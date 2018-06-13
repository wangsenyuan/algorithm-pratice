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

		res := solve(n, A, K)

		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}

func solve(n int, A []int, K int) bool {
	g := A[0]
	for i := 1; i < n; i++ {
		g = gcd(g, A[i])
	}

	if g <= K {
		// just divides g
		return true
	}

	h := g
	maxPrime := 1
	for x := 2; x*x <= g; x++ {
		for h%x == 0 {
			h /= x
			maxPrime = x
		}
	}

	if h > maxPrime {
		maxPrime = h
	}

	if maxPrime > K {
		return false
	}

	return true
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
