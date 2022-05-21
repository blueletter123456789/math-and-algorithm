package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scan = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, 2000000)
	scan.Buffer(buf, 2000000)
}

func IntStdin() int {
	scan.Scan()
	in_num, err := strconv.Atoi(scan.Text())
	if err != nil{
		panic(err)
	}
	return in_num
}

func IntsStdin(n int) []int{
	scan.Scan()
	lst := scan.Text()

	nums := make([]int, n)
	for i, v := range strings.Split(lst, " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

func gcd(a, b int) int {
	for a > 0 && b > 0 {
		if a > b {
			a %= b
		}else{
			b %= a
		}
	}
	if a == 0 {
		return b
	}
	return a
}

func lcm(a, b int) int {
	return (a/gcd(a, b)) * b
}

func main() {
	n := IntStdin()
	lst := IntsStdin(n)
	lcmNum := lcm(lst[0], lst[1])
	for i := 2; i < n; i++ {
		lcmNum = lcm(lcmNum, lst[i])
	}
	fmt.Println(lcmNum)
}
