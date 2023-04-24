package util

func IntSum(a []int, left, right int) int {

	s := 0

	for i := left; i <= right; i++ {
		s += a[i]
	}

	return s
}
