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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	firstLine := readNNums(scanner, 3)
	N := firstLine[0]
	M := firstLine[1]
	P := firstLine[2]
	songs := make([][]int, N)
	for i := 0; i < N; i++ {
		songs[i] = readNNums(scanner, 2)
	}
	albums := readNNums(scanner, M)
	fmt.Println(solve(N, M, P, songs, albums))
}

const MAXN = 5010

func solve(N, M, P int, songs [][]int, albums []int) int64 {
	av := make([]int, M)
	as := make([][]int, M)
	for i, song := range songs {
		a := song[0]
		v := song[2]
		av[a-1] += v
		if as[a-1] == nil {
			as[a-1] = make([]int, 0, 5)
		}
		as[a-1] = append(as[a-1], i)
	}
	dp := make([][]int64, MAXN)
	for i := 0; i < MAXN; i++ {
		dp[i] = make([]int64, P+1)
	}

	var idx int
	for i := 0; i < M; i++ {
		for _, j := range as[i] {
			idx++
			for price := 0; price <= P; price++ {
				dp[idx][price] = dp[idx-1][price]
				if price >= songs[j][1] && dp[idx-1][price-songs[j][1]]+int64(songs[j][2]) > dp[idx][price] {
					dp[idx][price] = dp[idx-1][price-songs[j][1]] + int64(songs[j][2])
				}
			}
		}
		ap := albums[i]
		idx++
		for price := 0; price <= P; price++ {
			dp[idx][price] = dp[idx-1][price]
			if price >= ap {
				if dp[idx-len(as)-1][price-ap]+int64(av[i]) > dp[idx][price] {
					dp[idx][price] = dp[idx-len(as)-1][price-ap] + int64(av[i])
				}
			}
		}
	}

	return dp[idx][P]
}
