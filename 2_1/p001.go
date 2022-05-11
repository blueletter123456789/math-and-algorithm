package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StrStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func IntStdin() (int, error) {
	in_str := StrStdin()
	return strconv.Atoi(in_str)
}

func main(){
	var x int = 5
	n, err := IntStdin()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(n + x)
}
