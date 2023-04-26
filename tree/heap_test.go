package tree

import (
	"github.com/alcomist/go-algorithms/provider"
	"reflect"
	"sort"
	"testing"
)

func TestIntHeapSort(t *testing.T) {

	s1 := provider.RandomIntSlice(1000)
	s2 := append([]int(nil), s1...)

	sort.Ints(s1)

	h := NewIntHeap()
	for _, n := range s2 {
		h.Push(n)
	}

	s3 := make([]int, 0, len(s2))
	for h.Len() > 0 {
		s3 = append(s3, h.Pop())
	}

	if reflect.DeepEqual(s1, s3) == false {
		t.Errorf("int heap sort differ\n\t=>%q\n\t=>%q\n",
			provider.IntSliceToString(s1),
			provider.IntSliceToString(s3))
	}
}
