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

type Point struct {
	X, Y int
}

type Queue struct {
	data []*Point
	size int
}

func New(cap int) *Queue {
	return &Queue{
		data: make([]*Point, 0, cap),
		size: 0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Top() *Point {
	return q.data[0]
}

func (q *Queue) Enqueue(val *Point) {
	q.data = append(q.data, val)
	q.size++
}

func (q *Queue) Dequeue() bool {
	if q.IsEmpty() {
		return false
	}
	q.data = q.data[1:]
	q.size--
	return true
}

func main() {
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	height := NumStdin()
	width := NumStdin()
	sp := &Point{Y: NumStdin()-1, X: NumStdin()-1}
	gp := &Point{Y: NumStdin()-1, X: NumStdin()-1}

	maps := make([]string, height)
	for i := 0; i < height; i++ {
		maps[i] = StrStdin()
	}

	dist := make([][]int, 0, height)
	for i := 0; i < height; i++ {
		row := make([]int, width)
		dist = append(dist, row)
		for j := 0; j < width; j++ {
			dist[i][j] = -1
		}
	}
	dist[sp.Y][sp.X] = 0

	que := New(height*width)
	que.Enqueue(sp)
	for !que.IsEmpty(){
		cp := que.Top()
		que.Dequeue()
		for i := 0; i < 4; i++ {
			ny, nx := cp.Y-dy[i], cp.X-dx[i]
			if ny < 0 || ny >= height || nx < 0 || nx >= width{
				continue
			}
			if dist[ny][nx] != -1 || string(maps[ny][nx]) == "#"{
				continue
			}
			dist[ny][nx] = dist[cp.Y][cp.X] + 1
			que.Enqueue(&Point{Y: ny, X: nx})
		}
	}
	fmt.Println(dist[gp.Y][gp.X])
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
