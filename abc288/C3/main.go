package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	debug int
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	flag.Parse()
	d := flag.Arg(0)
	if d != "" {
		debug = atoi(d)
	}
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())
	sc.Scan()
	m := atoi(sc.Text())

	g := newGraph(n)
	for i := 1; i <= m; i++ {
		sc.Scan()
		a := atoi(sc.Text())
		sc.Scan()
		b := atoi(sc.Text())

		g.addEdge(a, b)
		g.addEdge(b, a)
	}

	visited := g.newVisited()
	ans := 0
	for i := 1; i < len(visited); i++ {
		if visited[i] {
			continue
		}

		if debug > 0 {
			fmt.Printf("### visited: %+v\n", visited)
			fmt.Printf("### start: %+v\n", i)
		}

		g.dfs(i, visited, func(curr, next int) {
			if visited[next] {
				ans++
			}
		})
	}
	fmt.Println(ans / 2)

}

type graph struct {
	nodeSize int
	data     map[int][]int
}

func newGraph(nodeSize int) *graph {
	data := make(map[int][]int, nodeSize+1)
	return &graph{
		nodeSize: nodeSize,
		data:     data,
	}
}

func (g *graph) addEdge(a, b int) {
	g.data[a] = append(g.data[a], b)
}

func (g *graph) newVisited() []bool {
	visited := make([]bool, g.nodeSize+1)
	visited[0] = true
	return visited
}

func (g *graph) isCompleted(visited []bool) bool {
	completed := true
	for _, v := range visited {
		if !v {
			completed = false
			break
		}
	}
	if debug > 0 {
		var visitedP []int
		for i, v := range visited {
			if i != 0 && v {
				visitedP = append(visitedP, i)
			}
		}
		fmt.Printf("visited: %+v\n", visitedP)
		fmt.Printf("completed: %+v\n", completed)
	}
	return completed
}

// dfs conducts DFS and returns whether all nodes are visited or not and visited list.
func (g *graph) dfs(pos int, visited []bool, fn func(curr, next int)) {

	var dfs func(int, int)
	dfs = func(prev, curr int) {
		visited[curr] = true

		for _, next := range g.data[curr] {
			if next == prev {
				// skip going back
				continue
			}
			if fn != nil {
				fn(curr, next)
			}
			if !visited[next] {
				dfs(curr, next)
			}
		}
		// if revisit is needed, enable following
		//visited[curr] = false
	}
	if !visited[pos] {
		dfs(0, pos)
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	mustNil(err)
	return i
}

func mustNil(err error) {
	if err != nil {
		panic(err)
	}
}
