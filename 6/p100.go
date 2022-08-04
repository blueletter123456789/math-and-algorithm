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
	q := int(NumStdin())
	for i := 0; i < q; i++ {
		x, y, z := NumStdin(), NumStdin(), NumStdin()
		t := int(NumStdin())

		A := [3][3]float64{
			{1-x, y, 0},
			{0, 1-y, z},
			{x, 0, 1-z},
		}
		
		ans := powMatrix(A, t)

		out := []string{"", "", ""}
		for i, row := range ans {
			out[i] = strconv.FormatFloat(row[0] + row[1] + row[2], 
				'f', 10, 64)
		}
		fmt.Println(strings.Join(out, " "))
	}
}

func NumStdin() float64 {
	scan.Scan()
	num, err := strconv.ParseFloat(scan.Text(), 64)
	if err != nil {
		panic(err)
	}
	return num
}

func powMatrix(A [3][3]float64, t int) (ret [3][3]float64) {

	ret[0][0] = 1
	ret[1][1] = 1
	ret[2][2] = 1

	for i := 0; i < 25; i++ {
		if (t & (1 << i)) != 0 {
			ret = multiMatrix(A, ret)
		}
		A = multiMatrix(A, A)
	}

	return
}

func multiMatrix(A, B [3][3]float64) (C [3][3]float64) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return
}

// func powMatrix(A [][]float64, t int) (ret [][]float64) {

// 	for i := 0; i < len(A); i++ {
// 		ret = append(ret, make([]float64, len(A[i])))
// 		ret[i][i] = 1
// 	}

// 	for i := 0; i < 30; i++ {
// 		if (t & (1 << i)) != 0 {
// 			ret = multiMatrix(A, ret)
// 		}
// 		A = multiMatrix(A, A)
// 	}

// 	return
// }

// func multiMatrix(A, B [][]float64) (C [][]float64) {
// 	for i := 0; i < len(A); i++ {
// 		C = append(C, make([]float64, len(A[i])))
// 		for j := 0; j < len(A[i]); j++ {
// 			for k := 0; k < len(B); k++ {
// 				C[i][j] += A[i][k] * B[k][j]
// 			}
// 		}
// 	}

// 	return
// }
