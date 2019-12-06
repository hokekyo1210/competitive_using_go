package main

import (
	"bufio" //入力用
	"fmt"   //出力用
	"os"    //入力用
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var S string
	if scanner.Scan() {
		S = scanner.Text()
	}
	fmt.Println(S)
}
