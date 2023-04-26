package sort

import "github.com/alcomist/go-algorithms/typedef"

func insertion(data typedef.Sorter) {

	n := data.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func Insertion(data typedef.Sorter) {
	insertion(data)
}
