package backtrack

func genParens(left, right, n int, rs string) []string {

	ss := make([]string, 0)
	if left == n && right == n {
		ss = append(ss, rs)
	}

	if left < n {
		ss = append(ss, genParens(left+1, right, n, rs+"(")...)
	}
	if right < left {
		ss = append(ss, genParens(left, right+1, n, rs+")")...)
	}

	return ss
}

func GenerateParentheses(n int) []string {

	ss := genParens(0, 0, n, "")
	return ss
}
