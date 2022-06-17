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

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
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
	fmt.Println(G)

	for i := 1; i <= n; i++ {
		fmt.Println("{")
		for j, v := range G[i] {
			fmt.Print("  ", v)
			if j != len(G[i])-1 {
				fmt.Println(",")
			}
		}
		fmt.Println("")
		fmt.Println("},")
	}
}
