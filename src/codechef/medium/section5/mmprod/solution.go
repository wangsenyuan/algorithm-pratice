package main

import (
	"bufio"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, k := readTwoNums(scanner)
		A := readNNums(scanner, n)
		res := solve(n, k, A)
		fmt.Println(res)
	}
}

const MOD = 1e9 + 7

func solve(n, k int, A []int) int {
	sort.Sort(AbsValSlice(A))

	var negCnt int

	for i, j := n-1, 0; i >= 0 && j < k; i, j = i-1, j+1 {
		if A[i] < 0 {
			negCnt++
		}
	}

	if negCnt&1 == 0 {
		return caseOne(n, k, A)
	}

	//i1 is the first postive index that >= n - k
	i1 := n - k
	for i1 < n && A[i1] < 0 {
		i1++
	}
	//i2 is the first negative index that < n - k, from right to left
	i2 := n - k - 1
	for i2 >= 0 && A[i2] > 0 {
		i2--
	}

	i3 := n - k
	for i3 < n && A[i3] > 0 {
		i3++
	}
	i4 := n - k - 1
	for i4 >= 0 && A[i4] < 0 {
		i4--
	}

	if (i1 == n || i2 == -1) && (i3 == n || i4 == -1) {
		return caseTwo(n, k, A)
	}

	if i1 == n || i2 == -1 {
		// switch A[i3] & A[i4]
		A[i3], A[i4] = A[i4], A[i3]
		return caseOne(n, k, A)
	}

	if i3 == n || i4 == -1 {
		A[i1], A[i2] = A[i2], A[i1]
		return caseOne(n, k, A)
	}

	if A[i1]*A[i4] < A[i2]*A[i3] {
		A[i1], A[i2] = A[i2], A[i1]
		return caseOne(n, k, A)
	}
	A[i3], A[i4] = A[i4], A[i3]
	return caseOne(n, k, A)
}

func caseOne(n, k int, A []int) int {
	res := int64(1)
	for i, j := n-1, 0; i >= 0 && j < k; i, j = i-1, j+1 {
		res = modMul(res, int64(A[i]))
	}

	return int(res)
}

func caseTwo(n, k int, A []int) int {
	res := int64(1)
	for i := 0; i < k; i++ {
		res = modMul(res, int64(A[i]))
	}

	return int(res)
}

type AbsValSlice []int

func (this AbsValSlice) Len() int {
	return len(this)
}

func (this AbsValSlice) Less(i, j int) bool {
	return abs(this[i]) < abs(this[j])
}

func (this AbsValSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func modMul(a, b int64) int64 {
	c := (a * b) % MOD
	if c < 0 {
		return c + MOD
	}
	return c
}
