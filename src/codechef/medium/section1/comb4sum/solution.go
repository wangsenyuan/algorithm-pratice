package main

import (
	"bufio"
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
		fmt.Println(solve(n, A))
		t--
	}
}

func solve(n int, A []int) int64 {
	sort.Ints(A)
	sums := make([]int, n*n)
	var k int

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sums[k] = A[i] + A[j]
			k++
		}
	}

	sort.Ints(sums[:k])

	fn := func(arr []int) int64 {
		var res int64

		for i := 0; i < len(arr); i++ {
			a := i
			b := len(arr) - i - 1
			res += 2 * int64(a-b) * int64(arr[i])
		}

		return res
	}

	return fn(sums[:k]) - int64(n-2)*fn(A)
}
