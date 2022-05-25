package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func fact(n int) int {
	if n == 1{
		return 1
	}
	return fact(n-1) * n
}

func main() {
	n := NumStdin()
	res := fact(n)
	fmt.Println(res)
}
