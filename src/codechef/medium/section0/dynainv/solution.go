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

	n, m := readTwoNums(scanner)

	A := readNNums(scanner, n)

	for i := 0; i < m; i++ {
		scanner.Scan()
	}

	res := solve(n, A, m)

	fmt.Printf("%d", res[0])

	for i := 1; i < m; i++ {
		fmt.Printf(" %d", res[i])
	}
	fmt.Println()
}

func solve(n int, A []int, m int) []int {
	bit := BIT(n + 1)

	fn := func(read func(int) int, update func(int, int)) int64 {
		var cnt int64
		for i := n - 1; i >= 0; i-- {
			cnt += int64(read(A[i]))
			update(A[i], 1)
		}
		return cnt
	}

	invs := int(bit(fn) % 2)

	res := make([]int, m)

	for i := 0; i < m; i++ {
		invs = (invs + 1) % 2
		res[i] = invs
	}

	return res
}

func BIT(n int) func(fn func(func(int) int, func(int, int)) int64) int64 {
	tree := make([]int, n+1)

	read := func(i int) int {
		var sum int
		i++
		for i > 0 {
			sum += tree[i]
			i -= i & (-i)
		}

		return sum
	}

	update := func(i int, val int) {
		i++
		for i <= n {
			tree[i] += val
			i += i & (-i)
		}
	}

	gn := func(fn func(func(int) int, func(int, int)) int64) int64 {
		return fn(read, update)
	}

	return gn
}
