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
		n := readNum(scanner)
		A := readNNums(scanner, n)
		x := readNum(scanner)
		a := solve(n, A, x)
		fmt.Printf("%d %d\n", a, n-a)
	}
}

func solve(n int, A []int, x int) int {
	var i int
	var j = n - 1
	var sum1, sum2 int
	var cnt1, cnt2 int

	for i < j {
		if sum1 > sum2 {
			sum2 += x * A[j]
			cnt2++
			j--
		} else if sum1 < sum2 {
			sum1 += A[i]
			cnt1++
			i++
		} else {
			sum1 += A[i]
			cnt1++
			i++
			sum2 += x * A[j]
			cnt2++
			j--
		}
	}

	if i == j && sum1 < sum2 {
		cnt1++
	} else if i == j && sum1 > sum2 {
		cnt2++
	} else {
		if cnt1 >= cnt2 {
			cnt1++
		} else {
			cnt2++
		}
	}

	return cnt1
}
