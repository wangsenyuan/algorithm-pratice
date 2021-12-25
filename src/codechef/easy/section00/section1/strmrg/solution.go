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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n, m := readTwoNums(scanner)
		scanner.Scan()
		s1 := scanner.Text()
		scanner.Scan()
		s2 := scanner.Text()
		fmt.Println(solve(n, s1, m, s2))
		t--
	}
}

func solve(n int, s1 string, m int, s2 string) int {
	INF := 1 << 30

	// f[i][j][c] = mimmum value before (i, j) ending with c
	f := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			f[i][j] = make([]int, 2)
			f[i][j][0] = INF
			f[i][j][1] = INF
		}
	}

	f[1][0][0] = 1
	f[0][1][1] = 1
	S := []string{s1, s2}
	idx := []int{0, 0}
	sz := []int{n, m}
	for idx[0] = 0; idx[0] <= sz[0]; idx[0]++ {
		for idx[1] = 0; idx[1] <= sz[1]; idx[1]++ {
			for pz := 0; pz <= 1; pz++ {
				for nz := 0; nz <= 1; nz++ {
					if idx[nz] < sz[pz] && idx[pz] > 0 {
						ndx := []int{idx[0] + 1 - nz, idx[1] + nz}
						var tmp int
						if S[nz][idx[nz]] != S[pz][idx[pz]-1] {
							tmp = 1
						}
						f[ndx[0]][ndx[1]][nz] = min(f[ndx[0]][ndx[1]][nz], f[idx[0]][idx[1]][pz]+tmp)
					}
				}
			}
		}
	}
	return min(f[n][m][0], f[n][m][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
