package string

import (
	"math"
	"unicode/utf8"
)

const (
	base = 2
	mod  = 1000000007
)

type RabinKarp struct {
}

func hash(r rune, x, y int) int {
	return int(r) * int(math.Pow(float64(x), float64(y))) % mod
}

func CalculateHash(rs []rune) int {

	n := len(rs)

	h := 0

	for i, r := range rs {
		h += hash(r, base, n-i+1)
	}

	return h
}

func NewRK() *RabinKarp {

	r := &RabinKarp{}
	return r
}

func (r *RabinKarp) Search(p, s string) []int {

	pos := make([]int, 0)

	plen := utf8.RuneCountInString(p)
	slen := utf8.RuneCountInString(s)

	if plen > slen {
		return pos
	}

	if p == s {
		pos = append(pos, 0)
		return pos
	}

	prs := []rune(p)
	hash := CalculateHash(prs)

	rs := []rune(s)

	for i := 0; i <= slen-plen; i++ {

		srs := rs[i : i+plen]

		shash := CalculateHash(srs)
		if hash == shash {

			if string(prs) == string(srs) {
				pos = append(pos, i)
			}
		}
	}

	return pos
}
