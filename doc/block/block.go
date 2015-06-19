// Package block split document to fragments by set of tags
package block

import (
	"io/ioutil"

	"github.com/c4e8ece0/sereio/doc/attr"
)

// For better readability
var (
	set = struct{}{}
)

//
var (
	BlockDelimeterDefault = map[string]struct{}{
		"html":       set,
		"head":       set,
		"body":       set,
		"div":        set,
		"table":      set,
		"table > td": set, // special case
		"ul":         set,
		"ol":         set,
		"li":         set,
	}
)

//
var (
	ParagraphDelimeterDefault = map[string]struct{}{
		"p":          set,
		"hr":         set,
		"h1":         set,
		"h2":         set,
		"h3":         set,
		"h4":         set,
		"h5":         set,
		"h6":         set,
		"p > b":      set, // special case
		"p > strong": set, // special case
	}
)

//
func New(src io.Reader, delim Delimeter) {
	return &List{content: ioutil.ReadAll(r), make(attr.Attr, 0)}
}

//
type List struct {
	content string
	attr    attr.List
}
