package sort

import "github.com/alcomist/go-algorithms/typedef"

func selection(data typedef.Sorter) {

	n := data.Len()
	for i := 0; i < n; i++ {

		min := i
		for j := i + 1; j < n; j++ {
			if data.Less(j, min) {
				min = j
			}
		}

		data.Swap(i, min)
	}
}

func Selection(data typedef.Sorter) {
	selection(data)
}
