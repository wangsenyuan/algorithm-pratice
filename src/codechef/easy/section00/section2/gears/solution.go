package main

import (
	"bufio"
	"fmt"
	"os"
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
	n, m := readTwoNums(scanner)
	A := readNNums(scanner, n)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		bs := scanner.Bytes()
		var kind int
		pos := readInt(bs, 0, &kind)
		if kind == 1 || kind == 2 {
			queries[i] = make([]int, 3)
			queries[i][0] = kind
			pos = readInt(bs, pos+1, &queries[i][1])
			readInt(bs, pos+1, &queries[i][2])
		} else {
			queries[i] = make([]int, 4)
			queries[i][0] = kind
			pos = readInt(bs, pos+1, &queries[i][1])
			pos = readInt(bs, pos+1, &queries[i][2])
			pos = readInt(bs, pos+1, &queries[i][3])
		}
	}
	res := solve(n, m, A, queries)
	for _, ans := range res {
		if ans[0] == 0 {
			fmt.Println(0)
		} else {
			fmt.Printf("%d/%d\n", ans[0], ans[1])
		}
	}
}

func solve(n int, m int, A []int, queries [][]int) [][]int {
	uf := CreateUF(n)
	res := make([][]int, m)
	colors := make([]int, n)
	for i := 0; i < n; i++ {
		colors[i] = 1
	}
	blocked := make([]bool, n)
	var j int

	connect := func(x, y int) {
		px := uf.Find(x)
		py := uf.Find(y)

		if px != py {
			if colors[x] == colors[y] {
				if uf.Count(px) < uf.Count(py) {
					px, py = py, px
				}
				elems := uf.Elems(py)
				for _, i := range elems {
					colors[i] = -colors[i]
				}
			}
			uf.Union(x, y)
			blocked[uf.Find(x)] = blocked[px] || blocked[py]
			return
		}

		uf.Union(x, y)

		if colors[x] == colors[y] {
			// a block situation
			blocked[uf.Find(x)] = true
		}
	}

	for i := 0; i < m; i++ {
		query := queries[i]
		if query[0] == 1 {
			x, c := query[1]-1, query[2]
			A[x] = c
		} else if query[0] == 2 {
			x, y := query[1]-1, query[2]-1
			connect(x, y)
		} else {
			x, y, v := query[1]-1, query[2]-1, query[3]
			px, py := uf.Find(x), uf.Find(y)
			if px != py {
				// not connected
				res[j] = []int{0, 0}
				j++
			} else {
				// connected
				if blocked[px] {
					res[j] = []int{0, 0}
					j++
				} else {
					sign := colors[x] * colors[y]
					res[j] = calculate(A[x], A[y], v, sign)
					j++
				}
			}
		}
	}

	return res[:j]
}

func calculate(a, b, v, sign int) []int {
	g := gcd(a*v, b)
	return []int{sign * a * v / g, b / g}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type UF struct {
	set   []int
	elems [][]int
	count []int
}

func CreateUF(n int) UF {
	set := make([]int, n)
	count := make([]int, n)
	elems := make([][]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
		count[i] = 1
		elems[i] = make([]int, 0, 10)
		elems[i] = append(elems[i], i)
	}
	return UF{set, elems, count}
}

func (uf *UF) Find(x int) int {
	set := uf.set
	var loop func(x int) int
	loop = func(x int) int {
		if set[x] != x {
			set[x] = loop(set[x])
		}
		return set[x]
	}

	return loop(x)
}

func (uf *UF) Count(x int) int {
	return uf.count[uf.Find(x)]
}

func (uf *UF) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}
	if uf.count[pa] < uf.count[pb] {
		uf.set[pa] = pb
		uf.count[pb] += uf.count[pa]
		uf.elems[pb] = append(uf.elems[pb], uf.elems[pa]...)
		uf.elems[pa] = nil
	} else {
		uf.set[pb] = pa
		uf.count[pa] += uf.count[pb]
		uf.elems[pa] = append(uf.elems[pa], uf.elems[pb]...)
		uf.elems[pb] = nil
	}
	return true
}

func (uf *UF) Elems(x int) []int {
	return uf.elems[uf.Find(x)]
}
