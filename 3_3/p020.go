package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	initBuffSize int = 10000
	maxBufSize int = 1000000
)

var scan = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, initBuffSize)
	scan.Buffer(buf, maxBufSize)
}

func IntStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func IntsStdin(n int) []int {
	nums := make([]int, n)

	scan.Scan()
	for i, v := range strings.Split(scan.Text(), " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

func main() {
	n := IntStdin()
	lst := IntsStdin(n)
	expected := 1000

	ans := 0
	for i := 0; i<n; i++ {
		for j := i+1; j<n; j++ {
			for k := j+1; k<n; k++ {
				for l := k+1; l<n; l++ {
					for m := l+1; m<n; m++ {
						if expected - (lst[i] + lst[j] + lst[k] + lst[l] + lst[m]) == 0 {
							ans += 1
						}
					}
				}
			}
		}
	}

	fmt.Println(ans)
}
