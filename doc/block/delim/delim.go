// Package delim contains definitions for splitting data to the blocks
package delim

// "set" replaces with bool to be able reassign type of split

// Default vars with appending of se.Profile
var (
	// Move to se.Profile later
	Block = map[string]bool{
		"html":       true,
		"head":       true,
		"body":       true,
		"div":        true,
		"table":      true,
		"table > td": true,
		"ul":         true,
		"ol":         true,
		"li":         true,
	}

	Paragraph = map[string]bool{
		"p":          true,
		"hr":         true,
		"h1":         true,
		"h2":         true,
		"h3":         true,
		"h4":         true,
		"h5":         true,
		"h6":         true,
		"p > b":      true, // TODO: special case for hX search in sere/meta
		"p > strong": true, // special case
	}

	// Sentence and Word moved to the sere/doc/text
)
