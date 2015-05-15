// Package qolumn make the data collecting and preparation

package qolumn

func (s *Qol) Task() {
	// => [passage_id][]int32{<word_ids>} +++ map[token]count?

	// Common
	doc.View(SomeParametrisedClosure)

	// Or
	doc.
		Use.(se.YandexProfile). // se.Profile must contains specific parser delimiters and rules (context td)
		From(
		// Group of splitters with se.Profile usage (do closure realy need?)
		SplitByTagCollectionClosure(TAG_TD, TAG_DIV), // array of 1-len array blocks, i.e. <p><table>, <td><div>
		PassagesClosure(),                            // array of tokens array of bags (delimiter fields) [OR]
		BagOfWordClosure()).                          // array of words (bag of words)
		With(SYNSET).
		Make(TokenCount). // map[string]int
		Filter(Frequency(3)).
		Name("passages.synset.count.log") // auto || from predefined views
	Norm(log)

	// Qolumn.SetNext(doc.

}

// HDF5 data collection - HANDMADE INSANE CRAP?!!
func (q *Qolumn) SetNext(qol Qolumn) {
	name := qol.Name()
	flat := qol.IsFlat()
	q.Lock()
	i := len(q.list)
	q.list = append(q.list, name)
	q.flat[i] = flat
	TypedMakeCopy(q.data[i], qol.data)
	q.Unlock()
}

func (q *Qolumn) FromToNew(fromqol, toqol string, f func()) {
	// SetNext + Doc.Chain()
}

func (q *Qolumn) Create(name string) {
	// Create dir at diskplace(name)
	// Create sync worker to recovery copy
	// + flag on sync
	// + flag on write
	// + flag on compact
	// Format: gobot?
	// Files/Dirs to provider (sql, mongo, qml, inmem)
	// Each Qol to file... or FUSE (HDF5 more likely)?
}

type Qol struct {
	data interface{}
	size int
	name string
}
