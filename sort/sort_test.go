package sort

import (
	"github.com/alcomist/go-algorithms/provider"
	"reflect"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {

	s1 := provider.RandomIntSlice(1000)
	s2 := append([]int(nil), s1...)

	sort.Ints(s1)
	Selection(IntSlice(s2))

	if reflect.DeepEqual(s1, s2) == false {
		t.Errorf("selection sort differ\n\t=>%q\n\t=>%q\n",
			provider.IntSliceToString(s1),
			provider.IntSliceToString(s2))
	}
}

func TestInsertionSort(t *testing.T) {

	s1 := provider.RandomIntSlice(1000)
	s2 := append([]int(nil), s1...)

	sort.Ints(s1)
	Insertion(IntSlice(s2))

	if reflect.DeepEqual(s1, s2) == false {
		t.Errorf("insertion sort differ\n\t=>%q\n\t=>%q\n",
			provider.IntSliceToString(s1),
			provider.IntSliceToString(s2))
	}
}

func BenchmarkDefaultSort(b *testing.B) {

	s := provider.RandomIntSlice(1000)

	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
}

func BenchmarkSelectionSort(b *testing.B) {

	s := provider.RandomIntSlice(1000)

	for i := 0; i < b.N; i++ {
		Selection(IntSlice(s))
	}
}

func BenchmarkInsertionSort(b *testing.B) {

	s := provider.RandomIntSlice(1000)

	for i := 0; i < b.N; i++ {
		Insertion(IntSlice(s))
	}
}
