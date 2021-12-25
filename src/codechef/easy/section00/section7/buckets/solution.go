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
	n, k := readTwoNums(scanner)
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = readNNums(scanner, k)
	}
	res := solve(n, k, A)
	for i := 0; i < k; i++ {
		if i < k-1 {
			fmt.Printf("%.6f ", res[i])
		} else {
			fmt.Println(res[i])
		}
	}
}

func solve(n int, k int, A [][]int) []float64 {
	f := make([]float64, k)

	totalSum := func(i int) int {
		var res int
		for j := 0; j < k; j++ {
			res += A[i][j]
		}
		return res
	}

	sum := totalSum(0)
	for i := 0; i < k; i++ {
		f[i] = float64(A[0][i]) / float64(sum)
	}

	for i := 1; i < n; i++ {
		sum = totalSum(i)
		for j := 0; j < k; j++ {
			f[j] = (f[j] + float64(A[i][j])) / float64(sum+1)
		}
	}
	return f
}
