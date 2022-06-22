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
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

type Edge struct {
	V, W int
}

type EdgeCost struct {
	v, cost, index int
}

type PriorityQueue []*EdgeCost

func (pq PriorityQueue) Len()int {return len(pq)}

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
	n := 10
	k := NumStdin()
	G := make([][]Edge, k)
	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			G[i] = append(G[i], Edge{V: (i*n+j)%k, W: j})
		}
	}

	dist := make([]int, k)
	for i := 0; i < k; i++ {
		dist[i] = 1 << 30
	}

	que := make(PriorityQueue, 1, k)
	que[0] = &EdgeCost{
		cost: 20,
		v: 0,
		index: 0,
	}
	heap.Init(&que)

	for que.Len() > 0{
		edge := heap.Pop(&que).(*EdgeCost)
		if dist[edge.v] < edge.cost {
			continue
		}
		for _, to := range G[edge.v] {
			cost := dist[edge.v] + to.W
			if edge.v == 0 {
				cost = to.W
			}
			if dist[to.V] > cost {
				dist[to.V] = cost
				e := &EdgeCost{cost: dist[to.V], v: to.V}
				heap.Push(&que, e)
				que.update(e, e.v, e.cost)
			}
		}
	}

	fmt.Println(dist[0])
	
}


func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

// O(v^2)
// func main() {
// 	n := 10
// 	k := NumStdin()
// 	G := make([][]Edge, k)
// 	for i := 0; i < k; i++ {
// 		for j := 0; j < n; j++ {
// 			if i == 0 && j == 0 {
// 				continue
// 			}
// 			G[i] = append(G[i], Edge{V: (i*n+j)%k, W: j})
// 		}
// 	}

// 	dist := make([]int, k)
// 	used := make([]bool, k)
// 	for i := 0; i < k; i++ {
// 		dist[i] = 1 << 30
// 	}

// 	flg := true

// 	for i := 0; i < k; i++{
// 		v := -1
		
// 		for u := 0; u < k; u++ {
// 			if !used[u] && (v == -1 || dist[u] < dist[v]) {
// 				v = u
// 			}
// 		}

// 		if v == -1 {
// 			break
// 		}

// 		used[v] = true
// 		for _, e := range G[v] {
// 			if flg {
// 				dist[e.V] = min(e.W, dist[e.V])
// 			}
// 			dist[e.V] = min(dist[e.V], dist[v] + e.W)
// 		}
// 		flg = false
// 	}

// 	fmt.Println(dist[0])
	
// }

// func min(x, y int) int {
// 	if x > y {
// 		return y
// 	}
// 	return x
// }
