package main

import (
	"fmt"
)

func main() {
	var A, B, X int64
	fmt.Scan(&A, &B, &X)
	ans := solve(A, B, X)
	fmt.Println(ans)
}

func solve(A int64, B int64, X int64) int64 {

	// 範囲start < endを探索する
	var start int64 = 0
	var end int64 = 1000000000
	var N int64
	var ans int64 = 0
	for {
		if end < start {
			break
		}
		N = (start + end) / 2

		// fmt.Println(start, end, N)

		var dn int64 = 0
		var num int64 = N

		for num != 0 {
			num = num / 10
			dn++
		}

		if (A * N) <= (X - dn*B) {
			start = N + 1
			ans = N
		} else {
			end = N - 1
		}
	}
	return ans
}
