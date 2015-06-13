// Package fragment make data extraction from html-fragment
package fragment

import (
	"github.com/c4e8ece0/sereio/doc/attr"
)

type Fragment struct {
	content string
	attr    []*attr.Attr
}
