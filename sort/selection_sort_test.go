package sort

import (
	"github.com/alcomist/go-algorithms/provider"
	"reflect"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {

	s1 := provider.RandomIntSlice(100)
	s2 := append([]int(nil), s1...)

	sort.Ints(s1)
	Selection(IntSlice(s2))

	if reflect.DeepEqual(s1, s2) == false {
		t.Errorf("selection sort differ\n\t=>%q\n\t=>%q\n",
			provider.IntSliceToString(s1),
			provider.IntSliceToString(s2))
	}
}

func BenchmarkDefaultSort(b *testing.B) {

	s := provider.RandomIntSlice(100)

	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
}

func BenchmarkSelectionSort(b *testing.B) {

	s := provider.RandomIntSlice(100)

	for i := 0; i < b.N; i++ {
		Selection(IntSlice(s))
	}
}
