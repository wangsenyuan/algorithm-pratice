package main

import (
	"bufio"
	"fmt"
	"os"
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
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
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
		n, k := readTwoNums(scanner)
		A := readNNums(scanner, n)
		fmt.Println(solve(n, k, A))
	}
}

func solve(n int, k int, A []int) int64 {
	odd := make([]int64, n)
	even := make([]int64, n)

	for i := 0; i < n; i++ {
		num := int64(A[i])
		if i > 0 {
			even[i] = even[i-1]
			odd[i] = odd[i-1]
		}
		if A[i]&1 == 1 {
			//odd
			if i-k > 0 {
				odd[i] = max(odd[i-k-1]+num, odd[i-1])
			} else {
				odd[i] = max(odd[i], num)
			}

		} else {
			//even
			if i-k > 0 {
				even[i] = max(even[i-k-1]+num, even[i-1])
			} else {
				even[i] = max(even[i], num)
			}
		}
	}
	return even[n-1] + odd[n-1]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
