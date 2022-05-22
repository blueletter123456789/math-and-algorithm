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
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func ncr(n, r int) int{
	res := 1
	for i := 1; i <= r; i++{
		res *= n-i+1
		res /= i
	}
	return res
}

func main() {
	scan.Split(bufio.ScanWords)

	n := IntStdin()
	r := IntStdin()
	fmt.Println(ncr(n, r))
}

/* Sample code
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
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func calc(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func ncr(n, r int) int{
	return calc(n) / (calc(r)*calc(n-r))
}

func main(){
	scan.Split(bufio.ScanWords)

	n := IntStdin()
	r := IntStdin()

	fmt.Println(ncr(n, r))
}
*/
