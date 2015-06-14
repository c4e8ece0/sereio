// Package attr extract data from attributes of html-tags
package attr

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Создание нового объекта
func New(r io.Reader) *Attr {
	return &Attr{src: r}
}

type Attr struct {
	src io.Reader
}

// Predefined rules for attr.Fetch()
// Tag-rules delimited by "_" for two parts: Tag1Tag2_Attr1Attr2. Tag can be "*" for all.
var (
	Any_Title   = map[string]map[string]struct{}{"*": {"title": struct{}{}}}
	Any_Class   = map[string]map[string]struct{}{"*": {"class": struct{}{}}}
	A_Title     = map[string]map[string]struct{}{"a": {"title": struct{}{}}}
	Input_Title = map[string]map[string]struct{}{"input": {"title": struct{}{}}}
	A_Href      = map[string]map[string]struct{}{"a": {"href": struct{}{}}}
	Img_Alt     = map[string]map[string]struct{}{"img": {"alt": struct{}{}}}
	Img_Src     = map[string]map[string]struct{}{"img": {"src": struct{}{}}}
	Script_Src  = map[string]map[string]struct{}{"script": {"src": struct{}{}}}
)

// Helpers for parameters saveTag and saveAttr of Fetch()
var (
	NoTags    = false
	NeedTags  = true
	NoAttrs   = false
	NeedAttrs = true
)

// Method Fetch() extract attr values for specified tags ("*" for all) and their attrs
func (a *Attr) Fetch(rule map[string]map[string]struct{}, saveTag, saveAttr bool) (vals, tags, attrs []string) {
	z := html.NewTokenizer(a.src)
	for {
		//z.NextIsNotRawText()
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return

		case html.StartTagToken, html.SelfClosingTagToken:
			t, hasAttr := z.TagName()
			if !hasAttr {
				continue
			}
			realtag := string(t)
			if realtag == "script" {
				if skip_script(z) {
					continue
				} else {
					break
				}
			}
			if false {
				fmt.Printf("\n\t%v", z.Token())
			}
			searchtag := realtag
			if _, exists := rule[searchtag]; !exists {
				if _, exists := rule["*"]; !exists {
					continue
				} else {
					searchtag = "*"
				}
			}

			for {
				v, x, hasMore := z.TagAttr()
				attr := string(v)
				value := string(x)
				if _, exists := rule[searchtag][attr]; exists {
					vals = append(vals, value)
					if saveTag {
						tags = append(tags, realtag)
					}
					if saveAttr {
						attrs = append(attrs, attr)
					}
				}
				if !hasMore {
					break
				}
			}
		}
	}
	return
}

// Function skip_script get current *html.Tokenizer and skip content of <script>
// Returns recomendation to continue work with that struct.
func skip_script(z *html.Tokenizer) bool {
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return false
		case html.EndTagToken:
			t, _ := z.TagName()
			if string(t) == "script" {
				return true
			}
		}
	}
}

// Count frequency of tokens in attributes
func (a *Attr) Count(rule map[string]map[string]struct{}) (stat map[string]int) {
	stat = make(map[string]int)
	kw, _, _ := a.Fetch(rule, NoTags, NoAttrs)
	for _, t := range kw {
		stat[t]++
	}
	return
}

func (a *Attr) CountAnyTitle() map[string]int {
	return a.Count(Any_Title)
}

func (a *Attr) CountInputTitle() map[string]int {
	return a.Count(Input_Title)
}

func (a *Attr) CountATitle() map[string]int {
	return a.Count(A_Title)
}

func (a *Attr) CountAnyClass() map[string]int {
	return a.Count(Any_Class)
}

func (a *Attr) CountImgAlt() map[string]int {
	return a.Count(Img_Alt)
}
