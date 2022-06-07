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
	n := NumStdin()
	q := NumStdin()
 
	snows := make([]int, n+1)
	for i := 0; i < q; i++ {
		l := NumStdin()
		r := NumStdin()
		x := NumStdin()
		snows[l-1] += x
		snows[r] -= x
	}
 
	ans := make([]byte, 0, 1000000)
	for i := 1; i < n; i++ {
		if snows[i] == 0{
			ans = append(ans, "="...)
		}else if snows[i] > 0 {
			ans = append(ans, "<"...)
		}else {
			ans = append(ans, ">"...)
		}
	}
	fmt.Println(string(ans))
}
