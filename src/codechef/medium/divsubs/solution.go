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
		n := readNum(scanner)
		A := readNNums(scanner, n)
		begin, end := solve(n, A)
		fmt.Println(end - begin + 1)
		fmt.Printf("%d", begin+1)

		for i := begin + 1; i <= end; i++ {
			fmt.Printf(" %d", i+1)
		}
		fmt.Println()

		t--
	}
}

func solve(n int, A []int) (begin int, end int) {
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		prev[i] = n
	}
	prev[0] = -1
	var sum int
	for i := 0; i < n; i++ {
		sum = (sum + A[i]%n) % n
		if prev[sum] < i {
			begin, end = prev[sum]+1, i
			return
		}
		prev[sum] = i
	}
	return -1, -1
}
