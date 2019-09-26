package rgraph

import (
	"math/rand"
	"time"
)

var U []int
var Groups [][]int

type pair struct {
	x int
	y int
}

var pairs []pair

// should be ordered
func NewU(n, d int) {
	for i := 1; i < n*d+1; i++ {
		U = append(U, i)
	}
}

// should be ordered
func NewGroups(n, d int) {
	NewU(n, d)
	for i := 0; i < n; i++ {
		Groups = append(Groups, U[i*d:i*d+d])
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
	for i := range Groups {
		for j := range Groups[i] {
			if x == Groups[i][j] {
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
func TwoRandomPoints() (int, int) {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(len(U))
	for {
		y := rand.Intn(len(U))
		if x != y {
			if x < y {
				return x, y
			} else {
				return y, x
			}
		}
	}
}
