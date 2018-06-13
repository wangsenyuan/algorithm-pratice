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

		fmt.Println(solve(n, A))

		t--
	}
}

func solve(n int, A []int) int64 {
	C := make([][]int, 30)
	for i := 0; i < 30; i++ {
		C[i] = make([]int, 2)
		C[i][0] = 1
	}
	var S int
	var res int64
	for i := 0; i < n; i++ {
		S ^= A[i]

		for j := 0; j < 30; j++ {
			p := 1 << uint(j)
			if (S>>uint(j))&1 == 1 {
				res += int64(p) * int64(C[j][0])
			} else {
				res += int64(p) * int64(C[j][1])
			}
		}

		for j := 0; j < 30; j++ {
			if (S>>uint(j))&1 == 1 {
				C[j][1]++
			} else {
				C[j][0]++
			}
		}
	}

	return res
}
