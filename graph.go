package rgraph

type graph map[int][]int

func buildGraph() graph {
	g := make(graph)
	for i := range groups {
		g[i] = []int{}
	}
	for _, v := range pairs {
		g[v.x] = append(g[v.x], v.y)
		g[v.y] = append(g[v.y], v.x)
	}
	return g
}

func (g graph) isDRegular(d int) bool {
	for k := range g {
		if len(g[k]) != d {
			return false
		}
	}
	return true
}
