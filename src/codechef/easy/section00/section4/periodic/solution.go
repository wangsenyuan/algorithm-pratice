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
		n := readNum(scanner)
		A := readNNums(scanner, n)
		res := solve(n, A)
		if res == -1 {
			fmt.Println("impossible")
		} else if res == -2 {
			fmt.Println("inf")
		} else {
			fmt.Println(res)
		}
	}
}

func solve(n int, A []int) int {
	B := make([]int, n)
	var p int
	prev := -1
	var atLeast int
	for i := 0; i < n; i++ {
		if A[i] < 0 {
			continue
		}
		atLeast = max(atLeast, A[i])
		B[p] = i - A[i] + 1
		p++

		if prev >= 0 {
			if i-prev != A[i]-A[prev] && A[i] > i-prev {
				return -1
			}
		}
		prev = i
	}
	sort.Ints(B[:p])
	var pp int
	for i := 1; i <= p; i++ {
		if i < p && B[i] == B[i-1] {
			continue
		}
		B[pp] = B[i-1]
		pp++
	}
	if pp <= 1 {
		//INF
		return -2
	}
	var ans int
	for i := 1; i < pp; i++ {
		ans = gcd(ans, B[i]-B[i-1])
	}

	if ans < atLeast {
		return -1
	}

	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
