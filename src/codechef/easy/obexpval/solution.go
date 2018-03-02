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

	for i := 0; i < t; i++ {
		n, e := readTwoNums(scanner)
		x := readNNums(scanner, n)
		res, can := solve(n, x, e)
		if !can {
			fmt.Println("-1")
		} else {
			fmt.Printf("%.9f", res[0])
			for i := 1; i < n; i++ {
				fmt.Printf(" %.9f", res[i])
			}
			fmt.Println()
		}
	}
}

func solve(n int, x []int, e int) ([]float64, bool) {
	a, b := 0, 0
	for i := 1; i < n; i++ {
		if x[i] < x[a] {
			a = i
		}
		if x[i] > x[b] {
			b = i
		}
	}

	if x[a] > e || x[b] < e {
		return nil, false
	}
	res := make([]float64, n)
	if x[a] == x[b] {
		res[0] = 1.0
		return res, true
	}

	p := float64(x[b]-e) / float64(x[b]-x[a])
	res[a] = p
	res[b] = 1.0 - p
	return res, true
}
