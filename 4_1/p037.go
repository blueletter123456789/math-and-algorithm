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
	maxBufSize = 100000000
)

type Point struct {
	X, Y int
}

func init() {
	buf := make([]byte, initBufSize)
	scan.Buffer(buf, maxBufSize)
}

func main() {
	scan.Split(bufio.ScanWords)

	p1 := scanPoint()
	p2 := scanPoint()

	p3 := scanPoint()
	p4 := scanPoint()

	c1 := cross(p1, p2, p3)
	c2 := cross(p1, p2, p4)
	c3 := cross(p3, p4, p1)
	c4 := cross(p3, p4, p2)

	if c1 > c2{
		c1, c2 = c2, c1
	}
	if c3 > c4{
		c3, c4 = c4, c3
	}

	if c1 == 0 && c2 == 0 && c3 == 0 && c4 == 0 {
		if compare(p1, p2, p3, p4) {
			fmt.Println("Yes")
		}else{
			fmt.Println("No")
		}
	}else {
		if c1 <= 0 && c2 >= 0 && c3 <= 0 && c4 >= 0{
			fmt.Println("Yes")
		}else{
			fmt.Println("No")
		}
	}
}

func NumStdin() int {
	scan.Scan()
	num, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func cross(p1, p2, p3 *Point) int {
	return (p2.X-p1.X)*(p3.Y-p1.Y) - (p2.Y-p1.Y)*(p3.X-p1.X)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func compare(p1, p2, p3, p4 *Point) bool {
	if p1.X == p2.X && p1.X == p3.X && p1.X == p4.X {
		if p1.Y > p2.Y {
			p1, p2 = p2, p1
		}
		if p3.Y > p4.Y{
			p3, p4 = p4, p3
		}
		return max(p1.Y, p3.Y) <= min(p2.Y, p4.Y)
	}else{
		if p1.X > p2.X{
			p1, p2 = p2, p1
		}
		if p3.X > p4.X{
			p3, p4 = p4, p3
		}
		return max(p1.X, p3.X) <= min(p2.X, p4.X)
	}
}

func scanPoint() *Point {
	x := NumStdin()
	y := NumStdin()
	return &Point{x, y}
}
