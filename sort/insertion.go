package sort

func insertion(data Sorter) {

	n := data.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func Insertion(data Sorter) {
	insertion(data)
}
