package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var n, p int
		scanner.Scan()
		readInt(scanner.Bytes(), readInt(scanner.Bytes(), 0, &n)+1, &p)
		scanner.Scan()
		A := scanner.Text()
		scanner.Scan()
		B := scanner.Text()
		S, cnt := solve(n, A, B, p)
		if S == 0 {
			fmt.Println(-1)
		} else {
			fmt.Printf("%d %d\n", S, cnt)
		}
	}
}

func solve(n int, A string, B string, P int) (int, int) {
	f := make([][][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = make([]int, n)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				x := i*n*n + j*n + k
				if A[x] == B[x] {
					f[i][j][k] = 1
				}

				if i > 0 {
					f[i][j][k] += f[i-1][j][k]
				}
				if j > 0 {
					f[i][j][k] += f[i][j-1][k]
				}
				if k > 0 {
					f[i][j][k] += f[i][j][k-1]
				}
				if i > 0 && j > 0 {
					f[i][j][k] -= f[i-1][j-1][k]
				}
				if i > 0 && k > 0 {
					f[i][j][k] -= f[i-1][j][k-1]
				}
				if j > 0 && k > 0 {
					f[i][j][k] -= f[i][j-1][k-1]
				}
				if i > 0 && j > 0 && k > 0 {
					f[i][j][k] += 2 * f[i-1][j-1][k-1]
				}
			}
		}
	}

	S, cnt := 0, 0
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				for side := 1; side+a <= n && side+b <= n && side+c <= n; side++ {
					i := a + side - 1
					j := b + side - 1
					k := c + side - 1
					sz := side * side * side
					tmp := f[i][j][k]
					if a > 0 {
						tmp -= f[a-1][j][k]
					}
					if b > 0 {
						tmp -= f[i][b-1][k]
					}
					if c > 0 {
						tmp -= f[i][j][c-1]
					}
					if a > 0 && b > 0 {
						tmp += f[a-1][b-1][k]
					}
					if a > 0 && c > 0 {
						tmp += f[a-1][j][c-1]
					}
					if b > 0 && c > 0 {
						tmp += f[i][b-1][c-1]
					}
					if a > 0 && b > 0 && c > 0 {
						tmp -= 2 * f[a-1][b-1][c-1]
					}
					if tmp*100 >= P*sz {
						if side > S {
							S = side
							cnt = 1
						} else if side == S {
							cnt++
						}
					}
				}
			}
		}
	}
	return S, cnt
}
