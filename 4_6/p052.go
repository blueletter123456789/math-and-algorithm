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
	M = 1000000007
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	x := NumStdin()
	y := NumStdin()

	nume_a := 2*y - x
	nume_b := 2*x - y

	if nume_a % 3 != 0 || nume_b % 3 != 0 {
		fmt.Println(0)
		return
	}

	if nume_a < 0 || nume_b < 0 {
		fmt.Println(0)
		return
	}

	fmt.Println(ncr((x+y)/3, nume_a/3))
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func ncr(a, b int) int {
	lim := a
	if a < b {
		lim = b
	}
	fact := make([]int, lim+1)
	fact[0] = 1
	for i := 1; i <= lim; i++ {
		fact[i] = (fact[i-1]*i) % M
	}
	
	return (fact[a] * pow((fact[b]*fact[a-b]) % M, M-2)) % M
}

func pow(x, r int) int {
	p := x
	res := 1
	for i := 0; i <= 30; i++ {
		if (r & (1 << i)) != 0 {
			res *= p
			res %= M
		}
		p *= p
		p %= M
	}

	return res
}

/* TLE source code
func main() {
	x := NumStdin()
	y := NumStdin()

	patterns := make([][]int, y+1)
	for i := 0; i <= y; i++ {
		patterns[i] = make([]int, x+1)
	}

	patterns[0][0] = 1
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i+1 <= y && j+2 <= x {
				patterns[i+1][j+2] += patterns[i][j]
				patterns[i+1][j+2] %= M
			}
			if i+2 <= y && j+1 <= x {
				patterns[i+2][j+1] += patterns[i][j]
				patterns[i+2][j+1] %= M
			}
		}
	}

	fmt.Println(patterns[x][y])
}
*/
