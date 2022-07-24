package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdin)

const (
	initBufSize = 10000
	maxBufSize = 10000000
)

type Part struct {
	S, E int
}

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()
	movies := make([]Part, n)
	for i := 0; i < n; i++ {
		movies[i].S = NumStdin()
		movies[i].E = NumStdin()
	}

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].E < movies[j].E
	})

	ans := 1
	prev_e := movies[0].E
	for i := 1; i < n; i++ {
		if prev_e <= movies[i].S {
			ans++
			prev_e = movies[i].E
		}
	}

	fmt.Println(ans)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}
