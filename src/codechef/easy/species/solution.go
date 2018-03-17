package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
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

const MOD = 1e9 + 7

var pow3 []int
var pow2 []int

func init() {
	N := 25 * 50
	pow3 = make([]int, N+1)
	pow2 = make([]int, N+1)
	pow3[0] = 1
	pow2[0] = 1
	for i := 1; i <= N; i++ {
		pow3[i] = int((int64(pow3[i-1]) * 3) % MOD)
		pow2[i] = int((int64(pow2[i-1]) * 2) % MOD)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		grid := make([][]byte, n)

		for i := 0; i < n; i++ {
			scanner.Scan()
			grid[i] = scanner.Bytes()
		}
		fmt.Println(solve(n, grid))
		t--
	}
}

func solve(n int, grid [][]byte) int {
	dd := []int{-1, 0, 1, 0, -1}

	que := make([]int, n*n)

	bfs := func(a, b int, c byte) int {
		head, tail := 0, 0
		que[tail] = a*n + b
		tail++
		for head < tail {
			tt := tail
			for head < tt {
				v := que[head]
				head++
				i, j := v/n, v%n
				for k := 0; k < 4; k++ {
					x, y := i+dd[k], j+dd[k+1]
					if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == '?' {
						grid[x][y] = c
						que[tail] = x*n + y
						tail++
					}
				}
			}
		}

		return tail
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'B' || grid[i][j] == 'P' {
				bfs(i, j, grid[i][j])
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '.' && grid[i][j] != '?' {
				for k := 0; k < 4; k++ {
					x, y := i+dd[k], j+dd[k+1]
					if x >= 0 && x < n && y >= 0 && y < n && fight(grid[i][j], grid[x][y]) {
						return 0
					}
				}
			}
		}
	}

	var ones int
	var others int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '?' {
				grid[i][j] = '!'
				cnt := bfs(i, j, '!')
				if cnt == 1 {
					ones++
				} else {
					others++
				}
			}
		}
	}

	res1 := pow3[ones]
	res2 := pow2[others]

	return int((int64(res1) * int64(res2)) % MOD)
}

func fight(a, b byte) bool {
	if b == '.' {
		return false
	}
	if a == 'G' {
		return true
	}

	if b == '?' || b == a {
		return false
	}

	return true
}
