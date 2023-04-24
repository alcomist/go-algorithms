package tree

type Fenwick struct {
	t []int
}

func NewFenwickTree(ds []int) *Fenwick {

	fen := &Fenwick{}

	// make +1 space for 1 based index
	fen.t = make([]int, len(ds)+1)
	fen.init(ds)

	return fen
}

func (f *Fenwick) delta(i, v, n int) {

	for i <= n {
		f.t[i] += v
		i = i + (i & -i)
	}
}

func (f *Fenwick) init(ds []int) {

	for i, d := range ds {
		f.delta(i+1, d, len(ds))
	}
}

func (f *Fenwick) sum(i int) int {

	ret := 0

	for i > 0 {
		ret += f.t[i]
		i = i - (i & -i)
	}

	return ret
}

func (f *Fenwick) Sum(start, end int) int {

	return f.sum(end+1) - f.sum(start)
}
