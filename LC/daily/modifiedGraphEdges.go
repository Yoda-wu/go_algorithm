package lc_daily

import (
	_ "fmt"

	"math"
)

func modifiedGraphEdges(n int, edges [][]int, source, destination, target int) [][]int {
	type edge struct{ to, eid int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], edge{y, i})
		g[y] = append(g[y], edge{x, i})
	}
	var delta int
	dis := make([][2]int, n)
	for i := range dis {
		dis[i] = [2]int{math.MaxInt, math.MaxInt}
	}
	dis[source] = [2]int{}
	disjkstra := func(k int) {
		vis := make([]bool, n)
		for {
			x := -1
			for y, b := range vis {
				if !b && (x < 0 || dis[y][k] < dis[x][k]) {
					x = y
				}
			}
			if x == destination {
				return
			}
			vis[x] = true
			for _, e := range g[x] {
				y, eid := e.to, e.eid
				wt := edges[eid][2]
				if wt == -1 {
					wt = 1
				}
				if k == 1 && edges[eid][2] == -1 {
					w := delta + dis[y][0] - dis[x][1]
					if w > wt {
						wt = w
						edges[eid][2] = w
					}
				}
				dis[y][k] = min(dis[y][k], dis[x][k]+wt)
			}
		}

	}
	disjkstra(0)
	delta = target - dis[destination][0]
	if delta < 0 {
		return nil
	}
	disjkstra(1)
	if dis[destination][1] < target {
		return nil
	}
	for _, e := range edges {
		if e[2] == -1 { // 剩余没修改的边全部改成 1
			e[2] = 1
		}
	}
	return edges
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
