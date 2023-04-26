package typedef

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Searcher interface {
	Len() int
	Less(key any, i int) bool
	Equal(key any, i int) bool
}

type HeapInterface interface {
	Sorter
	Push(x any)
	Pop() any
}
