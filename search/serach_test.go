package search

import (
	"github.com/alcomist/go-algorithms/provider"
	"golang.org/x/exp/slices"
	"math/rand"
	"testing"
	"time"
)

type IntSlice []int

func (x IntSlice) Len() int                  { return len(x) }
func (x IntSlice) Less(key any, i int) bool  { return key.(int) < x[i] }
func (x IntSlice) Equal(key any, i int) bool { return key.(int) == x[i] }

func TestBinarySearch(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	s := provider.IntSlice(100)

	r := rand.Intn(100)

	a, ok := slices.BinarySearch(s, r)
	b := Binary(IntSlice(s), r)

	if ok && (a != b) {
		t.Errorf("binary search differ\n\t=>%v\n\t=>%v\n", a, b)
	}
}
