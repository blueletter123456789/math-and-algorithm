package main

import (
	"bufio"
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

func main() {
	n := NumStdin()
	m := NumStdin()

	G := make([][]int, n)
	for i := 0; i < m; i++ {
		a, b := NumStdin()-1, NumStdin()-1
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}

	que := New(n)
	que.Enqueue(0)
	dist[0] = 0
	for !que.IsEmpty() {
		cur := que.Top()
		que.Dequeue()
		for _, e := range G[cur] {
			if dist[e] >= 0 || dist[cur]+1 > 120 {
				continue
			}
			dist[e] = dist[cur] + 1
			que.Enqueue(e)
		}
	}

	ans := make([]byte, 0, 10000000)
	for _, d := range dist {
		if d < 0 {
			ans = append(ans, "120\n"...)
		}else{
			ans = append(ans, (strconv.Itoa(d) + "\n")...)
		}
	}
	fmt.Println(string(ans))
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}
