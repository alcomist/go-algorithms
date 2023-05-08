package string

type Kmp struct {
}

func BuildKmpTable(rs []rune) []int {

	n := len(rs)

	pi := make([]int, n)

	j := 0
	for i := 1; i < n; i++ {

		for j > 0 && rs[i] != rs[j] {
			j = pi[j-1]
		}

		if rs[i] == rs[j] {
			j++
			pi[i] = j
		}
	}

	return pi
}

func NewKmp() *Kmp {

	k := &Kmp{}
	return k
}

func (k *Kmp) Search(p, s string) []int {

	pr := []rune(p)
	pi := BuildKmpTable(pr)

	sr := []rune(s)
	pos := make([]int, 0)

	m := len(pr)
	n := len(sr)

	j := 0

	for i := 0; i < n; i++ {

		for j > 0 && sr[i] != pr[j] {
			j = pi[j-1]
		}

		if sr[i] == pr[j] {
			if j == m-1 {
				pos = append(pos, i-m+1)
				j = pi[j]
			} else {
				j++
			}
		}
	}

	return pos
}
