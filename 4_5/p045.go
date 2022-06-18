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

	G = make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := NumStdin(), NumStdin()
		G[u-1] = append(G[u-1], v-1)
		G[v-1] = append(G[v-1], u-1)
	}

	seen = make([]bool, n)
	fmt.Println(dfs(1))
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func dfs(u int) int {
	seen[u] = true
	res := 0
	cnt := 0
	for _, v := range G[u] {
		if u-v > 0 {
			cnt++
		}
		if seen[v] {
			continue
		}
		res += dfs(v)
	}
	if cnt == 1 {
		res++
	}
	return res
}


/* Another answer
func main(){
	n := NumStdin()
	m := NumStdin()
	counter := make([]int, n)
	for i := 0; i < m; i++ {
		u, v := NumStdin(), NumStdin()
		if u < v {
			counter[v-1]++
		}else{
			counter[u-1]++
		}
	}

	ans := 0
	for _, cnt := range counter {
		if cnt == 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
*/
