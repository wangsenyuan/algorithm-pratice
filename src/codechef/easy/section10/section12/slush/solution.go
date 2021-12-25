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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, M := readTwoNums(scanner)
		C := readNNums(scanner, M)
		D := make([]int, N)
		F := make([]int, N)
		B := make([]int, N)
		for i := 0; i < N; i++ {
			row := readNNums(scanner, 3)
			D[i] = row[0]
			F[i] = row[1]
			B[i] = row[2]
		}

		res, arr := solve(M, C, N, D, F, B)
		fmt.Println(res)
		for i := 0; i < N; i++ {
			if i < N-1 {
				fmt.Printf("%d ", arr[i])
			} else {
				fmt.Printf("%d\n", arr[i])
			}
		}
	}
}

func solve(M int, C []int, N int, D []int, F []int, B []int) (int, []int) {
	give := make([]int, N)
	var res int
	for i := 0; i < N; i++ {
		d := D[i] - 1
		if C[d] > 0 {
			C[d]--
			res += F[i]
			give[i] = d + 1
		} else {
			res += B[i]

		}
	}

	var j int

	for i := 0; i < N; i++ {
		if give[i] == 0 {
			for C[j] == 0 {
				j++
			}
			give[i] = j + 1
			C[j]--
		}
	}

	return res, give
}
