package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// Node : is something we manage in a priority queue.
type node struct {
	v int
	w int
}

// A PQ implements heap.Interface and holds Items.
type PQ []*node

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].w < pq[j].w
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(*node))
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

const INF int = 987654321

func main() {

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m, s int

	fmt.Fscan(r, &n, &m)
	fmt.Fscan(r, &s)
	//fmt.Println(vertex, edge, start)

	adj := make([][]node, n+1)

	for i := 0; i < m; i++ {

		var u, v, w int
		fmt.Fscan(r, &u, &v, &w)

		adj[u] = append(adj[u], node{v: v, w: w})
	}

	pq := PQ{}
	visited := make([]int, n+1)

	heap.Push(&pq, &node{s, 1})
	visited[s] = 1

	for pq.Len() > 0 {
		here := heap.Pop(&pq).(*node)

		if visited[here.v] < here.w {
			continue
		}

		for _, next := range adj[here.v] {
			if visited[next.v] == 0 || visited[next.v] > visited[here.v]+next.w {
				visited[next.v] = visited[here.v] + next.w
				heap.Push(&pq, &node{next.v, visited[next.v]})
			}
		}
	}

	for i := 1; i <= n; i++ {
		if visited[i] == 0 {
			w.WriteString("INF")
		} else {
			w.WriteString(strconv.Itoa(visited[i] - 1))
		}
		w.WriteByte('\n')
	}
}
