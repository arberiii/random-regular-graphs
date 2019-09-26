package rgraph

type U []int
type Groups [][]int

func NewU(n, d int) U {
	var ret []int
	for i := 1; i < n*d+1; i++ {
		ret = append(ret, i)
	}
	return ret
}

func NewGroups(n, d int) Groups {
	var ret [][]int
	u := NewU(n, d)
	for i := 0; i < n; i++ {
		ret = append(ret, u[i*d:i*d+d])
	}
	return ret
}
