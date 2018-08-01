package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		H, W := readTwoNums(scanner)
		M := make([][]int, H)
		for i := 0; i < H; i++ {
			M[i] = readNNums(scanner, W)
		}
		res := solve(H, W, M)
		fmt.Printf("Case #%d:\n", x)
		for i := 0; i < H; i++ {
			var s bytes.Buffer
			for j := 0; j < W; j++ {
				s.WriteByte(res[i][j])
				if j < W-1 {
					s.WriteByte(' ')
				} else {
					s.WriteByte('\n')
				}
			}
			fmt.Print(s.String())
		}
	}
}

var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}

func solve(H, W int, M [][]int) [][]byte {
	res := make([][]byte, H)
	for i := 0; i < H; i++ {
		res[i] = make([]byte, W)
		for j := 0; j < W; j++ {
			res[i][j] = ' '
		}
	}

	valid := func(x, y int) bool {
		return x >= 0 && x < H && y >= 0 && y < W
	}

	var dfs func(x, y int, cur byte) byte

	dfs = func(x, y int, cur byte) byte {
		if res[x][y] != ' ' {
			return res[x][y]
		}
		res[x][y] = cur
		a, b := -1, -1
		for k := 0; k < 4; k++ {
			u, v := x+dx[k], y+dy[k]
			if valid(u, v) && (a < 0 || M[a][b] > M[u][v]) {
				a, b = u, v
			}
		}
		if valid(a, b) && M[a][b] < M[x][y] {
			res[x][y] = dfs(a, b, cur)
		}
		return res[x][y]
	}

	var cur byte = 'a'
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if res[i][j] == ' ' {
				tmp := dfs(i, j, cur)
				if tmp == cur {
					cur++
				}
			}
		}
	}

	return res
}
