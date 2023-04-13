package tree

type Segment struct {
	d []int
	t map[int]int
}

func rootIndex() int {
	//return 1
	return 0
}

func leftChildIndex(i int) int {
	//return i*2
	return i*2 + 1
}

func rightChildIndex(i int) int {
	//return i*2 + 1
	return i*2 + 2
}

func parentIndex(i int) int {
	//return i / 2
	return (i - 1) / 2
}

func midIndex(i, j int) int {
	//return (i + j) / 2
	return i + (j-i)/2
}

func NewSegmentTree(d []int) *Segment {

	seg := &Segment{}
	seg.d = d
	seg.t = make(map[int]int)

	seg.init(0, len(d)-1, rootIndex())
	return seg
}

func (s *Segment) init(start, end, node int) int {

	if start == end {
		s.t[node] = s.d[start]
		return s.t[node]
	} else {

		mid := midIndex(start, end)
		s.t[node] =
			s.init(start, mid, leftChildIndex(node)) + s.init(mid+1, end, rightChildIndex(node))
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
		return s.sum(start, mid, leftChildIndex(node), left, right) + s.sum(mid+1, end, rightChildIndex(node), left, right)
	}
}

func (s *Segment) Sum(left, right int) int {

	return s.sum(0, len(s.d)-1, rootIndex(), left, right)
}
