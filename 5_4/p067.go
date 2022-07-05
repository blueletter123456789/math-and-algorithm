package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	h := NumStdin()
	w := NumStdin()

	area_y := make([]int, h)
	area_x := make([]int, w)
	in_lst := make([][]int, h)
	for i := 0; i < h; i++ {
		row := make([]int, w)
		row_sum := 0
		for j := 0; j < w; j++ {
			num := NumStdin()
			row[j] = num
			row_sum += num
			area_x[j] += num
		}
		in_lst[i] = row
		area_y[i] = row_sum
	}

	for i := 0; i < h; i++ {
		row := make([]string, w)
		for j := 0; j < w; j++ {
			row[j] = strconv.Itoa(area_y[i] + area_x[j] - in_lst[i][j])
		}
		fmt.Println(strings.Join(row, " "))
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
