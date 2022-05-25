package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scan = bufio.NewScanner(os.Stdin)

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

func interSum(l, r int, lst []int) int {
	if r - l <= 1{
		return lst[l]
	}
	m := (l + r) / 2

	s1 := interSum(l, m, lst)
	s2 := interSum(m, r, lst)

	return s1 + s2
}

func main() {
	n := NumStdin()
	lst := NumsStdin(n)

	res := interSum(0, n, lst)
	fmt.Println(res)
}
