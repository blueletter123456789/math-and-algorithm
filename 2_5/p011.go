package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	var n int
	fmt.Scan(&n)
	ans := make([]string, 0, n)

	for i := 2; i <= n; i++{
		is_prime := true
		for j := 2; j < i; j ++{
			if i % j == 0 {
				is_prime = false
				break
			}
		} 
		if is_prime {
			tmp := strconv.FormatInt(int64(i), 10)
			ans = append(ans, tmp)
		}
	}
	fmt.Println(strings.Join(ans, " "))
}
