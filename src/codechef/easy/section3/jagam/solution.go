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
		line := readNNums(scanner, 3)
		N, Z1, Z2 := line[0], line[1], line[2]
		A := readNNums(scanner, N)
		fmt.Println(solve(N, A, Z1, Z2))
	}
}

func solve(N int, A []int, Z1, Z2 int) int {
	for i := 0; i < N; i++ {
		x := abs(A[i])
		if x == abs(Z1) || x == abs(Z2) {
			return 1
		}
	}
	X := int64(Z1)
	Y := int64(Z2)
	// one move can't win
	var count int
	for i := 0; i < N; i++ {
		var cnt1, cnt2 int
		for j := 0; j < N; j++ {
			u := int64(A[i])
			v := int64(A[j])
			a := u - v
			b := u + v
			c := -u + v
			d := -u - v
			if a == X || a == Y || b == X || b == Y {
				// if vanju choose u, Miksi can find a number to get Z1 or Z2
				cnt1 = 1
			}
			if c == X || c == Y || d == X || d == Y {
				// if vanju choose -u, Miksi can find a number to get Z1 or Z2
				cnt2 = 1
			}
		}
		count += cnt1 + cnt2
	}
	if count == 2*N {
		return 2
	}
	// a tie
	return 0
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
