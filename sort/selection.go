package sort

func selection(data Sorter) {

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

func Selection(data Sorter) {
	selection(data)
}
