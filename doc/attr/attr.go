// Package attr extract data from attributes of tags and make some statistics
package attr

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// TODO: Self-typed returns, i.e. "vals Vals" instead of "vals []string" + export interfaces
// TODO: ++ collect all attrs, like tags

// Create new object
func New(r io.Reader) *Attr {
	return &Attr{src: r}
}

//
type Attr struct {
	src io.Reader
}

// Helpers for parameters saveTag and saveAttr of Fetch()
var (
	NoTags    = false
	NeedTags  = true
	NoAttrs   = false
	NeedAttrs = true
	NoIndex   = false
	NeedIndex = true
)

// For better readability
var (
	set = struct{}{}
)

// Predefined rules for attr.Fetch()
// Tag-rules delimited by "_" for two parts: Tag1Tag2_Attr1Attr2. Tag can be "*" for all.
var (
	Any_Resource = map[string]map[string]struct{}{
		"*":      {"src": set, "href": set},
		"script": {"src": set},
	}
	Any_Any     = map[string]map[string]struct{}{"*": {"*": set}}
	Any_Title   = map[string]map[string]struct{}{"*": {"title": set}}
	Any_Class   = map[string]map[string]struct{}{"*": {"class": set}}
	A_Title     = map[string]map[string]struct{}{"a": {"title": set}}
	Input_Title = map[string]map[string]struct{}{"input": {"title": set}}
	A_Href      = map[string]map[string]struct{}{"a": {"href": set}}
	Img_Alt     = map[string]map[string]struct{}{"img": {"alt": set}}
	Img_Src     = map[string]map[string]struct{}{"img": {"src": set}}
	Script_Src  = map[string]map[string]struct{}{"script": {"src": set}}
)

// Extract attr values for specified tags ("*" for all) and their attrs
func (a *Attr) Fetch(rule map[string]map[string]struct{}, saveTag, saveAttr, saveIndex bool) (vals, tags, attrs []string, tagIndex []uint32) {
	z := html.NewTokenizer(a.src)
	var i uint32 = 0
	for {
		//z.NextIsNotRawText()
		tt := z.Next()
		i++
		switch tt {
		case html.ErrorToken:
			return

		case html.StartTagToken, html.SelfClosingTagToken:
			t, hasAttr := z.TagName()

			realtag := string(t)
			searchtag := realtag
			if _, exists := rule[searchtag]; !exists {
				if _, exists := rule["*"]; !exists {
					continue
				} else {
					searchtag = "*"
				}
			}

			if realtag == "script" {
				if _, exists := rule["script"]; !exists {
					if skip_script(z) {
						continue
					} else {
						break
					}
				}
			}

			if !hasAttr {
				continue
			}

			for {
				v, x, hasMore := z.TagAttr()
				attr := string(v)
				value := string(x)
				_, exists := rule[searchtag][attr]
				_, all := rule[searchtag]["*"]
				if exists || all {
					vals = append(vals, value)
					if saveTag {
						tags = append(tags, realtag)
					}
					if saveAttr {
						attrs = append(attrs, attr)
					}
					if saveTagIndex {
						tagIndex = append(tagIndex, i)
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

// Get current *html.Tokenizer and skip content of <script>
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

func (a *Attr) CountAnyResource() map[string]int {
	return a.Count(Any_Resource)
}

func (a *Attr) CountAnyClass() map[string]int {
	return a.Count(Any_Class)
}

func (a *Attr) CountImgAlt() map[string]int {
	return a.Count(Img_Alt)
}
