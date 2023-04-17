package string

func KmpTable(s string) []int {

	pi := make([]int, len(s))
	j := 0

	for i := 1; i < len(s); i++ {

		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}

		if s[i] == s[j] {
			j++
			pi[i] = j
		}
	}

	return pi
}
