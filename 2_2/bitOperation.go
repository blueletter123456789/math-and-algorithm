package main

import "fmt"

func main(){
	var num1, num2 int
	fmt.Scanf("%d %d", &num1, &num2)
	fmt.Printf("bit: num1 = %b, num2 = %b\n", num1, num2)
	fmt.Println("AND =", num1 & num2)
	fmt.Println("OR = ", num1 | num2)
	fmt.Println("XOR = ", num1 ^ num2)
}
