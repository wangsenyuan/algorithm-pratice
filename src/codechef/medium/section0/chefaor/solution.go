package main

import (
	"bufio"
	"fmt"
	"os"
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
		n, K := readTwoNums(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, K, A))
		t--
	}
}

func solve(n, K int, A []int) int64 {
	f := make([][]int64, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int64, K+1)
	}
	// f[i][j] = max result when A[:(i+1)] splitted into j groups

	f[0][1] = int64(A[0])
	for i := 1; i < n; i++ {
		f[i][1] = f[i-1][1] | int64(A[i])
	}

	for k := 2; k <= K; k++ {
		for i := 1; i < n; i++ {
			// calculate f[i][k]
			sum := int64(A[i])
			best := f[i-1][k-1] + int64(A[i])
			for j := i - 1; j > k-2; j-- {
				sum = sum | int64(A[j])
				tmp := sum + f[j-1][k-1]
				if tmp > best {
					best = tmp
				}
			}
			f[i][k] = best
		}
	}
	return f[n-1][K]
}
