package sort

func merge(data Sorter, s, e int) {
	//fmt.Println(data[0])
	//fmt.Println(s, e)
}

// mergeSort range is inclusive (s <= index <= e)
func mergeSort(data Sorter, s, e int) {

	// if s >= e mean slice length is 0
	if s >= e {
		return
	}

	mid := (s + e) / 2

	mergeSort(data, s, mid)
	mergeSort(data, mid+1, e)
	merge(data, s, e)
}

func MergeSort(data Sorter) {

	mergeSort(data, 0, data.Len()-1)
}
