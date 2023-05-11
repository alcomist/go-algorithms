package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	//Data set 1: yes
	//Data set 2: yes
	//Data set 3: no

	for i := 0; i < n; i++ {
		scanner.Scan()

		s := scanner.Text()

		ss := strings.Fields(s)
		if len(ss) == 3 {
			r := interleaving([]rune(ss[0]), []rune(ss[1]), []rune(ss[2]))
			if r {
				writer.WriteString(fmt.Sprintf("Data set %d: yes\n", i+1))
			} else {
				writer.WriteString(fmt.Sprintf("Data set %d: no\n", i+1))
			}
		}
	}
}
