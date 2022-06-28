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

	if n < (abs(x) + abs(y)) {
		fmt.Println("No")
		return
	}
	if n % 2 == (abs(x) + abs(y)) % 2 {
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
