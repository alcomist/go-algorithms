package search

import (
	"github.com/alcomist/go-algorithms/provider"
	"golang.org/x/exp/slices"
	"math/rand"
	"testing"
	"time"
)

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
