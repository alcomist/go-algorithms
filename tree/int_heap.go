package tree

import "github.com/alcomist/go-algorithms/index"

type IntHeap struct {
	d       []int
	indexer index.Indexer
}

func NewIntHeap() *IntHeap {

	return &IntHeap{d: make([]int, 0), indexer: &index.RootZeroIndexer{}}
}

func (h *IntHeap) Len() int {
	return len(h.d)
}

func (h *IntHeap) up() {

	j := len(h.d) - 1
	for {

		i := h.indexer.Parent(j) // parent
		if i == j || h.d[i] < h.d[j] {
			break
		}

		h.d[i], h.d[j] = h.d[j], h.d[i]
		j = i
	}

}

func (h *IntHeap) Push(x int) {

	// add x to tail of h.d
	h.d = append(h.d, x)

	// then up heap
	h.up()
}

func (h *IntHeap) down() {

	i := h.indexer.Root()
	for {

		j1 := h.indexer.Left(i)
		if j1 >= len(h.d) || j1 < 0 { // j1 < 0 after int overflow
			break
		}

		j := j1 // left child
		if j2 := h.indexer.Right(i); j2 < len(h.d) && h.d[j2] < h.d[j1] {
			j = j2 // = 2*i + 2  // right child
		}

		if h.d[i] > h.d[j] {
			h.d[i], h.d[j] = h.d[j], h.d[i]
		}

		i = j
	}
}

func (h *IntHeap) Pop() int {

	top := h.d[0]
	h.d[0] = h.d[len(h.d)-1]
	h.d = h.d[0 : len(h.d)-1]

	// down heap
	h.down()
	return top
}
