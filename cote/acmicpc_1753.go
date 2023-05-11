package main

import (
	"bufio"
	"fmt"
	"github.com/alcomist/go-algorithms/tree"
	"os"
	"strconv"
)

// Node : is something we manage in a priority queue.
type node struct {
	next   int
	weight int
}

type Nodes []*node

func (ns *Nodes) Len() int           { return len(*ns) }
func (ns *Nodes) Less(i, j int) bool { return (*ns)[i].weight < (*ns)[j].weight }
func (ns *Nodes) Swap(i, j int)      { (*ns)[i], (*ns)[j] = (*ns)[j], (*ns)[i] }
func (ns *Nodes) Push(x any) {
	*ns = append(*ns, x.(*node))
}

func (ns *Nodes) Pop() any {
	old := *ns
	n := len(old)
	item := old[n-1]
	*ns = old[0 : n-1]
	return item
}

const INF int = 987654321

func main() {

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	// vertex, edge, start
	var v, e, s int

	fmt.Fscan(r, &v, &e)
	fmt.Fscan(r, &s)

	graph := make([][]node, v+1)

	for i := 0; i < e; i++ {

		var current, next, weight int
		fmt.Fscan(r, &current, &next, &weight)

		graph[current] = append(graph[current], node{next: next, weight: weight})
	}

	nodes := Nodes{}
	h := tree.NewHeap(&nodes)

	weights := make([]int, v+1)
	for i, _ := range weights {
		weights[i] = INF
	}

	weights[s] = 0

	h.Push(&node{s, 0})

	for h.Len() > 0 {

		cur := h.Pop().(*node)

		current := cur.next
		weight := cur.weight

		if weights[current] < weight {
			continue
		}

		for _, item := range graph[current] {

			nextWeight := weight + item.weight

			if nextWeight < weights[item.next] {
				weights[item.next] = nextWeight
				h.Push(&node{item.next, nextWeight})
			}
		}
	}

	for _, weight := range weights[1:] {
		if weight == INF {
			w.WriteString("INF")
		} else {
			w.WriteString(strconv.Itoa(weight))
		}
		w.WriteByte('\n')
	}
}
