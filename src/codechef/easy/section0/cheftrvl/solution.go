package main

import (
	"bufio"
	"os"
	"fmt"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
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

func readNNums(scanner *bufio.Scanner, n int) []int {
	scanner.Scan()
	bs := scanner.Bytes()
	res := make([]int, n)
	x := -1
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		pairs := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			pairs[i] = readNNums(scanner, 2)
		}
		A := readNNums(scanner, n)
		fmt.Println(solve(n, pairs, A))
		t--
	}
}

func solve(n int, pairs [][]int, A []int) int64 {
	neighbors := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		neighbors[i] = make(map[int]bool)
	}

	for i := 0; i < len(pairs); i++ {
		a, b := pairs[i][0]-1, pairs[i][1]-1
		neighbors[a][b] = true
		neighbors[b][a] = true
	}

	cnt := make([][]int64, 2)
	cnt[0] = make([]int64, 2)
	cnt[1] = make([]int64, 2)

	B := make([]int, n)

	for i := 0; i < n; i++ {
		a := A[i]
		var c int
		for a > 0 {
			c++
			a -= a & -a
		}
		B[i] = c & 1
	}

	var dfs func(p, u int, path int)

	dfs = func(p, u int, path int) {
		cnt[B[u]][path&1]++
		for v := range neighbors[u] {
			if p == v {
				continue
			}
			dfs(u, v, path+1)
		}
	}

	dfs(-1, 0, 0)
	var ans int64

	for a := 0; a < 2; a++ {
		for b := 0; b < 2; b++ {
			for c := 0; c < 2; c++ {
				for d := 0; d < 2; d++ {
					if a^b != c^d {
						ans += cnt[a][b] * cnt[c][d]
					}
				}
			}
		}
	}

	return ans / 2
}
