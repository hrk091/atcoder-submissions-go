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
	for i := 0; i < m; i++ {
		sc.Scan()
		u := atoi(sc.Text())
		sc.Scan()
		v := atoi(sc.Text())
		g.addEdge(u, v)
		g.addEdge(v, u)
	}

	count, _ := g.dfs(1, nil)
	fmt.Printf("%d\n", count)
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

// dfs conducts DFS and returns whether all nodes are visited or not and visited list.
func (g *graph) dfs(pos int, visited []bool) (int, []bool) {
	if visited == nil {
		visited = g.newVisited()
	}

	var dfs func(int)
	count := 0
	dfs = func(pos int) {
		if count == 1e6 {
			return
		}
		count++
		visited[pos] = true
		for _, next := range g.data[pos] {
			if !visited[next] {
				dfs(next)
			}
		}
		visited[pos] = false
	}
	dfs(pos)

	return count, visited
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
