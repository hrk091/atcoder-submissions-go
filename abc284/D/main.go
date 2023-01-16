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
	t := atoi(sc.Text())
	ns := make([]int, t+1)
	for i := 1; i <= t; i++ {
		sc.Scan()
		ns[i] = atoi(sc.Text())
	}

	max := int(math.Pow(9*10e19, 1.0/3)) + 1
	if debug > 0 {
		fmt.Printf("max: %+v\n", max)
	}
	deleted := make([]bool, max+1)

	deleted[0] = true
	deleted[1] = true
	for i := 2; i <= max; i++ {
		if deleted[i] == true {
			continue
		}
		for j := i * 2; j <= max; j += i {
			deleted[j] = true
		}
	}

	var ps []int
	for i, b := range deleted {
		if !b {
			ps = append(ps, i)
		}
	}
	if debug > 0 {
		fmt.Printf("ps: %+v\n", ps[:100])
	}

	for i := 1; i <= t; i++ {
		n := ns[i]
		for _, j := range ps {
			if n%j != 0 {
				continue
			}
			var p, q int
			if n%(j*j) == 0 {
				p = j
				q = n / p / p
			} else {
				q = j
				p = int(math.Sqrt(float64(n / q)))
			}
			fmt.Printf("%d %d\n", p, q)
			break
		}
	}
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
