package provider

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandomIntSlice(size int) []int {

	nums := make([]int, 0, size)
	for i := 0; i < size; i++ {
		nums = append(nums, rand.Intn(size))
	}

	return nums
}

func RandomUniqueIntSlice(size int) []int {

	a := make([]int, 0, size)
	for i := 0; i < size; i++ {
		a = append(a, i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(size, func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	return a
}

// IntSlice : return int slice [0, size)
func IntSlice(size int) []int {

	a := make([]int, 0, size)
	for i := 0; i < size; i++ {
		a = append(a, i)
	}

	return a
}

func RandomRange(i int) (int, int) {

	rand.Seed(time.Now().UnixNano())

	s, e := i, i
	for s == e {
		s = rand.Intn(i)
		e = rand.Intn(i)
	}

	if s > e {
		s, e = e, s
	}

	return s, e
}

func IntSliceToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}
