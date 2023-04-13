package search

func binary(data Searcher, x any) int {

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

func Binary(data Searcher, x any) int {
	return binary(data, x)
}
