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

var seen []bool
var G [][]int

func main() {
	n := NumStdin()
	m := NumStdin()

	G = make([][]int, n+1)
	var u, v int
	for i := 0; i < m; i++ {
		u = NumStdin()
		v = NumStdin()
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	seen = make([]bool, n+1)
	seen[0] = true

	dfs(1)

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

func dfs(current int) {
	seen[current] = true
	for _, v := range G[current] {
		if seen[v] {
			continue
		}
		dfs(v)
	}
}
