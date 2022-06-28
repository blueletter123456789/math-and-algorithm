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

type Matrix struct {
	size int
	p [16][16]int
}

var lst2 [4][4]int = [4][4]int{
		{0, 0, 0, 1},
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{1, 0, 0, 1},
	}

var lst3 [8][8]int = [8][8]int{
	{0, 0, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 1},
	{0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 1, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 1, 0, 0, 1, 0},
}

var lst4 [16][16]int = [16][16]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	{0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
	{1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1},
}

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	k := NumStdin()
	n := NumStdin()

	A := Matrix{size: 1<<k}
	for i := 0; i < (1 << k); i++ {
		for j := 0; j < (1 << k); j++ {
			if k == 2 {
				A.p[i][j] = lst2[i][j]
			}else if k == 3 {
				A.p[i][j] = lst3[i][j]
			}else if k == 4 {
				A.p[i][j] = lst4[i][j]
			}
		}
	}

	B := pow(A, n)

	fmt.Println(B.p[(1 << k)-1][(1 << k)-1])
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func pow(A Matrix, n int) (ret Matrix) {
	p := A
	flg := true
	for i := 0; i < 60; i++ {
		if (n & (1 << i)) != 0 {
			if flg {
				ret = p
				flg = false
			} else {
				ret = multi(ret, p)
			}
		}
		p = multi(p, p)
	}
	return
}

func multi(a, b Matrix) (ret Matrix) {
	ret.size = a.size

	for i := 0; i < a.size; i++ {
		for j := 0; j < a.size; j++ {
			for k := 0; k < a.size; k++ {
				ret.p[i][j] += a.p[i][k] * b.p[k][j]
				ret.p[i][j] %= M
			}
		}
	}

	return
}
