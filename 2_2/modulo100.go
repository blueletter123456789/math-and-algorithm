package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdin)

func IntStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(nil)
	}
	return num
}

func main(){
	var n int
	fmt.Scan(&n)
	scan.Split(bufio.ScanWords)
	ans := 0
	for i := 0; i < n; i++ {
		ans += IntStdin()
	}
	fmt.Println(ans%100)
}
