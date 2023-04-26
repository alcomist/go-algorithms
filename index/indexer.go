package index

// Indexer : indexer strategy interface with common methods
type Indexer interface {
	Root() int
	Left(i int) int
	Right(i int) int
	Parent(i int) int
}

type RootZeroIndexer struct {
}

func (r *RootZeroIndexer) Root() int {
	return 0
}

func (r *RootZeroIndexer) Left(i int) int {
	return i*2 + 1
}

func (r *RootZeroIndexer) Right(i int) int {
	return i*2 + 2
}

func (r *RootZeroIndexer) Parent(i int) int {
	return (i - 1) / 2
}

type RootOneIndexer struct {
}

func (r *RootOneIndexer) Root() int {
	return 1
}

func (r *RootOneIndexer) Left(i int) int {
	return i * 2
}

func (r *RootOneIndexer) Right(i int) int {
	return i*2 + 1
}

func (r *RootOneIndexer) Parent(i int) int {
	return i / 2
}
