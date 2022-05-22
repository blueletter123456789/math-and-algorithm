package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxBufSize = 1000000000000000
)

var scan = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, 100000)
	scan.Buffer(buf, maxBufSize)
}

func IntStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil{
		panic(err)
	}
	return num
}

func CntsStdin(n int) map[int]int{
	scan.Scan()
	lst := scan.Text()

	nums := make(map[int]int, n)
	for _, v := range strings.Split(lst, " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums[num] += 1
	}
	return nums
}

func main(){
	n := IntStdin()
	cnt := CntsStdin(n)

	expected := 100000
	ans := 0
	for i := 1; i < 50000; i++ {
		ans += cnt[i]*cnt[expected-i]
	}
	if cnt[50000] > 1{
		ans += (cnt[50000]*(cnt[50000]-1))/2
	}
	fmt.Println(ans)
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// const (
// 	maxBufSize = 1000000000000000
// )

// var scan = bufio.NewScanner(os.Stdin)

// func init() {
// 	buf := make([]byte, 100000)
// 	scan.Buffer(buf, maxBufSize)
// }

// func IntStdin() int {
// 	scan.Scan()
// 	num, err := strconv.Atoi(scan.Text())
// 	if err != nil{
// 		panic(err)
// 	}
// 	return num
// }

// func IntsStdin(n int) []int{
// 	scan.Scan()
// 	lst := scan.Text()

// 	nums := make([]int, n)
// 	for i, v := range strings.Split(lst, " ") {
// 		num, err := strconv.Atoi(v)
// 		if err != nil {
// 			panic(err)
// 		}
// 		nums[i] = num
// 	}
// 	return nums
// }

// func main(){
// 	n := IntStdin()
// 	lst := IntsStdin(n)

// 	cnt := make(map[int]int, n)
// 	for _, v := range lst {
// 		cnt[v] += 1
// 	}

// 	expected := 100000
// 	seen := make(map[int]bool, n)
// 	ans := 0
// 	for _, v := range lst {
// 		tgt := expected-v
// 		if seen[v] || seen[tgt]{
// 			continue
// 		}
// 		if cnt[tgt] > 0 && v != tgt{
// 			ans += cnt[v]*cnt[tgt]
// 			seen[v] = true
// 			seen[tgt] = true
// 		}else if v == tgt{
// 			ans += (cnt[v]*(cnt[v]-1))/2
// 			seen[v] = true
// 		}
// 	}
// 	fmt.Println(ans)
// }
