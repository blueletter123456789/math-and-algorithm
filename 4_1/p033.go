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
	maxBufSize = 1000000000000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	scan.Split(bufio.ScanWords)
	ax := NumStdin()
	ay := NumStdin()
	bx := NumStdin()
	by := NumStdin()
	cx := NumStdin()
	cy := NumStdin()

	// BC間の相対的な距離
	bcx := cx - bx
	bcy := cy - by
	cbx := bx - cx
	cby := by - cy

	// B, Cに対するAの相対的な距離
	bax := ax - bx
	bay := ay - by
	cax := ax - cx
	cay := ay - cy

	abc := bax * bcx + bay * bcy
	acb := cax * cbx + cay * cby
	var ans float64

	if abc < 0 {
		ans = math.Sqrt(
			math.Pow(math.Abs(float64(ax-bx)), 2) + 
			math.Pow(math.Abs(float64(ay-by)), 2))
	}else if acb < 0{
		ans = math.Sqrt(
			math.Pow(math.Abs(float64(ax-cx)), 2) + 
			math.Pow(math.Abs(float64(ay-cy)), 2))
	}else{
		area := math.Abs(float64(bax*bcy - bay*bcx))
		bottom := math.Sqrt(
			math.Pow(float64(bcx), 2) + 
			math.Pow(float64(bcy), 2))
		ans = area / bottom
	}

	fmt.Println(strconv.FormatFloat(ans, 'f', -1, 64))
}
