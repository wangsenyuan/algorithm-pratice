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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, k, A))
	}
}

func solve(N, K int, A []int) int {
	ans := make([]int64, N+1)
	sum := make([]int64, N+2)

	for i := 1; i <= N; i++ {
		x := int64(A[i-1])
		pos := i

		for j := int64(1); pos > 0 && j*j <= x; j, pos = j+1, pos-1 {
			ans[pos] += x / j
		}

		if pos == 0 {
			continue
		}

		prev := x / int64(i-pos)

		for x/int64(i-pos+1) == prev {
			ans[pos] += x / int64(i-pos+1)
			pos--
		}

		if pos == 0 {
			continue
		}

		for v := x / int64(i-pos+1); v > 0 && pos > 0; v-- {
			cnt := int(x/v) - int(x/(v+1))
			sum[max(1, pos-cnt+1)] += v
			sum[pos+1] -= v
			pos -= cnt
		}
	}

	var suma int64

	for i := 1; i <= N; i++ {
		suma += sum[i]
		if ans[i]+suma <= int64(K) {
			return i
		}
	}
	return N + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
