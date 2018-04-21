package main

import (
	"bufio"
	"fmt"
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

	first := readNNums(scanner, 3)
	n, x, y := first[0], first[1], first[2]

	A := readNNums(scanner, n)
	B := readNNums(scanner, n)
	fmt.Println(solve(n, x, y, A, B))
}

func solve(n, x, y int, A []int, B []int) int {
	ps := make([]Pair, n)
	for i := 0; i < n; i++ {
		ps[i] = Pair{A[i], B[i]}
	}

	sort.Sort(PS(ps))

	var ans int

	for i := 0; i < n; i++ {
		p := ps[i]
		if p.a > p.b {
			if x > 0 {
				ans += p.a
				x--
			} else {
				ans += p.b
				y--
			}
		} else {
			if y > 0 {
				ans += p.b
				y--
			} else {
				ans += p.a
				x--
			}
		}
	}

	return ans
}

type Pair struct {
	a, b int
}

type PS []Pair

func (this PS) Len() int {
	return len(this)
}

func (this PS) Less(i, j int) bool {
	x := abs(this[i].a - this[i].b)
	y := abs(this[j].a - this[j].b)
	return x > y
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (this PS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
