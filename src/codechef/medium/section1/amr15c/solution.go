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
		n, k := readTwoNums(scanner)

		res := solve(n, k)

		if res == nil {
			fmt.Println(-1)
		} else {
			for i := 0; i < n; i++ {
				if i < n-1 {
					fmt.Printf("%d ", res[i])
				} else {
					fmt.Printf("%d\n", res[i])
				}
			}
		}
		t--
	}
}

func solve(n, k int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i + 1
	}
	if k == 0 {
		return res
	}

	if k > n/2 {
		return nil
	}

	m := 2 * k

	var pos int
	for pos+m <= n {
		i := pos
		for j := i; j < i+k; j++ {
			res[j] = j + 1 + k
		}
		for j := i + k; j < i+m; j++ {
			res[j] = j + 1 - k
		}
		pos += m
	}

	if n == pos {
		return res
	}

	swap := func(start, end int) {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	pos -= k
	if pos+m < n {
		pos = n - m
	}

	swap(pos, n-1)
	swap(pos, n-k-1)
	swap(n-k, n-1)

	return res
}
