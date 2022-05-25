package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scan = bufio.NewScanner(os.Stdin)
var A []int
var C []int

const (
	initBufSize = 10000
	maxBufSize = 1000000000000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func NumsStdin(n int) []int {
	scan.Scan()
	nums := make([]int, n)

	for i, v := range strings.Split(scan.Text(), " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

func mergeSort(l, r int) {
	if r - l == 1 {
		return
	}
	m := (l + r) / 2
	
	mergeSort(l, m)
	mergeSort(m, r)

	c1 := l
	c2 := m
	cnt := 0

	for c1 != m || c2 != r{
		if c1 == m {
			C[cnt] = A[c2]
			c2++
		}else if c2 == r {
			C[cnt] = A[c1]
			c1++
		}else{
			if A[c1] < A[c2] {
				C[cnt] = A[c1]
				c1++
			}else{
				C[cnt] = A[c2]
				c2++
			}
		}
		cnt++
	}

	for i := 0; i < cnt; i++{
		A[l+i] = C[i]
	}
}

func main() {
	n := NumStdin()
	A = NumsStdin(n)
	C = make([]int, n)

	mergeSort(0, n)

	res := make([]string, n)
	for i, v := range A {
		res[i] = strconv.Itoa(v)
	}

	fmt.Println(strings.Join(res, " "))

}
