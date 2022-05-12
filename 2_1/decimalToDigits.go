package main

import "fmt"

func main(){
	var n int
	var digits string
	fmt.Scan(&n)
	for n > 0{
		if n % 2 == 1 {
			digits = "1" + digits
		}else{
			digits = "0" + digits
		}
		n /= 2
	}
	fmt.Println(digits)
}
