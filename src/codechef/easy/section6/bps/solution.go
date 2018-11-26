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
		n, m := readTwoNums(scanner)
		A := readNNums(scanner, n)
		ops := make([][]int, m)
		for i := 0; i < m; i++ {
			ops[i] = readNNums(scanner, 2)
		}
		p, q := solve(n, A, m, ops)

		fmt.Printf("%d/%d\n", p, q)
	}
}

func solve(n int, A []int, m int, ops [][]int) (int, int) {
	B := make([]int, n)
	copy(B, A)

	for i := 0; i < m; i++ {
		op := ops[i]
		x, y := op[0]-1, op[1]-1
		swap(B, x, y)
	}

	C := make([]int, n)
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = i
	}
	var q int
	all := 1 << uint(m)
	for tc := 0; tc < all; tc++ {
		copy(C, A)
		for i := 0; i < m; i++ {
			j := p[i]
			x, y := ops[j][0]-1, ops[j][1]-1
			swap(C, x, y)
		}
		same := true
		for i := 0; i < n && same; i++ {
			same = C[i] == B[i]
		}
		if same {
			q++
		}
		nextPermuatation(p)
	}
	g := gcd(q, all)
	return q / g, all / g
}

func nextPermuatation(arr []int) {
	n := len(arr)
	i := n - 1
	for i > 0 && arr[i-1] > arr[i] {
		i--
	}
	//arr[i-1] < arr[i]
	if i == 0 {
		//unable to permutate
		return
	}
	i--
	j := i + 1
	for k := j; k < n; k++ {
		if arr[k] > arr[i] {
			j = k
		}
	}
	arr[i], arr[j] = arr[j], arr[i]

	for a, b := i+1, n-1; a < b; a, b = a+1, b-1 {
		arr[a], arr[b] = arr[b], arr[a]
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func swap(arr []int, x, y int) {
	for i, j := x, y; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
