// Package delim contains definitions for splitting data to the blocks
package delim

var set = struct{}{}

// Default vars with appending of se.Profile
var (
	Block = map[string]struct{}{
		"html":       set,
		"head":       set,
		"body":       set,
		"div":        set,
		"table":      set,
		"table > td": set,
		"ul":         set,
		"ol":         set,
		"li":         set,
	}

	Paragraph = map[string]struct{}{
		"p":          set,
		"hr":         set,
		"h1":         set,
		"h2":         set,
		"h3":         set,
		"h4":         set,
		"h5":         set,
		"h6":         set,
		"p > b":      set, // TODO: special case for hX search in sere/meta
		"p > strong": set, // special case
	}

	Sentence = map[string]struct{}{
		". [A-Z]{3,}": set,
		"p":           set, // special case
	}

	Word = map[string]struct{}{
		" ": set,
		"-": set,
		",": set,
		".": set,
	}
)
