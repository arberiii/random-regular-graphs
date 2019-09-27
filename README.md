# Generating random regular graphs quickly
A regular graph is a graph where each vertex has the same number of neighbors. There are a lot of algorithms known for generating random regular graphs, but none of them is of practical. In paper "Generating random regular graphs quickly" by A. Steger and N. C. Wormald examined an algorithm which, although it does not generate uniformly at random, is provably close to a uniform generator when the degrees are relatively small.

# Algorithm
The configuration or pairing model of random d-regular graphs is as follows. Start with nd points (nd even) in n groups, and choose a random pairing of the points. Then create a graph with an edge from i to j if there is a pair containing points in the ith and jth groups. If no duplicated edge or loop (i.e., a pair of points in the same group) occurs the resulting d-regular graphs occur uniformly at random.

# Usage
```go
package main

import (
	"fmt"

	rgraph "github.com/arberiii/random-regular-graphs"
)

func main() {
	g := rgraph.RandomRegularGraph(5, 2)
	fmt.Println(g)
}
```

```
> go run main.go
map[0:[3 2] 1:[3 4] 2:[0 4] 3:[0 1] 4:[1 2]]
```
<img src="https://github.com/arberiii/random-regular-graphs/blob/master/star.png" class="center" width="200" height="200">
