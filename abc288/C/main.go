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

		g.addEdge(a, b, i)
		g.addEdge(b, a, i)
	}

	edgeVisited := newVisited(m)
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

		g.dfs(i, visited, edgeVisited, func() {
			ans++
		})
	}
	fmt.Println(ans)
}

// ---------------------
// Graph

type edge struct {
	id     int
	farend int
}

type graph struct {
	nodeSize int
	data     map[int][]edge
}

func newGraph(nodeSize int) *graph {
	data := make(map[int][]edge, nodeSize+1)
	return &graph{
		nodeSize: nodeSize,
		data:     data,
	}
}

func (g *graph) addEdge(a, b, m int) {
	g.data[a] = append(g.data[a], edge{
		id:     m,
		farend: b,
	})
}

func newVisited(m int) []bool {
	visited := make([]bool, m+1)
	visited[0] = true
	return visited
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
		fmt.Printf("visited: %+v\n", visited)
		fmt.Printf("visited: %+v\n", visitedP)
		fmt.Printf("completed: %+v\n", completed)
	}
	return completed
}

// dfs conducts DFS and returns whether all nodes are visited or not and visited list.
func (g *graph) dfs(pos int, visited []bool, edgeVisited []bool, fn func()) {

	var dfs func(int)
	dfs = func(curr int) {
		visited[curr] = true

		for _, next := range g.data[curr] {
			if edgeVisited[next.id] {
				continue
			}
			edgeVisited[next.id] = true
			if fn != nil && visited[next.farend] {
				fn()
			}
			if !visited[next.farend] {
				dfs(next.farend)
			}
		}
		// if revisit is needed, enable following
		//visited[curr] = false
	}
	if !visited[pos] {
		dfs(pos)
	}
}

//func (g *graph) wfs(pos int, visited []bool, fn func(curr, next int)) (bool, []bool) {
//	if visited == nil {
//		visited = g.newVisited()
//	}
//
//	q := newQueue()
//	if !visited[pos] {
//		q.push(pos)
//	}
//
//	for !q.empty() {
//		curr := q.pop()
//		visited[curr] = true
//
//		for _, next := range g.data[curr] {
//			if fn != nil {
//				fn(curr, next)
//			}
//			if !visited[next] {
//				q.push(next)
//			}
//			// if multiple adding to queue is needed, disable following
//			visited[next] = true
//		}
//	}
//
//	return g.isCompleted(visited), visited
//}

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
