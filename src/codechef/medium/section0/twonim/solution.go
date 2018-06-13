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

func readNums(scanner *bufio.Scanner) []int {
	res := make([]int, 0, 10)
	x := -1
	scanner.Scan()
	bs := scanner.Bytes()
	for x < len(bs) {
		var num int
		x = readInt(bs, x+1, &num)
		res = append(res, num)
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

		if res {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}

		t--
	}
}

func solve(n int, A []int) bool {
	if n%2 == 0 {
		return true
	}

	var res int
	for _, a := range A {
		res ^= a
	}
	return res == 0
}
