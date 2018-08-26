package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, M := readTwoNums(scanner)
		plan := make([]string, N)
		output := make([]string, N)
		for i := 0; i < N; i++ {
			scanner.Scan()
			s := scanner.Text()
			ss := strings.Split(s, " ")
			plan[i] = ss[0]
			output[i] = ss[1]
		}
		fmt.Println(solve(N, M, plan, output))
	}
}

func solve(N, M int, plan []string, output []string) string {
	// check in valid cases

	for i := 0; i < N; i++ {
		if plan[i] == "correct" {
			for j := 0; j < M; j++ {
				if output[i][j] == '0' {
					// a fail case
					return "INVALID"
				}
			}
		}
	}
	// all cases are valid
	for i := 0; i < N; i++ {
		if plan[i] == "wrong" {
			pass := true
			for j := 0; j < M; j++ {
				if output[i][j] == '0' {
					pass = false
					break
				}
			}
			if pass {
				return "WEAK"
			}
		}
	}

	return "FINE"
}
