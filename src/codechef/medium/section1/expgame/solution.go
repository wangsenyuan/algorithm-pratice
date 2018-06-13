package main

import (
	"bufio"
	"bytes"
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

	var buf bytes.Buffer

	for t > 0 {
		n := readNum(scanner)
		A := readNNums(scanner, n)
		res := solve(n, A)
		if res {
			buf.WriteString("Little Chef\n")
		} else {
			buf.WriteString("Head Chef\n")
		}
		t--
	}
	fmt.Print(buf.String())
}

const MaxN = 100334

var sg = make([]int, MaxN)

func init() {
	deg := make([]int, 123)
	for i := 1; i <= 7; i++ {
		deg[i] = 1
		for j := 1; j <= i; j++ {
			deg[i] = deg[i] * i
		}
	}

	var o int
	used := make([]int, 1000)
	for i := 1; i < MaxN; i++ {
		o++
		for j := 1; deg[j] <= i; j++ {
			used[sg[i-deg[j]]] = o
		}
		for j := 0; j < 1000; j++ {
			if used[j] != o {
				sg[i] = j
				break
			}
		}
	}
}

func solve(n int, A []int) bool {
	var res int

	for i := 0; i < n; i++ {
		res ^= A[i]
	}

	return res > 0
}
