package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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
	//scanner := bufio.NewScanner(os.Stdin)
	//tc := readNum(scanner)
	var tc int
	fmt.Scanf("%d", &tc)

	for x := 1; x <= tc; x++ {
		var N, M int
		fmt.Scanf("%d %d", &N, &M)
		G := make([]string, N)
		for i := 0; i < N; i++ {
			//scanner.Scan()
			fmt.Scanf("%s", &G[i])
		}
		//fmt.Fprintf(os.Stderr, "[debug] %d %d %v\n", N, M, G)

		res := solve(N, M, G)
		fmt.Printf("Case #%d: %d\n", x, res)
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(N int, M int, G []string) int {
	forest := make([]int, N*M)
	var idx int
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if G[i][j] == 'T' {
				forest[idx] = i*M + j
				idx++
			}
		}
	}

	que := make([]int, N*M)
	dist := make([]int, N*M)
	vis := make([]int, N*M)
	distance := func(x, y int, flag int) int {
		front, end := 0, 0
		que[end] = x
		end++
		vis[x] = flag
		dist[x] = 0
		for front < end {
			u := que[front]
			front++
			if u == y {
				return dist[u]
			}
			a, b := u/M, u%M
			for k := 0; k < 4; k++ {
				c, d := a+dd[k], b+dd[k+1]
				if c >= 0 && c < N && d >= 0 && d < M && G[c][d] != '.' && vis[c*M+d] != flag {
					dist[c*M+d] = dist[u] + 1
					vis[c*M+d] = flag
					que[end] = c*M + d
					end++
				}
			}
		}
		return -1
	}
	ff := make([]int, idx*idx)
	var flag int
	for i := 0; i < idx; i++ {
		for j := i + 1; j < idx; j++ {
			flag++
			x := distance(forest[i], forest[j], flag)
			y := x / 2
			c := x * (x + 1) / 2
			c -= y * (y + 1)
			if x&1 == 0 {
				c += y
			}
			ff[i*idx+j] = c
			ff[j*idx+i] = c
		}
	}

	mstOfForest := func() int {
		var res int
		dd := make([]int, idx)
		for i := 0; i < idx; i++ {
			dd[i] = math.MaxInt32
		}
		vis := make([]bool, idx)
		dd[0] = 0
		for {
			i := -1
			for j := 0; j < idx; j++ {
				if !vis[j] && (i == -1 || dd[j] < dd[i]) {
					i = j
				}
			}
			if i < 0 {
				return res
			}
			res += dd[i]
			vis[i] = true

			for j := 0; j < idx; j++ {
				if !vis[j] && dd[j] > ff[i*idx+j] {
					dd[j] = ff[i*idx+j]
				}
			}
		}
	}

	res := mstOfForest()

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if G[i][j] == '#' {
				d := -1
				for k := 0; k < idx; k++ {
					flag++
					x := distance(forest[k], i*M+j, flag)
					if d < 0 || x < d {
						d = x
					}
				}
				res += d
			}
		}
	}

	fmt.Fprintf(os.Stderr, "[debug]cost result %v\n", res)
	//fmt.Fprintf(os.Stderr, "[debug] %d %d => %d\n", N, M, ans)
	return res
}
