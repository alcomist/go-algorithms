package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func distance(s1, s2 []rune) int {

	if len(s1) == 0 {
		return len(s2)
	}

	if len(s2) == 0 {
		return len(s1)
	}

	if s1[0] == s2[0] {
		return distance(s1[1:], s2[1:])
	}

	ds := make([]int, 3)
	ds[0] = distance(s1[1:], s2)
	ds[1] = distance(s1[1:], s2[1:])
	ds[2] = distance(s1, s2[1:])

	sort.Ints(ds)
	return ds[0] + 1
}

func main() {

	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b string
	if r.Scan() {
		a = r.Text()
	}
	if r.Scan() {
		b = r.Text()
	}

	w.WriteString(fmt.Sprintf("%d", distance([]rune(a), []rune(b))))
}
