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
	h := NumStdin()
	w := NumStdin()

	if h == 1 || w == 1 {
		fmt.Println(1)
		return
	}
	
	area := h * w
	if area % 2 == 0 {
		fmt.Println(area / 2)
	} else {
		fmt.Println(area / 2 + 1)
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
