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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)
		A := readNNums(scanner, k)
		res := solve(n, k, A)
		if res == nil {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			for i := 0; i < n; i++ {
				if i < n-1 {
					fmt.Printf("%d ", res[i])
				} else {
					fmt.Printf("%d\n", res[i])
				}
			}
		}
	}
}

func solve(n, k int, A []int) []int {
	B := make([]int, n+1)
	if k == 1 {
		if n == 1 {
			B[1] = 1
			return B[1:]
		}
		return nil
	}

	if k == 2 {
		if A[1]-A[0] > 1 {
			return nil
		}
		num := n
		for i := 1; i < A[0]; i++ {
			B[i] = num
			num--
		}
		B[A[1]] = num
		num--
		B[A[0]] = num
		num--

		for i := A[1] + 1; i <= n; i++ {
			B[i] = num
			num--
		}
		return B[1:]
	}

	num := n

	for i := 1; i < A[0]; i++ {
		B[i] = num
		num--
	}
	for i := A[0] + 1; i < A[1]; i++ {
		B[i] = num
		num--
	}
	for j := k - 1; j >= 0; j-- {
		B[A[j]] = num
		num--
	}
	for i := A[1] + 1; i <= n; i++ {
		if B[i] == 0 {
			B[i] = num
			num--
		}
	}
	return B[1:]
}
