package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scan = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, 800000)
	scan.Buffer(buf, 800000)
}

func main()  {
	scan.Scan()
	_ = scan.Text()
	coins := map[string]int{"100": 0, "200": 0, "300": 0, "400": 0}

	scan.Scan()
	nums := strings.Split(scan.Text(), " ")
	for _, v := range nums{
		coins[v] += 1
	}

	comb := coins["100"]*coins["400"] + coins["200"]*coins["300"]
	fmt.Println(comb)
}
