package sereio

import (
	"strings"
)

type TokenID uint32
type WordID uint32
type PassageID uint32
type DocID uint32
type NameID uint32
type HostID uint32

type Weight float32
type Count uint32

type Storager interface {
	Set(id int) error
	Get(id int) (map[string]string, error)
}

type Getter interface {
	Get(id int) error
}

type Counter struct {
	Tokens []*TokenID
	Freq   map[TokenID]Count
}

type Token struct {
	Words []WordID
}

func (t *Token) Pack() {
}

type Word struct {
	src string
	lib interface{}
}

func (w *Word) Strict() string {
	return w.src
}

// Приведение к нижнему регистру
func (w *Word) Lower() string {
	return strings.ToLower(w.src)
}

// Получение базовой формы слова
func (w *Word) Base() string {
	return nlp.BaseForm(w.src)
}

// Расширение слова на группу синонимов
func (w *Word) Expand() []Word {
	var hl string = "" // -geo
}

// Предствление предложения
type Sentence struct {
	bag  []Word
	inst interface{}
}

// "ya", "translit", "?gg?"
func (s *Sentence) Translate() []string {
}
