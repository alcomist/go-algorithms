package distance

import "sort"

func distance(s1, s2 []rune) int {

	if len(s1) == 0 {
		return len(s2)
	}

	if len(s2) == 0 {
		return len(s1)
	}

	if s1[0] == s2[0] {
		return distance(s1[1:], s2[1:])
	}

	ds := make([]int, 3)
	ds[0] = distance(s1[1:], s2)     // delete
	ds[1] = distance(s1[1:], s2[1:]) // update
	ds[2] = distance(s1, s2[1:])     // insert

	sort.Ints(ds)
	return ds[0] + 1
}

func Distance(s1, s2 string) int {

	return distance([]rune(s1), []rune(s2))
}
