package main

import (
	"fmt"
)

func main(){
	var n int

	fmt.Scan(&n)

	for i := 1; i*i <= n; i++{
		if n % i != 0{
			continue
		}
		fmt.Println(i)
		if i != n / i{
			fmt.Println(n/i)
		}
	}
}
