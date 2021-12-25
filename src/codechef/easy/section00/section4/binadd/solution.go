package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		A := scanner.Text()
		scanner.Scan()
		B := scanner.Text()
		fmt.Println(solve(A, B))
	}
}

func solve(A, B string) int {
	n := max(len(A), len(B)) + 1

	A = normalize(A, n)
	B = normalize(B, n)
	var cnt int
	var res int
	var carry int

	for i := n - 1; i >= 0; i-- {
		carry += int(A[i]-'0') + int(B[i]-'0')
		if carry > 1 {
			cnt++
		} else {
			cnt = 0
		}

		carry /= 2

		res = max(res, cnt)
	}

	x := strings.Index(B, "1")

	if x >= 0 {
		res++
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func normalize(s string, n int) string {
	var buf bytes.Buffer

	for i := 0; i < n-len(s); i++ {
		buf.WriteByte('0')
	}

	buf.WriteString(s)

	return buf.String()
}
