package rgraph

import (
	"math/rand"
	"time"
)

const MaxPairIteration = 100000
const MaxGraphIteration = 100

var u []int
var groups [][]int

type pair struct {
	x int
	y int
}

var pairs []pair

// RandomRegularGraph is a practical algorithm for generating random regular graphs with n vertices of degree d.
func RandomRegularGraph(n, d int) graph {
	newGroups(n, d)
	for i := 0; i < MaxGraphIteration; i++ {
		g := iteration()
		if g.isDRegular(d) {
			return g
		}
	}
	return nil
}

// should be ordered
func newU(n, d int) {
	for i := 1; i < n*d+1; i++ {
		u = append(u, i)
	}
}

// should be ordered
func newGroups(n, d int) {
	newU(n, d)
	for i := 0; i < n; i++ {
		var group []int
		for _, v := range u[i*d : i*d+d] {
			group = append(group, v)
		}
		groups = append(groups, group)
	}
}

// x < y
func suitable(x, y int) bool {
	// check that  they lie in different group
	indexOfX := indexInGroups(x)
	indexOfY := indexInGroups(y)
	if indexOfX == indexOfY {
		return false
	}

	// check that  no currently existing pair contains points
	// in the same two group
	if pairExist(indexOfX, indexOfY) {
		return false
	}

	return true
}

func indexInGroups(x int) int {
	for i := range groups {
		for j := range groups[i] {
			if x == groups[i][j] {
				return i
			}
		}
	}
	return -1
}

func indexInU(x int) int {
	for i := range u {
		if u[i] == x {
			return i
		}
	}
	return -1
}

func pairExist(x, y int) bool {
	for i := range pairs {
		if pairs[i].x == x && pairs[i].y == y {
			return true
		}
	}
	return false
}

// order tuple
func twoRandomPoints() (int, int) {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(len(u))
	for {
		y := rand.Intn(len(u))
		if x != y {
			if x < y {
				return u[x], u[y]
			} else {
				return u[y], u[x]
			}
		}
	}
}

func pairNewPoints() {
	if len(u) > 0 {
		x, y := twoRandomPoints()
		if suitable(x, y) {
			pairs = append(pairs, pair{x: indexInGroups(x), y: indexInGroups(y)})
			deleteIndex(y)
			deleteIndex(x)
		}
	}
}

func deleteIndex(x int) {
	i := indexInU(x)
	u = append(u[:i], u[i+1:]...)
}

func iteration() graph {
	for i := 0; i < MaxPairIteration; i++ {
		pairNewPoints()
	}
	return buildGraph()
}
