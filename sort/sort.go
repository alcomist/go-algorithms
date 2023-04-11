package sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Functor interface {
	Sort(Sorter, int, int)
}

func Sort(f Functor, data Sorter) {

	f.Sort(data, 0, data.Len())
}
