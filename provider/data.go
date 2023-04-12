package provider

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandomIntSlice(size int) []int {

	nums := make([]int, 0, size)
	for i := 0; i < size; i++ {
		nums = append(nums, rand.Intn(size))
	}

	return nums
}

func IntSliceToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}
