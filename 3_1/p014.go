package main

import (
	"fmt"
)

func main(){
	var n int
	fmt.Scan(&n)

	for i := 2; i*i <= n; i++{
		for n % i == 0{
			fmt.Print(i, " ")
			n /= i
		}
	}

	if n >= 2{
		fmt.Print(n, " ")
	}
	fmt.Println()
}
