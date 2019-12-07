package main

import (
	"fmt"
)

// Edge desu
type Edge struct {
	from  int
	to    int
	index int
}

var color []int
var edge [][]Edge

//K from ABC146
var K int

func main() {
	var N int
	fmt.Scan(&N)
	var deg []int
	var av, bv int
	var i = 0
	for i < N {
		var array []Edge
		edge = append(edge, array)
		deg = append(deg, 0)
		i++
	}
	i = 0
	for i < N-1 {
		fmt.Scan(&av, &bv)
		av--
		bv--
		color = append(color, -1)
		edge[av] = append(edge[av], Edge{av, bv, i})
		edge[bv] = append(edge[bv], Edge{bv, av, i})
		deg[av]++
		deg[bv]++
		i++
	}
	start := 0
	i = 0
	for i < N {
		if deg[i] == 1 {
			start = i
			break
		}
		i++
	}
	K = 0
	dfs(-1, 0, start)
	fmt.Println(K)
	for _, c := range color {
		fmt.Println(c)
	}
}

func dfs(from int, beforeColor int, at int) {
	edges := edge[at]
	var useColor = 1
	for _, e := range edges {
		to := e.to
		if to == from {
			continue
		}
		if useColor == beforeColor {
			useColor++
		}
		color[e.index] = useColor //色確定
		dfs(at, useColor, to)
		K = max(K, useColor)
		useColor++
	}
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
