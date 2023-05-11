package interleaving

func interleaving(a, b, c []rune) bool {

	if len(a) == 0 && len(b) == 0 && len(c) == 0 {
		return true
	}

	if len(c) == 0 {
		return false
	}

	if len(a) == 0 && len(b) == 0 {
		return false
	}

	first, second := false, false

	if len(a) > 0 && len(c) > 0 && a[0] == c[0] {
		first = interleaving(a[1:], b, c[1:])
	}

	if len(b) > 0 && len(c) > 0 && b[0] == c[0] {
		second = interleaving(a, b[1:], c[1:])
	}

	return first || second
}

func Interleaving(a, b, c string) bool {

	return interleaving([]rune(a), []rune(b), []rune(c))
}
