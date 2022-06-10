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

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	t := NumStdin()
	n := NumStdin()

	members := make([]int, t+1)
	for i := 0; i < n; i++ {
		l, r := NumStdin(), NumStdin()
		members[l]++
		members[r]--
	}

	sum := 0
	ans := make([]byte, 0, 1000000)
	for i := 0; i < t; i++{
		sum += members[i]
		ans = append(ans, (strconv.Itoa(sum) + "\n")...)
	}
	fmt.Print(string(ans))
}

// Code2
// func main() {
// 	t := NumStdin()
// 	n := NumStdin()

// 	members := make([]int, t+1)
// 	for i := 0; i < n; i++ {
// 		l, r := NumStdin(), NumStdin()
// 		members[l]++
// 		members[r]--
// 	}

// 	sum := 0
// 	for i := 0; i < t; i++{
// 		sum += members[i]
// 		fmt.Println(sum)
// 	}
// }
