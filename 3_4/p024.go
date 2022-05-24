package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initBufSize = 10000
	maxBufSize = 10000000
)

var scan = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
}

func IntStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func main(){
	scan.Split(bufio.ScanWords)

	n := IntStdin()
	var ans float64 = 0

	for i := 0; i < n; i++{
		p := IntStdin()
		q := IntStdin()
		ans += 1 / float64(p) * float64(q)
	}
	fmt.Println(ans)
}

