package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdin)

func IntsStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func main(){
	scan.Split(bufio.ScanWords)
	n := IntsStdin()
	ans := 0
	for i := 0; i < n; i++ {
		ans += IntsStdin()
	}
	fmt.Println(ans)
}

// func main(){
// 	sum := 0
// 	var n int
// 	fmt.Scan(&n)
// 	var in_val int
// 	for i := 0; i < n; i++{
// 		fmt.Scan(&in_val)
// 		sum += in_val
// 	}
// 	fmt.Println(sum)
// }
