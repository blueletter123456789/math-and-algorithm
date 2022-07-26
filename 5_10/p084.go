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
	maxBufSize = 1000000000
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

	if c-a-b < 0 {
		fmt.Println("No")
		return
	}

	if 4*a*b < (c-a-b)*(c-a-b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 誤差
// func main() {
// 	a := NumStdin()
// 	b := NumStdin()
// 	c := NumStdin()

// 	ab := math.Sqrt(a) + math.Sqrt(b)
// 	if ab < math.Sqrt(c) {
// 		fmt.Println("Yes")
// 	} else {
// 		fmt.Println("No")
// 	}
// }

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}
