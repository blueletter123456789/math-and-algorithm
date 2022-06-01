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
	maxBufSize = 100000000
)

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
}

func NumStdin() float64 {
	scan.Scan()
	num, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	scan.Split(bufio.ScanWords)

	a := NumStdin()
	b := NumStdin()
	h := NumStdin()
	m := NumStdin()

	radA := h*30.0 + m*0.5
	radB := m*6.0
	radAB := math.Abs(radA-radB)

	if radAB == 180 {
		fmt.Println(a+b)
	// 以下余弦定理
	}else if radAB < 180 {
		fmt.Println(math.Sqrt(math.Pow(a, 2) + 
			math.Pow(b, 2)-2*a*b*math.Cos(radAB * math.Pi / 180)))
	}else{
		fmt.Println(math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2) - 
			2*a*b*math.Cos((360.0-radAB) * math.Pi / 180)))
	}

}

// Sample code ベクトル
// func coordinate(a, angle float64) (x, y float64){
// 	x = a * math.Cos(angle * math.Pi / 180)
// 	y = a * math.Sin(angle * math.Pi / 180)
// 	return x, y
// }

// func main() {
// 	scan.Split(bufio.ScanWords)

// 	a := NumStdin()
// 	b := NumStdin()
// 	h := NumStdin()
// 	m := NumStdin()

// 	angleA := h*30.0 + m*0.5
// 	angleB := m*6.0

// 	ax, ay := coordinate(a, angleA)
// 	bx, by := coordinate(b, angleB)
	
// 	fmt.Println(math.Sqrt((ax-bx)*(ax-bx) + (ay-by)*(ay-by)))

// }
