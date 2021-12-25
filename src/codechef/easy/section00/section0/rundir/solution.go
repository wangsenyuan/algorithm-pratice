package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		X := make([]int, n)
		V := make([]int, n)
		for i := 0; i < n; i++ {
			a, b := readTwoNums(scanner)
			X[i] = a
			V[i] = b
		}
		res := solve(n, X, V)
		if res < INF {
			fmt.Printf("%.8f\n", res)
		} else {
			fmt.Println(-1)
		}
		t--
	}
}

const INF = 1 << 30

func time(s, v int) float64 {
	if v <= 0 {
		return INF
	}
	return float64(s) / float64(v)
}

func solve(n int, X []int, V []int) float64 {
	if n == 1 {
		return -1
	}
	ps := make(PS, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{X[i], V[i]}
	}

	sort.Sort(ps)

	var left float64 = INF
	var right float64 = INF

	for i := 1; i < n; i++ {
		a := time(ps[i].x-ps[i-1].x, ps[i].v-ps[i-1].v)
		// i goes left, i - 1 goes left
		a = math.Min(a, left)
		b := time(ps[i].x-ps[i-1].x, ps[i].v+ps[i-1].v)
		// i goes left, i - 1 goes right
		b = math.Min(b, right)
		x := math.Max(a, b)

		// i goes right, i - 1 goes left
		c := left
		// i goes right i - 1 goes right
		d := time(ps[i].x-ps[i-1].x, ps[i-1].v-ps[i].v)
		d = math.Min(d, right)
		y := math.Max(c, d)
		left, right = x, y
	}

	return math.Max(left, right)
}

type Pair struct {
	x, v int
}

type PS []Pair

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	return ps[i].x < ps[j].x
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
