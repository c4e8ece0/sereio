// Package block split document to fragments by set of tags
package block

import (
	"io/ioutil"

	"github.com/c4e8ece0/sereio/doc/attr"
	"github.com/c4e8ece0/sereio/doc/block/delim"
)

//
func New(src io.Reader) {
	return &List{content: ioutil.ReadAll(r), make(attr.Attr, 0)}
}

//
type List struct {
	content string
	attr    attr.List
}
