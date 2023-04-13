package tree

import (
	"github.com/alcomist/go-algorithms/provider"
	"math/rand"
	"testing"
	"time"
)

func sum(a []int, left, right int) int {

	s := 0

	for i := left; i <= right; i++ {
		s += a[i]
	}

	return s
}

func TestSegmentTree(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	size := 1000

	a := provider.IntSlice(size)

	for i := 0; i < 1000; i++ {

		tree := NewSegmentTree(a)

		s, e := provider.RandomRange(size - 1)

		s1 := tree.Sum(s, e)
		s2 := sum(a, s, e)

		if s1 != s2 {
			t.Errorf("[%d:%d] segment tree sum error (tree sum=>%d, sum=>%d)\n", s, e, s1, s2)
		}
	}

}
