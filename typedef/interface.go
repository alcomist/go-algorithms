package typedef

type Lesser interface {
	Less(i, j int) bool
}

type Swapper interface {
	Swap(i, j int)
}

type Lener interface {
	Len() int
}

type Sorter interface {
	Lener
	Swapper
	Lesser
}

type Assigner interface {
	Assign(i int, x any)
}

type Valuer interface {
	Value(i int) any
}

type Searcher interface {
	Lener
	Less(key any, i int) bool
	Equal(key any, i int) bool
}

type HeapInterface interface {
	Sorter
	Push(x any)
	Pop() any
}
