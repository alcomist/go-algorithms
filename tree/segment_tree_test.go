package tree

import (
	"github.com/alcomist/go-algorithms/provider"
	"github.com/alcomist/go-algorithms/util"
	"testing"
)

func TestSegmentTree(t *testing.T) {

	//rand.Seed(time.Now().UnixNano())

	size := 1000

	a := provider.IntSlice(size)

	for i := 0; i < 1000; i++ {

		tree := NewSegmentTree(a)

		s, e := provider.RandomRange(size - 1)

		s1 := tree.Sum(s, e)
		s2 := util.IntSum(a, s, e)

		if s1 != s2 {
			t.Errorf("[%d:%d] segment tree sum error (tree sum=>%d, sum=>%d)\n", s, e, s1, s2)
		}
	}
}
