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

	us := make([]int, m+1)
	vs := make([]int, m+1)
	g := newGraph(n)
	for i := 1; i <= m; i++ {
		sc.Scan()
		us[i] = atoi(sc.Text())
		sc.Scan()
		vs[i] = atoi(sc.Text())

		u, v := us[i], vs[i]
		g.addEdge(u, v)
		g.addEdge(v, u)
	}
	if debug > 0 {
		fmt.Printf("graph: %+v\n", g.data)
	}

	size := 0
	visited := g.newVisited()

	for {
		var done bool
		for i, v := range visited {
			if !v {
				done, visited = g.dfs(i, visited)
				break
			}
		}
		size++
		if done {
			break
		}
	}
	fmt.Println(size)
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

func (g graph) addEdge(a, b int) {
	g.data[a] = append(g.data[a], b)
}

func (g graph) newVisited() []bool {
	visited := make([]bool, g.nodeSize+1)
	visited[0] = true
	return visited
}

// dfs conducts DFS and returns whether all nodes are visited or not and visited list.
func (g graph) dfs(pos int, visited []bool) (bool, []bool) {
	if visited == nil {
		visited = make([]bool, g.nodeSize+1)
		visited[0] = true
	}

	var dfs func(int)
	dfs = func(pos int) {
		visited[pos] = true
		for _, next := range g.data[pos] {
			if !visited[next] {
				dfs(next)
			}
		}
	}
	dfs(pos)

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

	return completed, visited
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	mustNil(err)
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func btoi(b byte) int {
	if b < '0' || '9' < b {
		panic(fmt.Errorf("cannot convert %s to int", []byte{b}))
	}
	return atoi(string(b))
}

func itob(i int) byte {
	if i < 0 || i > 9 {
		panic(fmt.Errorf("cannot convert %d to byte", i))
	}
	return byte(i + '0')
}

func mustNil(err error) {
	if err != nil {
		panic(err)
	}
}
