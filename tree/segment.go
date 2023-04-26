package tree

import "github.com/alcomist/go-algorithms/index"

func midIndex(i, j int) int {
	return i + (j-i)/2
}

type Segment struct {
	d       []int
	t       map[int]int
	indexer index.Indexer
}

func NewSegmentTree(d []int) *Segment {

	seg := &Segment{}
	seg.d = d

	// use map for reducing extra spaces
	seg.t = make(map[int]int)

	//seg.indexer = &RootOneIndexer{}
	seg.indexer = &index.RootZeroIndexer{}

	seg.init(0, len(d)-1, seg.indexer.Root())
	return seg
}

func (s *Segment) init(start, end, node int) int {

	if start == end {
		s.t[node] = s.d[start]
		return s.t[node]
	} else {

		mid := midIndex(start, end)
		s.t[node] =
			s.init(start, mid, s.indexer.Left(node)) + s.init(mid+1, end, s.indexer.Right(node))
		return s.t[node]
	}
}

func (s *Segment) sum(start, end, node, left, right int) int {

	if left > end || right < start {
		return 0
	}

	// wrong if statement : if start <= left && right <= end {
	if left <= start && end <= right {

		return s.t[node]
	} else {

		mid := midIndex(start, end)
		return s.sum(start, mid, s.indexer.Left(node), left, right) + s.sum(mid+1, end, s.indexer.Right(node), left, right)
	}
}

func (s *Segment) Sum(left, right int) int {

	return s.sum(0, len(s.d)-1, s.indexer.Root(), left, right)
}
