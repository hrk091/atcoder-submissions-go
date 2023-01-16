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
	ss := make([]string, n+1)

	for i := 1; i <= n; i++ {
		sc.Scan()
		ss[i] = sc.Text()
	}
	for i := n; i > 0; i-- {
		fmt.Println(ss[i])
	}
	// main
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
