package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1000

var C [][]int64

const MOD = 1000000080798150871

func init() {
	C = make([][]int64, N+1)
	for i := 0; i <= N; i++ {
		C[i] = make([]int64, N+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % MOD
		}
	}
}

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
	*val = sign * tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readFour(scanner *bufio.Scanner) (a int, b int, c int, d int) {
	scanner.Scan()
	bs := scanner.Bytes()
	x := readInt(bs, 0, &a)
	x = readInt(bs, x+1, &b)
	x = readInt(bs, x+1, &c)
	x = readInt(bs, x+1, &d)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for i := 0; i < t; i++ {
		n, m, l, k := readFour(scanner)
		grid := make([][]byte, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			grid[j] = scanner.Bytes()
		}
		fmt.Println(solve(n, m, grid, l, k))
	}
}

func solve(n, m int, grid [][]byte, l, k int) int64 {
	if l > n || k > m {
		return 0
	}
	count := func(cols int) int {
		var res int
		for i := 0; i < n; i++ {
			allOne := true

			for j := 0; j < m; j++ {
				if cols&(1<<uint(j)) > 0 {
					if grid[i][j] == '0' {
						allOne = false
						break
					}
				}
			}

			if allOne {
				res++
			}
		}
		return res
	}
	var ans int64
	for mask := 0; mask < (1 << uint(m)); mask++ {
		var cols int
		for j := 0; j < m; j++ {
			if mask&(1<<uint(j)) > 0 {
				cols++
			}
		}
		if cols != k {
			continue
		}
		rows := count(mask)
		ans = (ans + C[rows][l]) % MOD
	}

	return ans
}
