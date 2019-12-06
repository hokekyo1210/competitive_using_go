package main

import (
	"fmt"
	"strings"
)

func main() {
	var data string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var N int
	var S string
	var ans string = ""
	fmt.Scan(&N, &S)

	var i = 0
	for i < len(S) {
		var t = S[i : i+1]
		var at = strings.Index(data, t)
		var targetJ int = (at + N) % 26
		var target = data[targetJ : targetJ+1]
		ans = ans + target
		i++
	}
	fmt.Println(ans)
}
