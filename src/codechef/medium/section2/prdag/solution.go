package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	mat := make([][]byte, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		mat[i] = scanner.Bytes()
	}

	res := solve(n, mat)
	fmt.Println(len(res))

	for _, ans := range res {
		fmt.Printf("%d %d\n", ans[0], ans[1])
	}
}

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

func solve(n int, mat [][]byte) [][]int {
	outDegrees := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == '1' {
				outDegrees[i]++
			}
		}
	}
	used := make([]bool, n)
	newEdge := make([][]bool, n)
	for i := 0; i < n; i++ {
		newEdge[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		v := -1
		for j := n - 1; j >= 0; j-- {
			if !used[j] && outDegrees[j] == 0 {
				v = j
				break
			}
		}
		if v < 0 {
			break
		}
		used[v] = true
		for j := 0; j < n; j++ {
			if !used[j] && mat[j][v] == '0' {
				newEdge[j][v] = true
			}
		}
		for j := 0; j < n; j++ {
			if mat[j][v] == '1' {
				outDegrees[j]--
			}
		}
	}

	res := make([][]int, 0, n*n/2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if newEdge[i][j] {
				res = append(res, []int{i + 1, j + 1})
			}
		}
	}
	return res
}
