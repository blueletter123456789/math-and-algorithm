package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scan = bufio.NewScanner(os.Stdin)

func numStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func numsStdin(n int) []int {
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

func main() {
	n := numStdin()
	ints := numsStdin(n)

	fmt.Println(ints)

	for i:= 0; i < n; i++ {
		tmp := i
		for j := i+1; j < n; j++ {
			if ints[tmp] > ints[j] {
				tmp = j
			}
		}
		ints[i], ints[tmp] = ints[tmp], ints[i]
	}

	fmt.Println(ints)
}