package tree

// Indexer : indexer strategy interface with common methods
type Indexer interface {
	root() int
	left(i int) int
	right(i int) int
	parent(i int) int
}

type RootZeroIndexer struct {
}

func (r *RootZeroIndexer) root() int {
	return 0
}

func (r *RootZeroIndexer) left(i int) int {
	return i*2 + 1
}

func (r *RootZeroIndexer) right(i int) int {
	return i*2 + 2
}

func (r *RootZeroIndexer) parent(i int) int {
	return (i - 1) / 2
}

type RootOneIndexer struct {
}

func (r *RootOneIndexer) root() int {
	return 1
}

func (r *RootOneIndexer) left(i int) int {
	return i * 2
}

func (r *RootOneIndexer) right(i int) int {
	return i*2 + 1
}

func (r *RootOneIndexer) parent(i int) int {
	return i / 2
}

func midIndex(i, j int) int {
	return i + (j-i)/2
}

type Segment struct {
	d       []int
	t       map[int]int
	indexer Indexer
}

func NewSegmentTree(d []int) *Segment {

	seg := &Segment{}
	seg.d = d

	// use map for reducing extra spaces
	seg.t = make(map[int]int)

	//seg.indexer = &RootOneIndexer{}
	seg.indexer = &RootZeroIndexer{}

	seg.init(0, len(d)-1, seg.indexer.root())
	return seg
}

func (s *Segment) init(start, end, node int) int {

	if start == end {
		s.t[node] = s.d[start]
		return s.t[node]
	} else {

		mid := midIndex(start, end)
		s.t[node] =
			s.init(start, mid, s.indexer.left(node)) + s.init(mid+1, end, s.indexer.right(node))
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
		return s.sum(start, mid, s.indexer.left(node), left, right) + s.sum(mid+1, end, s.indexer.right(node), left, right)
	}
}

func (s *Segment) Sum(left, right int) int {

	return s.sum(0, len(s.d)-1, s.indexer.root(), left, right)
}
