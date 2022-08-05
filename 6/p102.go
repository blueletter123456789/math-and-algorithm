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
	c := StrStdin()

	ans := 0
	for i := 0; i < n; i++ {
		var code int
		switch string(c[i]) {
		case "B":
			code = 0
		case "W":
			code = 1
		case "R":
			code = 2
		default:
			panic("no character.")
		}

		ans += code * ncr(n-1, i)
		ans %= 3
	}

	if n % 2 == 0 {
		ans = (3 - ans) % 3
	}

	fmt.Println(string("BWR"[ans]))
}

func StrStdin() string {
	scan.Scan()
	return scan.Text()
}

func NumStdin() int {
	num, err := strconv.Atoi(StrStdin())
	if err != nil {
		panic(err)
	}
	return num
}

func ncr(n, r int) int {
	if n < 3 && r < 3 {
		if n < r {
			return 0
		}
		if n == 2 && r == 1 {
			return 2
		}
		return 1
	}

	return ncr(n / 3, r / 3) * ncr(n % 3, r % 3) % 3
}
