package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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
		n := readNum(scanner)
		A := readNNums(scanner, n)
		res := solve(n, A)
		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			buf.WriteByte(byte('0' + res[i]))
		}
		fmt.Println(buf.String())
		t--
	}
}

func solve(n int, A []int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	// k = 1
	res[0] = 1

	sums := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + int64(A[i])
	}

	check := func(k int) bool {
		if sums[n]%int64(k) != 0 {
			return false
		}
		part := sums[n] / int64(k)

		for i := 1; i <= k; i++ {
			left := part * int64(i)
			j := sort.Search(n+1, func(j int) bool {
				return sums[j] >= left
			})

			if j == n+1 || sums[j] > left {
				// not found
				return false
			}
		}
		return true
	}

	for i := 1; i < n; i++ {
		if res[i] == -1 {
			// no calculate yet
			can := check(i + 1)
			if can {
				res[i] = 1
			} else {
				// it i can't, the muliple of i can't either
				k := i + 1
				for j := k; j <= k; j += k {
					res[j-1] = 0
				}
			}
		}
	}

	return res
}
