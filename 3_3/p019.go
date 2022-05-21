package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CardsType int

const(
	Red CardsType = iota+1
	Yello
	Blue
)

var scan = bufio.NewScanner(os.Stdin)

func init(){
	buf := make([]byte, 1000000)
	scan.Buffer(buf, 1000000)
}

func IntStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil{
		panic(err)
	}
	return num
}

func IntsStdin(n int) []int{
	scan.Scan()
	lst := scan.Text()

	nums := make([]int, n)
	for i, v := range strings.Split(lst, " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

func ncr(n int) int {
	return n*(n-1) / 2
}

func main() {
	n := IntStdin()
	lst := IntsStdin(n)

	r := 0
	y := 0
	b := 0
	for _, c := range lst {
		switch c{
		case int(Red):
			r += 1
		case int(Yello):
			y += 1
		case int(Blue):
			b += 1
		}
	}
	
	fmt.Println(ncr(r) + ncr(y) + ncr(b))
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// var scan = bufio.NewScanner(os.Stdin)

// func init(){
// 	buf := make([]byte, 1000000)
// 	scan.Buffer(buf, 1000000)
// }

// func main() {
// 	scan.Scan()
// 	_ = scan.Text()
// 	cards := map[string]int{"1": 0, "2": 0, "3": 0}

// 	scan.Scan()
// 	lst := strings.Split(scan.Text(), " ")

// 	for _, c := range lst {
// 		cards[c] += 1
// 	}

// 	sum := 0
// 	for _, v := range cards {
// 		if v >= 2 {
// 			sum += (v*(v-1))/2
// 		}
// 	}
// 	fmt.Println(sum)
// }
