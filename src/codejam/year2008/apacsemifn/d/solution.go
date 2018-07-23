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

func readFloat64(bs []byte, from int, val *float64) int {
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}

	var dec float64
	for i < len(bs) && bs[i] != '.' && bs[i] != ' ' {
		dec = dec*10 + float64(bs[i]-'0')
		i++
	}
	*val = dec

	if i == len(bs) || bs[i] == ' ' {
		//no fraction
		return i
	}
	i++
	var frac float64
	base := 1.0
	for i < len(bs) && bs[i] != ' ' {
		frac = frac*10 + float64(bs[i]-'0')
		base *= 10
		i++
	}
	*val += frac / base
	return i * sign
}

func main() {

	f, err := os.Open("./D-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		N := readNum(scanner)

		Large := make([][]int, N-1)
		for i := 0; i < N-1; i++ {
			Large[i] = readNNums(scanner, 2)
		}

		M := readNum(scanner)
		Small := make([][]int, M-1)

		for i := 0; i < M-1; i++ {
			Small[i] = readNNums(scanner, 2)
		}
		res := solve(N, Large, M, Small)

		if !res {
			fmt.Printf("Case #%d: NO\n", x)
		} else {
			fmt.Printf("Case #%d: YES\n", x)
		}
	}
}

func solve(N int, Large [][]int, M int, Small [][]int) bool {
	//Fix Small rooted at 0
	// fmt.Printf("[debug]---%d %v %d %v\n", N, Large, M, Small)
	deg1 := degree(N, Large)
	deg2 := degree(M, Small)
	conn1 := connections(N, Large)
	conn2 := connections(M, Small)

	leaf := func(p, u int, deg []int) bool {
		if p >= 0 {
			return deg[u] == 1
		}
		return deg[u] == 0
	}

	isomorphism := func(r1 int) bool {
		vis := make([]int, M)
		var flag int

		var dfs2 func(u int, can map[int]map[int]bool, pair []int) bool

		dfs2 = func(u int, can map[int]map[int]bool, pair []int) bool {
			if vis[u] == flag {
				return false
			}
			vis[u] = flag

			for v := range can[u] {
				if pair[v] == -1 || dfs2(pair[v], can, pair) {
					pair[v] = u
					return true
				}
			}
			return false
		}
		var dfs1 func(p1, u1, p2, u2 int) bool

		dfs1 = func(p1, u1, p2, u2 int) bool {
			if leaf(p2, u2, deg2) {
				return true
			} else if leaf(p1, u1, deg1) {
				return false
			}

			if len(conn1[u1]) < len(conn2[u2]) {
				return false
			}
			can := make(map[int]map[int]bool)
			for _, v2 := range conn2[u2] {
				if p2 == v2 {
					continue
				}
				can[v2] = make(map[int]bool)
				for _, v1 := range conn1[u1] {
					if v1 == p1 {
						continue
					}
					if dfs1(u1, v1, u2, v2) {
						can[v2][v1] = true
					}
				}
				if len(can[v2]) == 0 {
					return false
				}
			}
			pair := make([]int, N)
			for i := 0; i < N; i++ {
				pair[i] = -1
			}

			for _, v2 := range conn2[u2] {
				if p2 == v2 {
					continue
				}
				flag++
				if !dfs2(v2, can, pair) {
					return false
				}
			}
			return true
		}

		var dfs3 func(p, u int) bool

		dfs3 = func(p, u int) bool {
			if dfs1(p, u, -1, 0) {
				return true
			}

			for _, v := range conn1[u] {
				if v == p {
					continue
				}
				if dfs3(u, v) {
					return true
				}
			}
			return false
		}

		return dfs3(-1, r1)
	}

	for i := 0; i < N; i++ {
		if isomorphism(i) {
			return true
		}
	}
	return false
}

func degree(N int, G [][]int) []int {
	deg := make([]int, N)
	for _, edge := range G {
		u, v := edge[0]-1, edge[1]-1
		deg[u]++
		deg[v]++
	}
	return deg
}

func connections(N int, G [][]int) [][]int {
	res := make([][]int, N)
	for i := 0; i < N; i++ {
		res[i] = make([]int, 0, 3)
	}

	for _, edge := range G {
		u, v := edge[0]-1, edge[1]-1
		res[u] = append(res[u], v)
		res[v] = append(res[v], u)
	}

	return res
}
