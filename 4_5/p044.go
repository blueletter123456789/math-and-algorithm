package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initBufSize = 10000
	maxBufSize = 10000000
)

var scan = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()
	m := NumStdin()
	G := make([][]int, n+1)
	
	var u, v int
	for i := 0; i < m; i++ {
		u = NumStdin()
		v = NumStdin()
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	dist := make([]int, n)
	dist[0] = 0
	for i := 1; i < n; i++ {
		dist[i] = -1
	}

	que := New(n)
	que.Enqueue(1)
	for !que.IsEmpty(){
		current := que.Top()
		que.Dequeue()
		for _, v := range G[current] {
			if dist[v-1] != -1 {
				continue
			}
			dist[v-1] = dist[current-1] + 1
			que.Enqueue(v)
		}
	}

	for _, d := range dist {
		fmt.Println(d)
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

type Queue struct {
	data []int
	size int
}

func New(cap int) *Queue {
	return &Queue{
		data: make([]int, 0, cap),
		size: 0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Top() int {
	return q.data[0]
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
	q.size++
}

func (q *Queue) Dequeue() bool {
	if q.IsEmpty() {
		return false
	}
	q.data = q.data[1:]
	q.size--
	return true
}
