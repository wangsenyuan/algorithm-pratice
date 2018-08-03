package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
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

	for tc > 0 {
		tc--
		scanner.Scan()
		bs := scanner.Bytes()
		var R, C int
		pos := readInt(bs, 0, &R)
		pos = readInt(bs, pos+1, &C)
		var G int64
		readInt64(bs, pos+1, &G)
		res := solve(R, C, G)
		fmt.Println(len(res))
		n := len(res)
		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
	}
}

const MAX_N = 100
const MAX_G = 1e12

var X [][]int64

func init() {
	X = make([][]int64, MAX_N)
	for i := 0; i < MAX_N; i++ {
		X[i] = make([]int64, MAX_N)
	}

	X[0][0] = 1
	for i := 1; i < MAX_N; i++ {
		X[i][0] = 1
		X[i][i] = 1
		for j := 1; j < i; j++ {
			if X[i-1][j]+X[i-1][j-1] > MAX_G {
				X[i][j] = MAX_G
			} else {
				X[i][j] = X[i-1][j] + X[i-1][j-1]
			}
		}
	}
	// fmt.Fprintf(os.Stderr, "[debug]%v\n", X[MAX_N])
}

func solve(R, C int, G int64) []int64 {
	res := make([]int64, 0, 10)
	for G > 0 {
		// X[r][c] == X[r][r-c]
		// if 2*C > R {
		// 	C = R - C
		// }
		i := sort.Search(MAX_N, func(i int) bool {
			return X[i][C] > G
		})
		i--
		res = append(res, X[i][C])
		G -= X[i][C]
		C--
	}

	return res
}
