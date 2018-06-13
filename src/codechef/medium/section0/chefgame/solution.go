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
		N, D := readTwoNums(scanner)
		points := make([][]int, N)
		for i := 0; i < N; i++ {
			points[i] = readNNums(scanner, D)
		}
		fmt.Println(solve(N, D, points))
		t--
	}
}

const MOD = 747474747

func solve(N int, D int, points [][]int) int64 {
	// sort.Sort(Points(points))

	dist := make([]int64, N)
	dist[0] = 0

	for i := 1; i < N; i++ {
		dist[i] = distance(points[0], points[i])
	}

	var score int64 = 1

	seen := make([]bool, N)
	seen[0] = true

	for {
		u := -1
		for i := 0; i < N; i++ {
			if seen[i] {
				continue
			}
			if u == -1 || dist[u] < dist[i] {
				u = i
			}
		}
		if u < 0 || dist[u] == 0 {
			break
		}
		score *= (dist[u]) % MOD
		score %= MOD

		seen[u] = true
		for i := 0; i < N; i++ {
			if seen[i] {
				continue
			}
			tmp := distance(points[i], points[u])
			if tmp > dist[i] {
				dist[i] = tmp
			}
		}
	}

	return score
}

func distance(a, b []int) int64 {
	var res int64

	for i := 0; i < len(a); i++ {
		d := int64(a[i]) - int64(b[i])
		res += d * d
	}

	return res
}
