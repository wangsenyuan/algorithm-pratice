package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			bs := scanner.Bytes()
			var m int
			pos := readInt(bs, 0, &m)
			A[i] = make([]int, m)
			for j := 0; j < m; j++ {
				pos = readInt(bs, pos+1, &A[i][j])
			}
		}
		res := solve(n, A)
		if res > 0 {
			fmt.Println("EVEN")
		} else if res < 0 {
			fmt.Println("ODD")
		} else {
			fmt.Println("DON'T PLAY")
		}
	}
}

func solve(n int, A [][]int) int {
	var res int64

	for i := 0; i < n; i++ {
		res += calculate(A[i])
	}
	if res > 0 {
		return 1
	} else if res < 0 {
		return -1
	}
	return 0
}

const INF = 1 << 48

func calculate(A []int) int64 {
	sort.Ints(A)
	var j int

	for i := 1; i <= len(A); i++ {
		if i < len(A) && A[i] == A[i-1] {
			continue
		}
		A[j] = A[i-1]
		j++
	}

	val := int64(INF)
	res := val
	if A[0]%2 == 1 {
		res = -val
	}
	change := false
	for i := 1; i < j; i++ {
		change = change || A[i-1]%2 != A[i]%2
		if change {
			val /= 2
		}

		if A[i]%2 == 0 {
			res += val
		} else {
			res -= val
		}
	}

	return res
}
