package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func mergeSort(lst []int) []int {
	if len(lst) <= 1 {
		return lst
	}
	l, r := 0, len(lst)
	m := (l + r) / 2
	l1 := mergeSort(lst[l:m])
	l2 := mergeSort(lst[m:r])
	l3 := make([]int, len(lst))
	var i, j, k int
	for i < len(l1) && j < len(l2) {
		if l1[i] <= l2[j] {
			l3[k] = l1[i]
			i++
		}else{
			l3[k] = l2[j]
			j++
		}
		k++
	}
	for i < len(l1){
		l3[k] = l1[i]
		i++
		k++
	}
	for j < len(l2){
		l3[k] = l2[j]
		j++
		k++
	}
	return l3
}

func main() {
	n := NumStdin()
	lst := NumsStdin(n)

	res := make([]string, n)
	for i, num := range mergeSort(lst){
		res[i] = strconv.Itoa(num)
	}
	fmt.Println(strings.Join(res, " "))
}
