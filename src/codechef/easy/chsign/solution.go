package main

import (
	"bufio"
	"os"
	"fmt"
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
		n := readNum(scanner)
		A := readNNums(scanner, n)
		res := solve(n, A)

		fmt.Printf("%d", res[0])
		for i := 1; i < n; i++ {
			fmt.Printf(" %d", res[i])
		}
		fmt.Println()
		t--
	}
}

func solve(n int, A []int) []int {
	B := make([]int, n+2)
	copy(B[1:], A)
	B[0] = 1 << 30
	B[n+1] = 1 << 30
	C := make([]int, n+2)
	i := 1
	for i <= n {
		if B[i] < B[i-1] && B[i] < B[i+1] {
			j := i
			i += 2
			for i <= n && B[i] < B[i-1] && B[i] < B[i+1] && B[i]+B[i-2] >= B[i-1] {
				i += 2
			}
			a := 0
			b := B[j]
			k := j + 2
			for k < i {
				if a+B[k] >= b {
					C[k] = a + B[k]
				} else {
					C[k] = b
				}
				a, b = b, C[k]
				k += 2
			}

			k -= 2

			for k >= j {
				if k == j || C[k] > C[k-2] {
					B[k] = -B[k]
					k -= 4
				} else {
					k -= 2
				}
			}

		} else {
			i++
		}

	}

	return B[1 : n+1]
}
