package main

import (
	"fmt"
)

func IsPrime(n int) bool {
	ans := true
	for i := 2; i*i <= n; i++{
		if n % i == 0{
			ans = false
			break
		}
	}
	return ans
}

func main(){
	var n int
	fmt.Scan(&n)
	res := IsPrime(n)
	if res{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}
