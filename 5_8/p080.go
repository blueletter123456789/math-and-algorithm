package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdin)

const (
	initBufSize = 10000
	maxBufSize = 10000000
	INF = 100000000000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

type Edge struct {
	T, W int
}

type EdgeCost struct {
	v, cost, index int
}

type PriorityQueue []*EdgeCost

func (pq PriorityQueue) Len()int {return len(pq)}

// 優先度の要素
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	edge := x.(*EdgeCost)
	edge.index = n
	*pq = append(*pq, edge)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	edge := old[n-1]
	old[n-1] = nil
	edge.index = -1
	*pq = old[0 : n-1]
	return edge
}

func (pq *PriorityQueue) update(edge *EdgeCost, v, cost int) {
	edge.v = v
	edge.cost = cost
	heap.Fix(pq, edge.index)
}

func main() {
	n := NumStdin()
	m := NumStdin()
	G := make([][]Edge, n)
	for i := 0; i < m; i++ {
		a := NumStdin()-1
		b := NumStdin()-1
		c := NumStdin()
		G[a] = append(G[a], Edge{b, c})
		G[b] = append(G[b], Edge{a, c})
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[0] = 0
	
	que := make(PriorityQueue, 1, n)
	que[0] = &EdgeCost{
		cost: 0,
		v: 0,
		index: 0,
	}
	heap.Init(&que)

	for que.Len() > 0 {
		min_v := heap.Pop(&que).(*EdgeCost)
		if dist[min_v.v] < min_v.cost {
			continue
		}
		for _, edge := range G[min_v.v] {
			cost := dist[min_v.v] + edge.W
			if dist[edge.T] > cost {
				dist[edge.T] = cost
				e := &EdgeCost{cost: dist[edge.T], v: edge.T}
				heap.Push(&que, e)
				que.update(e, e.v, e.cost)
			}
		}
	}

	if dist[n-1] < INF {
		fmt.Println(dist[n-1])
	} else {
		fmt.Println(-1)
	}
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

// func main() {
// 	n := NumStdin()
// 	m := NumStdin()
// 	G := make([][]Edge, n)
// 	for i := 0; i < m; i++ {
// 		a := NumStdin()-1
// 		b := NumStdin()-1
// 		c := NumStdin()
// 		G[a] = append(G[a], Edge{b, c})
// 		G[b] = append(G[b], Edge{a, c})
// 	}

// 	dist := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		dist[i] = INF
// 	}
// 	seen := make([]bool, n)
// 	dist[0] = 0

// 	for {
// 		v := -1
// 		for u := 0; u < n; u++ {
// 			if !seen[u] && (v == -1 || dist[u] < dist[v]) {
// 				v = u
// 			}
// 		}

// 		if v == -1 {
// 			break
// 		}

// 		seen[v] = true

// 		for _, e := range G[v] {
// 			dist[e.T] = min(dist[e.T], dist[v]+e.W)
// 		}
// 	}

// 	if dist[n-1] < INF {
// 		fmt.Println(dist[n-1])
// 	} else {
// 		fmt.Println(-1)
// 	}
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }
