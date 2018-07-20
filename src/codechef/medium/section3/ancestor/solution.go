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
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		n := readNum(scanner)
		A := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			A[i] = readNNums(scanner, 2)
		}
		B := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			B[i] = readNNums(scanner, 2)
		}

		res := solve(n, A, B)

		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}

		tc--
	}
}
func solve(n int, A [][]int, B [][]int) []int {
	outA := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		outA[i] = make(map[int]bool)
	}

	for _, a := range A {
		u, v := a[0]-1, a[1]-1
		outA[u][v] = true
		outA[v][u] = true
	}

	var dfs func(p, u int, time *int)
	start := make([]int, n)
	end := make([]int, n)
	time := new(int)
	dfs = func(p, u int, time *int) {
		start[u] = *time
		*time++

		for v := range outA[u] {
			if p == v {
				continue
			}
			dfs(u, v, time)
		}
		end[u] = *time
	}

	dfs(0, 0, time)

	outB := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		outB[i] = make(map[int]bool)
	}

	for _, b := range B {
		u, v := b[0]-1, b[1]-1
		outB[u][v] = true
		outB[v][u] = true
	}

	res := make([]int, n)

	var dfs2 func(p, u int)
	bit := CreateBIT(*time + 1)

	dfs2 = func(p, u int) {
		res[u] = bit.Get(start[u])
		bit.Set(start[u], 1)
		bit.Set(end[u]+1, -1)
		for v := range outB[u] {
			if p == v {
				continue
			}
			dfs2(u, v)
		}
		bit.Set(start[u], -1)
		bit.Set(end[u]+1, 1)
	}

	dfs2(0, 0)

	return res
}

type BIT struct {
	array []int
	size  int
}

func CreateBIT(n int) BIT {
	array := make([]int, n+1)
	return BIT{array, n}
}

func (bit *BIT) Set(pos int, val int) {
	array := bit.array
	pos++
	for pos <= bit.size {
		array[pos] += val
		pos += pos & (-pos)
	}
}

func (bit BIT) Get(pos int) int {
	var res int
	pos++
	for pos > 0 {
		res += bit.array[pos]
		pos -= pos & (-pos)
	}
	return res
}

func solve1(n int, A [][]int, B [][]int) []int {
	outA := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		outA[i] = make(map[int]bool)
	}

	for _, a := range A {
		u, v := a[0]-1, a[1]-1
		outA[u][v] = true
		outA[v][u] = true
	}

	var dfs func(p, u int, time *int)
	start := make([]int, n)
	end := make([]int, n)

	dfs = func(p, u int, time *int) {
		start[u] = *time
		*time++

		for v := range outA[u] {
			if p == v {
				continue
			}
			dfs(u, v, time)
		}
		end[u] = *time
	}

	dfs(0, 0, new(int))

	isAncestorInTreeA := func(p, u int) bool {
		return start[p] <= start[u] && end[u] <= end[p]
	}

	res := make([]int, n)

	outB := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		outB[i] = make(map[int]bool)
	}

	for _, b := range B {
		u, v := b[0]-1, b[1]-1
		outB[u][v] = true
		outB[v][u] = true
	}

	PP := make([][]int, n)
	PB := make([][]bool, n)
	for i := 0; i < n; i++ {
		PP[i] = make([]int, 20)
		PB[i] = make([]bool, 20)
	}

	LCA := func(u int) int {
		for i := 19; i >= 0; i-- {
			if !PB[u][i] {
				u = PP[u][i]
			}
		}
		return PP[u][0]
	}

	var dfs2 func(p, u int)

	dfs2 = func(p, u int) {
		if u > 0 {
			PP[u][0] = p
			for i := 1; i < 20; i++ {
				PP[u][i] = PP[PP[u][i-1]][i-1]
			}

			PB[u][0] = isAncestorInTreeA(p, u)
			for i := 1; i < 20; i++ {
				PB[u][i] = PB[u][i-1] || PB[PP[u][i-1]][i-1]
			}
		}
		if u > 0 {
			cp := LCA(u)
			res[u] = res[cp] + 1
		}

		for v := range outB[u] {
			if p == v {
				continue
			}
			dfs2(u, v)
		}
	}

	dfs2(0, 0)

	return res
}
