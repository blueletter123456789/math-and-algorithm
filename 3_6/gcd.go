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

func gcd(a, b int) int {
	if b == 0{
		return a
	}
	return gcd(b, a%b)
}

func main() {
	scan.Split(bufio.ScanWords)

	a := NumStdin()
	b := NumStdin()

	res := gcd(a, b)
	fmt.Println(res)
}
