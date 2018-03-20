package main

import (
	"bufio"
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
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	x = readInt(scanner.Bytes(), x+1, &b)
	readInt(scanner.Bytes(), x+1, &c)
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

func processNNums(bs []byte, n int) []int {
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
		n, m, K := readThreeNums(scanner)
		A := readNNums(scanner, n)

		inc := make([][]int, 0, m)
		desc := make([][]int, 0, m)

		for i := 0; i < m; i++ {
			scanner.Scan()
			bs := scanner.Bytes()
			tmp := processNNums(bs[2:], 2)
			if bs[0] == 'I' {
				inc = append(inc, tmp)
			} else {
				desc = append(desc, tmp)
			}
		}

		fmt.Println(solve(n, A, K, inc, desc))
		t--
	}
}

const MOD = 10e9 + 7

func solve(n int, A []int, K int, inc [][]int, desc [][]int) int {
	//fmt.Printf("[debug] %d %d %v %v %v\n", n, K, A, inc, desc)
	B, C := make([]int, n+1), make([]int, n+1)
	for i := 0; i < len(inc); i++ {
		a, b := inc[i][0], inc[i][1]
		B[a]++
		B[b]--
	}

	for i := 0; i < len(desc); i++ {
		a, b := desc[i][0], desc[i][1]
		C[a]++
		C[b]--
	}

	var sumb, sumc int

	for i := 0; i < n; i++ {
		sumb += B[i]
		B[i] = sumb
		sumc += C[i]
		C[i] = sumc
	}

	for i := 0; i < n; i++ {
		if B[i] > 0 && C[i] > 0 {
			return 0
		}
		if B[i] > 0 && A[i-1] >= 0 && A[i] >= 0 && A[i]-A[i-1] != 1 {
			return 0
		}
		if C[i] > 0 && A[i-1] >= 0 && A[i] >= 0 && A[i]-A[i-1] != -1 {
			return 0
		}
	}

	var dfs func(v int) int

	dfs = func(i int) int {
		if i == n {
			return 1
		}
		if i == n-1 {
			if A[i] == -1 {
				return K
			}
			return 1
		}

		if A[i] == -1 && B[i+1] == 0 && C[i+1] == 0 {
			return (K * dfs(i+1)) % MOD
		}
		if B[i+1] == 0 && C[i+1] == 0 {
			return dfs(i + 1)
		}

		j := i + 1
		for j < n && (B[j] != 0 || C[j] != 0) {
			j++
		}

		l := i
		for l < j && A[l] == -1 {
			l++
		}
		if l < j {
			for m := l - 1; m >= i; m-- {
				if A[m] == -1 {
					if B[m+1] > 0 {
						A[m] = A[m+1] - 1
					} else {
						A[m] = A[m+1] + 1
					}
				}
			}
			for m := l + 1; m < j; m++ {
				if A[m] == -1 {
					if B[m+1] > 0 {
						A[m] = A[m-1] + 1
					} else {
						A[m] = A[m-1] - 1
					}
				}
			}
			for m := i; m < j; m++ {
				if A[m] < 1 || A[m] > K {
					return 0
				}
				if B[m] > 0 && A[m]-A[m-1] != 1 {
					return 0
				}
				if C[m] > 0 && A[m]-A[m-1] != -1 {
					return 0
				}
			}

			return dfs(j)
		}

		d := make([]int, j-i)
		d[0] = 0
		for m := i + 1; m < j; m++ {
			if B[m] > 0 {
				d[m-i] = d[m-i-1] + 1
			} else {
				d[m-i] = d[m-i-1] - 1
			}
		}
		max, min := 0, 0
		for m := i + 1; m < j; m++ {
			if max < d[m-i] {
				max = d[m-i]
			}
			if min > d[m-i] {
				min = d[m-i]
			}
		}
		x, y := 1-min, K-max
		if y >= x {
			return ((y - x + 1) * dfs(j)) % MOD
		}
		return 0
	}

	return dfs(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
