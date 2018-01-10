package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
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
	t := readNum(scanner)
	for i := 0; i < t; i++ {
		n := readNum(scanner)
		rs := readNNums(scanner, n)
		es := make([]int64, n)
		for j := 0; j < n; j++ {
			es[j] = int64(rs[j])
		}
		from, end := readTwoNums(scanner)
		fmt.Println(solve(n, es, from, end))
	}
}

func solve(n int, es []int64, from int, end int) int64 {
	rs := make([]int64, 2*n)
	copy(rs, es)
	for i := n; i < 2*n; i++ {
		rs[i] = rs[i-n]
	}

	sum := make([]int64, 2*n+1)

	for i := 1; i <= 2*n; i++ {
		sum[i] = sum[i-1] + rs[i-1]
	}

	if from > end {
		from, end = end, from
	}

	max, tmp := int64(0), int64(0)
	for i := from - 1; i < end-1; i++ {
		tmp += rs[i]
		if tmp > max {
			max = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}

	res1 := sum[n+from-1] - sum[end-1] + 2*(sum[end-1]-sum[from-1]-max)

	max, tmp = int64(0), int64(0)
	for i := end - 1; i < n+from-1; i++ {
		tmp += rs[i]
		if tmp > max {
			max = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	res2 := sum[end-1] - sum[from-1] + 2*(sum[n+from-1]-sum[end-1]-max)
	return min(res1, res2)
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
