package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fenwick struct {
	t []int64
}

func newFenwick(ns []int64) *fenwick {

	f := &fenwick{}
	f.t = make([]int64, len(ns)+1)

	for i := 0; i < len(ns); i++ {
		f.update(i+1, len(ns), ns[i])
	}

	return f
}

func (f *fenwick) update(i, n int, num int64) {

	for i <= n {
		f.t[i] += num
		i = i + (i & -i)
	}
}

func (f *fenwick) sum(i int) int64 {

	var r int64

	for i >= 1 {
		r += f.t[i]
		i = i - (i & -i)
	}

	return r
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	nmk := strings.Split(scanner.Text(), " ")

	n, _ := strconv.Atoi(nmk[0])
	m, _ := strconv.Atoi(nmk[1])
	k, _ := strconv.Atoi(nmk[2])

	ns := make([]int64, 0)

	for i := 0; i < n; i++ {
		scanner.Scan()
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err == nil {
			ns = append(ns, num)
		}
	}

	f := newFenwick(ns)

	for i := 0; i < m+k; i++ {

		scanner.Scan()
		abc := strings.Split(scanner.Text(), " ")

		a, _ := strconv.Atoi(abc[0])
		b, _ := strconv.Atoi(abc[1])
		c, _ := strconv.ParseInt(abc[2], 10, 64)

		if a == 1 {
			// 더해주는 연산만 함
			t := c - ns[b-1]
			ns[b-1] = c
			f.update(b, len(ns), t)

		} else if a == 2 {
			// 구간합 값을 출력
			fmt.Fprintf(writer, "%d\n", f.sum(int(c))-f.sum(b-1))
		}
	}
}
