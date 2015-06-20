// Package text make many-dimension text view and generate some more practice variants
package text

var (
	// Move to profile later
	Sentence = map[string]bool{
		". [A-Z]{3,}": true,
		"p":           true, // special case
	}
	Word = map[string]bool{
		" ": true,
		"-": true,
		",": true,
		".": true,
	}
)

// View($)
const (
	STRICT = iota + 1
	LOWER
	SYNSET
	EXPAND_YANDEX // Need extern services
	COLLOC_YANDEX // Need extern services
	HL_YANDEX     // Need extern services
)

// WHERE AND WHEN CHECK spacebattle-tags with contexts?
// <div></div><span></span> is it ok?

type Token struct {
	str                      string
	spacebattle_danger_usage bool
}

type Text struct {
	content string
	sereio.DataFrame
}

func New() {
	return &Text{}
}

// Get only the html.Node with best text stats (stats will come from where?)
func BestCandidate() {

}
