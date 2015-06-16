// Package tag checks tags structure and make some statistics
package tag

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Создание нового объекта
func New(r io.Reader) *TagList {
	return &List{src: r}
}

type Tag struct {
	Name string
	Type html.TokenType
}

type List struct {
	src  io.Reader
	list []Tag
}

// Build array of tags in document
func (t *List) Build() {
	if len(t.list) > 0 {
		return
	}

	z := html.NewTokenizer(a.src)
	for {
		z.NextIsNotRawText()
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken, html.SelfClosingTagToken, html.EndTagToken, html.DoctypeToken:
			p, _ := z.TagName()
			realtag := string(t)
			t.list = append(t.list, []Tag{Name: realtag, Type: tt})

			if realtag == "script" && skip_script(z) {
				t.list = append(t.list, []Tag{Name: realtag, Type: html.EndTagToken})
			}
		}
	}
	return
}

// Todo
func (t *List) String() {
	return fmt.Sprintf("%v", t)
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