package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initBufSize = 10000
	maxBufSize = 10000000
)

var scan = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, initBufSize)
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


func main() {
	scan.Split(bufio.ScanWords)

	n := IntStdin()
	var aSum float64 = 0
	var bSum float64 = 0
	for i := 0; i < n; i++ {
		a := IntStdin()
		aSum += 1.0/3.0  * float64(a)
	}
	for i := 0; i < n; i++ {
		b := IntStdin()
		bSum += 2.0/3.0 * float64(b)
	}
	fmt.Println(aSum + bSum)
}
