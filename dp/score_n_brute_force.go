package dp

func WaysToScore(n int) int {

	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	return WaysToScore(n-10) + WaysToScore(n-5) + WaysToScore(n-3)
}

func WaysToScore2(n int, ns []int, result *[][]int) {

	if n < 0 {
		return
	}

	if n == 0 {

		//ns2 := make([]int, len(ns))
		//copy(ns2, ns)
		*result = append(*result, ns)
		return
	}

	WaysToScore2(n-10, append(ns, 10), result)
	WaysToScore2(n-5, append(ns, 5), result)
	WaysToScore2(n-3, append(ns, 3), result)
}

func WaysToScoreDP(n int) int {

	ns := make([]int, n+1)

	ns[0] = 1

	for i := 1; i <= n; i++ {
		if i-3 >= 0 {
			ns[i] += ns[i-3]
		}
		if i-5 >= 0 {
			ns[i] += ns[i-5]
		}
		if i-10 >= 0 {
			ns[i] += ns[i-10]
		}
	}

	return ns[n]
}
