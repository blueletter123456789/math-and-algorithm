package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdin)

const (
	initBufSize = 10000
	maxBufSize = 1000000000000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
}

func NumStdin() float64 {
	scan.Scan()
	num, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	scan.Split(bufio.ScanWords)
	
	x1 := NumStdin()
	y1 := NumStdin()
	r1 := NumStdin()
	x2 := NumStdin()
	y2 := NumStdin()
	r2 := NumStdin()

	if r1 < r2 {
		x1, y1, r1, x2, y2, r2 = x2, y2, r2, x1, y1, r1
	}

	dis := math.Sqrt(math.Pow(x1 - x2, 2) + math.Pow(y1 - y2, 2))

	if r1+r2 < dis {
		fmt.Println(5)
	}else if r1+r2 == dis {
		fmt.Println(4)
	}else if dis+r2 > r1 {
		fmt.Println(3)
	}else if dis+r2 == r1 {
		fmt.Println(2)
	}else{
		fmt.Println(1)
	}
}

// func main() {
// 	scan.Split(bufio.ScanWords)
	
// 	x1 := NumStdin()
// 	y1 := NumStdin()
// 	r1 := NumStdin()
// 	x2 := NumStdin()
// 	y2 := NumStdin()
// 	r2 := NumStdin()

// 	dis := math.Sqrt(math.Pow(x1 - x2, 2) + math.Pow(y1 - y2, 2))

// 	if dis < math.Abs(r1-r2) {
// 		fmt.Println(1)
// 	}else if dis == math.Abs(r1-r2) {
// 		fmt.Println(2)
// 	}else if dis < r1+r2 {
// 		fmt.Println(3)
// 	}else if dis == r1+r2 {
// 		fmt.Println(4)
// 	}else {
// 		fmt.Println(5)
// 	}
// }
