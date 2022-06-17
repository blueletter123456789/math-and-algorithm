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

	seen := make([]bool, n+1)
	seen[0] = true
	stack := New(n)
	stack.Push(1)
	for !stack.IsEmpty() {
		cur := stack.Top()
		stack.Pop()

		for _, v := range G[cur] {
			if seen[v]{
				continue
			}
			seen[v] = true
			stack.Push(v)
		}
	}

	for _, result := range seen {
		if !result{
			fmt.Println("The graph is not connected.")
			return
		}
	}
	fmt.Println("The graph is connected.")
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

type Stack struct {
	data []int
	size int
}

func New(cap int) *Stack {
	return &Stack{
		data: make([]int, 0, cap),
		size: 0,
	}
}

func (st *Stack) IsEmpty() bool {
	return st.size == 0
}

func (st *Stack) Top() int {
	return st.data[st.size-1]
}

func (st *Stack) Push(val int) {
	st.data = append(st.data, val)
	st.size++
}

func (st *Stack) Pop() bool {
	if st.IsEmpty() {
		return false
	}
	st.size--
	st.data = st.data[:st.size]
	return true
}
