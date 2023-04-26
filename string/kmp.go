package string

type Kmp struct {
	pi []int
	s  string
}

func build(s string) []int {

	m := len(s)
	pi := make([]int, m)

	j := 0
	for i := 1; i < m; i++ {
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

func (k *Kmp) Search(p string) []int {

	pos := make([]int, 0)

	n := len(k.s)
	m := len(p)
	j := 0

	for i := 0; i < n; i++ {
		for j > 0 && k.s[i] != p[j] {
			j = k.pi[j-1]
		}
		if k.s[i] == p[j] {
			if j == m-1 {
				pos = append(pos, i-m+1)
				j = k.pi[j]
			} else {
				j++
			}
		}
	}

	return pos
}

func NewKmp(s string) *Kmp {

	k := &Kmp{}
	k.s = s
	k.pi = build(s)
	return k
}
