package main

import (
	"fmt"
)

func main(){
	var a, b int
	fmt.Scanf("%v %v", &a, &b)

	for a > 0 && b > 0{
		if a > b{
			a %= b
		}else{
			b %= a
		}
	}
	if a == 0{
		fmt.Println(b)
	}else{
		fmt.Println(a)
	}
}
