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
	_ = atoi(sc.Text())

}

// ---------------------
// Math

func abs(v int) int {
	return int(math.Abs(float64(v)))
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func max(values ...int) int {
	var mx int
	for _, v := range values {
		if v > mx {
			mx = v
		}
	}
	return mx
}

func min(values ...int) int {
	mn := math.MaxInt64
	for _, v := range values {
		if v < mn {
			mn = v
		}
	}
	return mn
}

func sum(values ...int) int {
	var a int
	for _, v := range values {
		a += v
	}
	return a
}

// ---------------------
// Graph

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
		dfs(-1, pos)
	}
}

func (g *graph) wfs(pos int, visited []bool, fn func(curr, next int)) {

	q := newQueue()
	if !visited[pos] {
		q.push(pos)
	}

	for !q.empty() {
		curr := q.pop()
		visited[curr] = true

		for _, next := range g.data[curr] {
			if fn != nil {
				fn(curr, next)
			}
			if !visited[next] {
				q.push(next)
			}
			// if multiple adding to queue is needed, disable following
			visited[next] = true
		}
	}
}

// ---------------------
// Union Find

type unionFind struct {
	parents []int
	sizes   []int
}

func newUnionFind(numNodes int) *unionFind {
	parents := make([]int, numNodes)
	fillSlice(parents, -1)
	sizes := make([]int, numNodes)
	fillSlice(sizes, 1)
	return &unionFind{parents, sizes}
}

func (uf *unionFind) root(v int) int {
	for uf.parents[v] != -1 {
		v = uf.parents[v]
	}
	return v
}

func (uf *unionFind) unite(u, v int) {
	rootU := uf.root(u)
	rootV := uf.root(v)
	if rootU == rootV {
		return
	}
	if uf.sizes[rootU] < uf.sizes[rootV] {
		uf.parents[rootU] = rootV
		uf.sizes[rootV] = uf.sizes[rootU] + uf.sizes[rootV]
	}
}

func (uf *unionFind) isSameGroup(u, v int) bool {
	return uf.root(u) == uf.root(v)
}

// ---------------------
// Queue

type queue struct {
	data []int
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) push(v int) {
	q.data = append(q.data, v)
}

func (q *queue) pop() int {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *queue) len() int {
	return len(q.data)
}

func (q *queue) empty() bool {
	return len(q.data) == 0
}

type pQueue struct {
	data    []int
	compare func(highP, lowP int) bool
}

func newPQueue(fn func(highP, lowP int) bool) *pQueue {
	return &pQueue{
		data:    []int{},
		compare: fn,
	}
}

func (q *pQueue) push(v int) {
	cur := len(q.data)
	q.data = append(q.data, v)
	for {
		if cur == 0 {
			break
		}
		next := (cur - 1) / 2
		if q.compare(q.data[cur], q.data[next]) {
			q.data[cur], q.data[next] = q.data[next], q.data[cur]
		} else {
			break
		}
		cur = next
	}
}

func (q *pQueue) pop() int {
	val := q.data[0]
	last := len(q.data) - 1
	q.data[0] = q.data[last]
	q.data = q.data[0:last]

	if last == 1 {
		return val
	}
	if last == 2 {
		if q.compare(q.data[1], q.data[0]) {
			q.data[0], q.data[1] = q.data[1], q.data[0]
		}
		return val
	}

	cur := 0
	for {
		l, r := cur*2+1, cur*2+2
		if r >= len(q.data) {
			break
		}
		var next int
		if q.compare(q.data[l], q.data[r]) {
			next = l
		} else {
			next = r
		}
		if q.compare(q.data[next], q.data[cur]) {
			q.data[cur], q.data[next] = q.data[next], q.data[cur]
		} else {
			break
		}
		cur = next
	}
	return val
}

func (q *pQueue) len() int {
	return len(q.data)
}

func (q *pQueue) empty() bool {
	return len(q.data) == 0
}

// ---------------------
// Segment Tree

type segmentTree struct {
	data   []int
	size   int
	eval   func(a, b int) int
	bottom int
}

func newSegmentTree(requiredSize int, bottom int, eval func(a, b int) int) *segmentTree {
	size := 1
	for size < requiredSize {
		size *= 2
	}
	data := make([]int, size*2)
	if bottom != 0 {
		for i, _ := range data {
			data[i] = bottom
		}
	}
	return &segmentTree{
		data:   data,
		size:   size,
		eval:   eval,
		bottom: bottom,
	}
}

func (s *segmentTree) update(pos, val int) {
	p := s.size + pos - 1
	s.data[p] = val
	for p > 1 {
		p /= 2
		s.data[p] = s.eval(s.data[p*2], s.data[p*2+1])
	}
}

func (s *segmentTree) query(l, r int) int {
	// 半開区間なので、 [1, size+1)
	return s._query(l, r, 1, s.size+1, 1)
}

