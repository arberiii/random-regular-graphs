package rgraph

type Graph map[int][]int

func (g Graph) IsDRegular(d int) bool {
	for k := range g {
		if len(g[k]) != d {
			return false
		}
	}
	return true
}
