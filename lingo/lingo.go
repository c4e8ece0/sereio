// Package lingo make some fake linguistics
package lingo

type Word struct {
}

func New() *Word {
	return &Word{}
}

func BaseForm(s string) string {
	return "baseform"
}
