package lc_daily

func IsBipartite(graph [][]int) bool {
	n := len(graph)
	uf := NewUnionSet(n)
	for i:= 0; i < n; i++{
		for _, v := range graph[i] {
			if uf.IsConnected(i, v) {
				return false
			}
			uf.Union(graph[i][0], v)
		}
	}
	return true
}

type UnionSet struct{
	roots []int
	n 	int
}

func NewUnionSet(n int) *UnionSet {
	roots := make([]int, n)
	for i := 0; i < n; i++ {
		roots[i] = i
	}
	return &UnionSet{
		roots: roots,
		n: n,
	}
}

func (u *UnionSet) Find(x int) int {
	if u.roots[x] == x {
		return x
	}
	return u.Find(u.roots[x])
}

func (u *UnionSet) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}


func (u *UnionSet) Union(x, y int) {
	u.roots[u.Find(x)] = u.Find(y)
}
