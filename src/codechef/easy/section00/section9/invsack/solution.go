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

	N := readNum(scanner)

	query := func(i, j int) int {
		fmt.Printf("1 %d %d\n", i, j)
		return readNum(scanner)
	}

	w := make([]int, N)
	v := make([]int, N)

	solve(N, query, w, v)

	fmt.Println(2)

	for i := 0; i < N; i++ {
		if i < N-1 {
			fmt.Printf("%d ", w[i])
		} else {
			fmt.Printf("%d\n", w[i])
		}
	}

	for i := 0; i < N; i++ {
		if i < N-1 {
			fmt.Printf("%d ", v[i])
		} else {
			fmt.Printf("%d\n", v[i])
		}
	}
}

func solve(n int, query func(i, j int) int, w []int, v []int) {

	checkWeight := func(i int, W, V int) int {
		var left = 1
		var right = W

		for left < right {
			mid := (left + right) / 2
			x := query(i, mid)
			if x == V {
				right = mid
			} else {
				left = mid + 1
			}
		}

		return right
	}

	V := query(n, 10000)
	W := checkWeight(n, 10000, V)

	// dp[n][W] = V

	for i := n - 1; i > 0; i-- {
		x := query(i, W)
		v[i] = V - x
		y := checkWeight(i, W, x)
		w[i] = W - y
		V = x
		W = y
	}
	w[0] = W
	v[0] = V
}
