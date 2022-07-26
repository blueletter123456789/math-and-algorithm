package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scan = bufio.NewScanner(os.Stdin)

const (
	initBufSize = 10000
	maxBufSize = 10000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
	scan.Split(bufio.ScanWords)
}

func main() {
	n := NumStdin()
	s := StrStdin()

	cnt := make(map[string]int, 2)
	cnt["("], cnt[")"] = 0, 0
	ans := true

	for i := 0; i < n; i++ {
		cnt[string(s[i])]++
		if cnt["("] < cnt[")"] {
			ans = false
			break
		}
	}
	if ans && cnt["("] == cnt[")"] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func StrStdin() string {
	scan.Scan()
	return scan.Text()
}

func NumStdin() int {
	num, err := strconv.Atoi(StrStdin())
	if err != nil {
		panic(err)
	}
	return num
}

// Smaple code
// func main() {
// 	n := NumStdin()
// 	s := StrStdin()

// 	depth := 0
// 	for i := 0; i < n; i++ {
// 		if string(s[i]) == "(" {
// 			depth++
// 		}
// 		if string(s[i]) == ")" {
// 			depth--
// 		}
// 		if depth < 0 {
// 			fmt.Println("No")
// 			return
// 		}
// 	}

// 	if depth == 0 {
// 		fmt.Println("Yes")
// 	} else {
// 		fmt.Println("No")
// 	}
// }
