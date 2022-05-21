package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdin)

func IntStdin() int{
	scan.Scan()
	in_int, err := strconv.Atoi(scan.Text())
	if err != nil{
		panic(err)
	}
	return in_int
}

func gcd(a, b int) int {
	for a > 0 && b > 0{
		if a > b{
			a %= b
		}else{
			b %= a
		}
	}
	if a == 0{
		return b
	}
	return a
}

func main(){
	scan.Split(bufio.ScanWords)
	n := IntStdin()
	a := IntStdin()

	for i := 0; i < n-1; i++{
		b := IntStdin()
		a = gcd(a, b)
	}

	fmt.Println(a)
}
