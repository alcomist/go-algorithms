package main

import (
	"bufio"
	"fmt"
	"os"
)

type kmp struct {
	table []int
}

func (k *kmp) build(s string) {

	rs := []rune(s)

	k.table = make([]int, len(rs))

	j := 0
	for i := 1; i < len(rs); i++ {

		for j > 0 && rs[i] != rs[j] {
			j = k.table[j-1]
		}

		if rs[i] == rs[j] {
			j++
			k.table[i] = j
		}
	}
}

func (k *kmp) search(s, p string) []int {

	pos := make([]int, 0)

	rs := []rune(s)
	rp := []rune(p)

	j := 0

	for i := 0; i < len(rs); i++ {

		for j > 0 && rs[i] != rp[j] {
			j = k.table[j-1]
		}

		if rs[i] == rp[j] {
			if j == len(rp)-1 {
				pos = append(pos, i-len(rp)+1)
				j = k.table[j]
			} else {
				j++
			}
		}
	}

	return pos
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	const maxCapacity = 1000001
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	t := scanner.Text()

	scanner.Scan()
	p := scanner.Text()

	k := &kmp{}
	k.build(p)

	pos := k.search(t, p)

	writer.WriteString(fmt.Sprintf("%d\n", len(pos)))
	for _, pp := range pos {
		writer.WriteString(fmt.Sprintf("%d ", pp+1))
	}
	writer.WriteString("\n")
}
