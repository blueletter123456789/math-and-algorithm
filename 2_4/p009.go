package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdin)

func IntStdin() int {
	scan.Scan()
	res, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return res
}

func main(){
	scan.Split(bufio.ScanWords)
	
	n := IntStdin()
	s := IntStdin()
	A := make([]int, 0, n)
	
	for i := 0; i < n; i++ {
		A = append(A, IntStdin())
	}

	ans := false
	for i := 0; i < (1 << n); i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				sum += A[j]
			}
		}
		if s == sum{
			ans = true
			break
		}
	}
	if ans{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}
