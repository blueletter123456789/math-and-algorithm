package main

import (
	"bufio"
	"fmt"
	"math"
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
	k := NumStdin()

	is_one := false
	v := make([]int, k)
	for i := 0; i < k; i++ {
		v[i] = NumStdin()
		if v[i] == 1 {
			is_one = true
			break
		}
	}

	if is_one {
		fmt.Println(n)
		return
	}

	pat := int(math.Pow(2.0, float64(k)))
	ans := 0

	for i := 1; i < pat; i++ {
		tgt := 1
		is_even := 0
		for j := 0; j < k; j++ {
			if (i & (1 << j)) != 0 {
				tgt = lcm(tgt, v[j])
				is_even++
			}
		}
		if is_even % 2 == 1 {
			ans += n / tgt
		} else {
			ans -= n / tgt
		}
	}
	
	fmt.Println(ans)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func lcm(a, b int) int {
	return (a*b)/gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}
