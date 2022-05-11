package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func StrStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func IntsStdin() (after_conv []int) {
	before_conv := StrStdin()
	for _, v := range strings.Split(before_conv, " "){
		s, err := strconv.Atoi(v)
		if err != nil{
			panic(err)
		}
		after_conv = append(after_conv, s)
	}
	return
}

func main(){
	inputs := IntsStdin()
	ans := 0
	for _, v := range inputs{
		ans += v
	}
	fmt.Println(ans)
}

// fmtパッケージを使用した標準入力
// fmt.Scan(&str)
// fmt.Scanf("%d %d", &b, &c)

// bufio
// scanner := bufio.NewScanner(os.Stdin) 標準入力を受け付けるスキャナ
// or scanner.Buffer(make([]byte, 64*1024), 100001) 最大100,001文字に設定
// scanner.Scan() １行分の入力を取得する
// scanner.Text()
// inputs := strings.Split(scanner.Text(), " ") 空白区切り

/* 空白区切りでの読み込み(int変換)
var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}

func main() {
    sc.Split(bufio.ScanWords)
    n := nextInt()
    x := 0
    for i := 0; i < n; i++ {
        x += nextInt()
    }
    fmt.Println(x)
}
*/
