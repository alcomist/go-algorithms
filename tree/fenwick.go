package tree

type Fenwick struct {
	t []int
}

func NewFenwickTree(ns []int) *Fenwick {

	f := &Fenwick{}

	f.t = make([]int, len(ns)+1)

	for i, n := range ns {
		f.delta(i+1, n)
	}

	return f
}

func (f *Fenwick) delta(i, n int) {

	for i < len(f.t) {
		f.t[i] += n
		i = i + (i & -i)
	}
}

// Sum = sum(hi+1) - sum(lo)
func (f *Fenwick) Sum(lo, hi int) int {

	return f.sum(hi+1) - f.sum(lo)
}

func (f *Fenwick) sum(i int) int {

	r := 0
	for i > 0 {
		r += f.t[i]
		i = i - (i & -i)
	}

	return r
}
