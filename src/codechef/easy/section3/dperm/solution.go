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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, d := readTwoNums(scanner)
		A := readNNums(scanner, n)

		res := solve(n, d, A)
		fmt.Println(res)
	}
}

func solve(n int, d int, A []int) int {
	for i := 0; i < n; i++ {
		A[i]--
	}
	for i := 0; i < n; i++ {
		if i%d != A[i]%d {
			return -1
		}
	}

	arrs := make([][]int, d)

	for i := 0; i < d; i++ {
		arrs[i] = make([]int, 0, n/d)
	}

	for i := 0; i < n; i++ {
		j := A[i] % d
		arrs[j] = append(arrs[j], A[i]/d)
	}

	var res int

	for i := 0; i < d; i++ {
		res += findInversions(arrs[i])
	}

	return res
}

func findInversions(arr []int) int {
	n := len(arr)
	bit := make([]int, n+1)

	set := func(i int) {
		i++
		for i <= n {
			bit[i]++
			i += i & (-i)
		}
	}

	get := func(i int) int {
		var r int
		i++
		for i > 0 {
			r += bit[i]
			i -= i & (-i)
		}
		return r
	}

	var res int

	for i := 0; i < n; i++ {
		res += i - get(arr[i])
		set(arr[i])
	}

	return res
}
