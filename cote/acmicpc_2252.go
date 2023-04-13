package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func topo(om map[int][]int) []int {

	var order []int

	seen := make(map[int]bool)

	var visit func(items []int)

	visit = func(items []int) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visit(om[item])
				order = append(order, item)
			}
		}
	}

	var keys []int
	for key := range om {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	visit(keys)
	return order
}

func main() {

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int

	fmt.Fscan(r, &n, &m)

	height := make(map[int][]int)

	for i := 0; i < m; i++ {

		var r1, r2 int
		fmt.Fscan(r, &r1, &r2)

		orders, ok := height[r1]
		if !ok {
			orders = make([]int, 0)
		}

		orders = append(orders, r2)
		height[r1] = orders
	}

	orders := topo(height)

	for _, order := range orders {
		fmt.Fprintf(w, "%d ", order)
	}

	fmt.Fprintln(w)
}
