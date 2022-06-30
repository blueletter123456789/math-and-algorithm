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
	maxBufSize = 1000000000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()
	k := NumStdin()

	telLst := make([]int, n+1)
	for i := 1; i <= n; i++ {
		t := NumStdin()
		telLst[i] = t
	}

	init := getInit(telLst, n)
	initLen := len(init)

	looped := getLooped(telLst, init[initLen-1], n)
	// fmt.Println(init, looped)

	if k < len(init) {
		fmt.Println(init[k])
	} else {
		k -= initLen
		loopedLen := len(looped)
		fmt.Println(looped[k%loopedLen])
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

func getInit(telLst []int, n int) []int {
	seen := make([]bool, n+1)
	init := make([]int, 0, n)

	t := 1
	seen[t] = true
	init = append(init, t)
	for {
		v := telLst[t]
		if seen[v] {
			break
		}
		seen[v] = true
		init = append(init, v)
		t = v
	}
	
	return init
}

func getLooped(telLst []int, t, n int) []int {
	seen := make([]bool, n+1)
	looped := make([]int, 0, n)

	for {
		v := telLst[t]
		if seen[v] {
			break
		}
		seen[v] = true
		looped = append(looped, v)
		t = v
	}
	
	return looped
}
