// Package stats make calculations over matrices and vectors
package stat

// TODO float, categorial, binomial

type Stat struct {
	name string
	rows int
	cols int
	arr  []int
}

func (s *Stat) Show() map[Factor]float32 {
	// Набор данных для эксорта в Qolumn
}
