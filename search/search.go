package search

type Searcher interface {
	Len() int
	Less(key any, i int) bool
	Equal(key any, i int) bool
}

type IntSlice []int

func (x IntSlice) Len() int                  { return len(x) }
func (x IntSlice) Less(key any, i int) bool  { return key.(int) < x[i] }
func (x IntSlice) Equal(key any, i int) bool { return key.(int) == x[i] }
