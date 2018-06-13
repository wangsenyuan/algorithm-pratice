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

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &n)

	matrix := make([][]byte, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]byte, n)
		scanner.Scan()
		bs := scanner.Bytes()

		for j, k := 0, 0; j < len(bs); j, k = j+2, k+1 {
			matrix[i][k] = bs[j]
		}
	}
	scanner.Scan()
	s := scanner.Bytes()
	res := solve(n, matrix, s)

	fmt.Println(res)
}

var dd = [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

func solve(n int, matrix [][]byte, s []byte) int {
	f := make([][][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = make([]int, 2)
		}
	}

	p := 0
	for l := 1; l <= len(s); l++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[i][j][1-p] = 0
				if matrix[i][j] == s[l-1] {
					if l == 1 {
						f[i][j][1-p] = 1
					} else {
						for k := 0; k < len(dd); k++ {
							x, y := i+dd[k][0], j+dd[k][1]
							if x >= 0 && x < n && y >= 0 && y < n {
								f[i][j][1-p] += f[x][y][p]
							}
						}
					}
				}
			}
		}
		p = 1 - p
	}

	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += f[i][j][p]
		}
	}
	return ans
}