func (s *segmentTree) _query(l, r, curL, curR, curN int) int {
	// 半開区間なので、端点が一致しても積は空集合
	if r <= curL || curR <= l {
		return s.bottom
	}
	if l <= curL && curR <= r {
		return s.data[curN]
	}
	m := (curL + curR) / 2
	ansL := s._query(l, r, curL, m, curN*2)
	ansR := s._query(l, r, m, curR, curN*2+1)
	return s.eval(ansL, ansR)
}

func (s *segmentTree) showDebug() {
	if debug == 0 {
		return
	}
	fmt.Printf("---")
	ypos := 0
	for i := 1; i <= s.size*2-1; i++ {
		if i >= pow(2, ypos) {
			ypos++
			fmt.Printf("\n%d: ", ypos)
		}
		fmt.Printf("%d ", s.data[i])
	}
	fmt.Println()
	fmt.Println("---")
}

// ---------------------
// LinkedList & OrderedMap

type linkedListNode struct {
	next *linkedListNode
	prev *linkedListNode

	key   int
	value interface{}
}

type linkedList struct {
	head *linkedListNode
	tail *linkedListNode
	len  int
}

func newLinkedList() *linkedList {
	list := &linkedList{
		head: &linkedListNode{value: nil},
		tail: &linkedListNode{value: nil},
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

func (l *linkedList) append(key int, value interface{}) *linkedListNode {
	n := &linkedListNode{
		prev:  l.tail.prev,
		next:  l.tail,
		key:   key,
		value: value,
	}

	l.tail.prev.next = n
	l.tail.prev = n
	l.len++

	return n
}

func (l *linkedList) remove(n *linkedListNode) bool {
	if n == l.head || n == l.tail {
		return false
	}

	n.prev.next = n.next
	n.next.prev = n.prev
	l.len--

	return true
}

func (l *linkedList) Iterate() chan *linkedListNode {
	ch := make(chan *linkedListNode)

	go func() {
		n := l.head

		for n.next != l.tail {
			ch <- n.next
			n = n.next
		}

		close(ch)
	}()

	return ch
}

type orderedDict struct {
	lookup map[int]*linkedListNode
	list   *linkedList
}

func newOrderedDict() *orderedDict {
	return &orderedDict{
		lookup: make(map[int]*linkedListNode),
		list:   newLinkedList(),
	}
}

func (d *orderedDict) set(key int, value interface{}) {
	if n, ok := d.lookup[key]; ok {
		d.list.remove(n)
	}

	d.lookup[key] = d.list.append(key, value)
}

func (d *orderedDict) get(key int) interface{} {
	if v, ok := d.lookup[key]; ok {
		return v.value
	} else {
		return -1
	}
}

func (d *orderedDict) remove(key int) bool {
	if n, ok := d.lookup[key]; ok {
		if ok := d.list.remove(n); !ok {
			return false
		}
		delete(d.lookup, key)
		return true
	}
	return false
}

func (d *orderedDict) removeLast() bool {
	n := d.list.head.next
	if ok := d.list.remove(n); !ok {
		return false
	}
	delete(d.lookup, n.key)
	return true
}

func (d *orderedDict) moveToTop(key int) bool {
	if n, ok := d.lookup[key]; ok {
		d.list.remove(n)
		d.lookup[key] = d.list.append(n.key, n.value)
		return true
	}
	return false
}

func (d *orderedDict) len() int {
	return d.list.len
}

func (d *orderedDict) iterate() chan interface{} {
	ch := make(chan interface{})

	go func() {
		for v := range d.list.Iterate() {
			ch <- v.value
		}

		close(ch)
	}()

	return ch
}

// ---------------------
// Util

func bitToList(b int, len int) ([]int, int) {
	// 63, 8 => [1,1,1,1,1,1,0,0]
	var ret []int
	count := 0
	for i := 0; i < len; i++ {
		v := b >> i & 1
		if v == 1 {
			count++
		}
		ret = append(ret, v)
	}
	return ret, count
}

func binarySearch(l, r int, fn func(int) int) int {
	// fn must be the one that returns true only when the result is greater than the given value.
	for l < r {
		mid := (l + r) / 2
		if ret := fn(mid); ret > 0 {
			l = mid + 1
		} else if ret < 0 {
			r = mid
		} else {
			return mid
		}
	}
	return l
}

func makeMatrix(n, m int) [][]int {
	matrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		matrix[i] = make([]int, m+1)
	}
	return matrix
}

func fillSlice(s []int, v int) {
	s[0] = v
	for p := 1; p < len(s); p *= 2 {
		copy(s[p:], s[:p])
	}
}

func fillMatrix(s [][]int, v int) {
	fillSlice(s[0], v)
	for p := 1; p < len(s); p++ {
		copy(s[p], s[0])
	}
}

func scanLineInt(sc *bufio.Scanner, size, offset int) []int {
	items := make([]int, size+offset)
	for i := 0; i < size; i++ {
		sc.Scan()
		items[i+offset] = atoi(sc.Text())
	}
	return items
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
