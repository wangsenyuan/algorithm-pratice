package main

import (
	"bufio"
	"fmt"
	"math"
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

	for i := 0; i < t; i++ {
		n, m := readTwoNums(scanner)
		C := readNNums(scanner, n)
		P := make([]int, m)
		meals := make([][]int, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			bs := scanner.Bytes()
			x := -1
			var p, q int
			x = readInt(bs, 0, &p)
			P[j] = p
			x = readInt(bs, x+1, &q)
			meals[j] = make([]int, q)
			for k := 0; k < q; k++ {
				x = readInt(bs, x+1, &meals[j][k])
			}
		}
		fmt.Println(solve(n, C, m, P, meals))
	}
}

func solve(n int, C []int, m int, P []int, meals [][]int) int64 {
	N := 1 << uint(n)
	dp1 := make([]int64, N)
	for mask := 0; mask < N; mask++ {
		dp1[mask] = math.MaxInt64
	}
	for i := 0; i < n; i++ {
		dp1[1<<uint(i)] = int64(C[i])
	}

	for i := 0; i < m; i++ {
		var mask int
		for j := 0; j < len(meals[i]); j++ {
			mask |= 1 << uint(meals[i][j]-1)
		}
		dp1[mask] = min(dp1[mask], int64(P[i]))
	}
	for mask := N - 1; mask >= 0; mask-- {
		for i := 0; i < n; i++ {
			if mask&(1<<uint(i)) > 0 {
				dp1[mask^(1<<uint(i))] = min(dp1[mask^(1<<uint(i))], dp1[mask])
			}
		}
	}
	dp2 := make([]int64, N)
	for i := 1; i < N; i++ {
		dp2[i] = math.MaxInt64
	}
	for mask := 0; mask < N; mask++ {
		subMask := mask & (mask - 1)
		for {
			dp2[mask] = min(dp2[mask], dp2[subMask]+dp1[mask^subMask])
			subMask = mask & (subMask - 1)
			if subMask == 0 {
				break
			}
		}
	}

	return dp2[N-1]
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
