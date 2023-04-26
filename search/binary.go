package search

import "github.com/alcomist/go-algorithms/typedef"

func binary(data typedef.Searcher, x any) int {

	lo := 0
	hi := data.Len() - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if data.Equal(x, mid) {
			return mid
		} else if data.Less(x, mid) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}

func Binary(data typedef.Searcher, x any) int {
	return binary(data, x)
}
