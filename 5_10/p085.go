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
	x := NumStdin()
	y := NumStdin()

	ans := false
	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			for c := b; c <= n; c++ {
				for d := c; d <= n; d++ {
					if a+b+c+d == x && a*b*c*d == y {
						ans = true
						break
					}
				}
			}
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
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
