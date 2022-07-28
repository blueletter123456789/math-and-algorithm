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
	a := NumStdin()
	b := NumStdin()
	c := NumStdin()

	// cが1の時はMAX_Bまでループする
	// 1 <= aのためcが1の時は場合分け
	if c == 1 {
		fmt.Println("No")
		return
	}

	// a < c^b
	// a/c < c^(b-1)
	x := 1.0
	for i := 1; i <= int(b); i++ {
		if (a / c) < x {
			fmt.Println("Yes")
			return
		}
		x *= c
	}
	fmt.Println("No")
}

func NumStdin() float64 {
	scan.Scan()
	num, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		panic(err)
	}
	return num
}
