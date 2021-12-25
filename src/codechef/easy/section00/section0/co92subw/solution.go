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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	x = readInt(scanner.Bytes(), x+1, &b)
	readInt(scanner.Bytes(), x+1, &c)
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

func processNNums(bs []byte, n int) []int {
	res := make([]int, n)
	x := -1
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	var res bytes.Buffer

	for t > 0 {
		X := readNum(scanner)
		Y := solve(X)
		res.WriteByte(byte(Y + '0'))
		res.WriteByte('\n')
		t--
	}
	fmt.Print(res.String())
}

func solve(X int) int {
	Y := int64(X)
	n := sort.Search(44721, func(i int) bool {
		j := int64(i)
		return j*(j+1) >= 2*Y
	})

	if n*(n+1) == 2*X {
		return n
	}

	m := n - 1

	a := X - m*(m+1)/2
	b := n*(n+1)/2 - X

	if a+m < b+n {
		return a + m
	}
	return b + n
}
