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

	parent := gcd(a, b)
	if a / parent > 1000000000000000000 / b {
		
		fmt.Println("Large")
		return
	}

	fmt.Println(a * (b / parent))
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}
