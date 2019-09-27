package rgraph

import (
	"fmt"
	"math/rand"
	"time"
)

var u []int
var groups [][]int

type pair struct {
	x int
	y int
}

var pairs []pair

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
		groups = append(groups, u[i*d:i*d+d])
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
				return x, y
			} else {
				return y, x
			}
		}
	}
}

func ForTest() {
	newGroups(4, 2)
	fmt.Println(twoRandomPoints())
}
