package tree

import (
	"github.com/alcomist/go-algorithms/index"
	"github.com/alcomist/go-algorithms/typedef"
)

type Heap struct {
	hi      typedef.HeapInterface
	indexer index.Indexer
}

func NewHeap(hi typedef.HeapInterface) *Heap {

	return &Heap{hi: hi, indexer: &index.RootZeroIndexer{}}
}

func (h *Heap) Len() int {
	return h.hi.Len()
}

func (h *Heap) up() {

	j := h.Len() - 1
	for {

		i := h.indexer.Parent(j) // parent
		//if i == j || h.d[i] < h.d[j] {
		if i == j || h.hi.Less(i, j) {
			break
		}

		//h.d[i], h.d[j] = h.d[j], h.d[i]
		h.hi.Swap(i, j)
		j = i
	}

}

func (h *Heap) Push(x any) {

	//h.d = append(h.d, x)
	h.hi.Push(x)

	// add x to tail of h.d
	// then up heap
	h.up()
}

func (h *Heap) down() {

	i := h.indexer.Root()
	for {

		j1 := h.indexer.Left(i)
		if j1 >= h.Len() || j1 < 0 { // j1 < 0 after int overflow
			break
		}

		j := j1 // left child
		if j2 := h.indexer.Right(i); j2 < h.Len() && h.hi.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}

		if h.hi.Less(j, i) {
			h.hi.Swap(i, j)
		}

		i = j
	}
}

func (h *Heap) Pop() any {

	h.hi.Swap(0, h.Len()-1)

	top := h.hi.Pop()

	// down heap
	h.down()
	return top
}
