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
	maxBufSize = 10000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()

	a := int(math.Sqrt(float64(n)))
	ans := 1 << 59
	for i := 1; i <= a; i++ {
		b := n / i
		if (i * b) != n {
			continue
		}
		ans = min(ans, 2 * i + 2 * b)
	}

	fmt.Println(ans)
}

// Another answer
// func main() {
// 	n := NumStdin()

// 	ans := 1 << 59
// 	for i := 1; i*i <= n; i++ {
// 		if n % i == 0 {
// 			ans = min(ans, 2*i + 2*(n/i))
// 		}
// 	}

// 	fmt.Println(ans)
// }

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
