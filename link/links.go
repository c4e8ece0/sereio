// Package links extract links from document without mercy

package links

const (
	FOLLOW = 1 << iota
	NOFOLLOW
	INTERNAL // for Anchors()
	EXTERNAL // for Anchors()
	NOMATTER = FOLLOW | NOFOLLOW | INTERNAL | EXTERNAL
)

// Список внутренних ссылок
func (s *Doc) Internal(flags int) []Link {
	// Find links
	// Return only internal
}

// Список внешних ссылок
func (s *Doc) External(flags int) []Link {
	// Find links
	// Return only external
}

// Выбрать только тексты ссылок
func (s *Doc) Anchors(flags int) []string {
}

// Описание пути в ссылке или ресурсе
type Path struct {
	External  bool
	Relative  bool // first "/" not exists
	DotPrefix int8 // number of dots in prefix,
}

// Структура для описания документа
type Link struct {
	Pre              string
	Anchor           string
	Post             string
	Scheme           string // http, https, tel, javascript, mailto, call, any...
	Image            []Attr // Index of image + details
	Nofollow         bool
	NofollowExternal bool
	Rel              string
	path             string
	frag             string
}

// Проверка наличия аттрибута rel=nofollow
func (a *Link) Nofollow() bool {
	return a.Nofollow || a.NofollowExternal
}