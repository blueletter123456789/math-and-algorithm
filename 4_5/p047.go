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

var seen []int
var G [][]int

func main() {
	n := NumStdin()
	m := NumStdin()

	G = make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := NumStdin(), NumStdin()
		G[u-1] = append(G[u-1], v-1)
		G[v-1] = append(G[v-1], u-1)
	}

	seen = make([]int, n)
	for i := 0; i < n; i++ {
		seen[i] = -1
	}

	res := false
	for i := 0; i < n; i++ {
		if seen[i] != -1{
			continue
		}
		res = dfs(i, -1)
		if !res{
			break
		}
	}

	if res {
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
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

func dfs(u, flg int) bool {
	if flg == -1{
		flg = 0
	}
	seen[u] = flg

	ret := true
	for _, v := range G[u] {
		if seen[u] == seen[v]{
			ret = false
			break
		}else if seen[v] == -1{
			ret = dfs(v, 1-flg)
			if !ret{
				break
			}
		}
	}

	return ret
}
