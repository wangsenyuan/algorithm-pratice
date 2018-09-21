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
	n, q := readTwoNums(scanner)

	A := readNNums(scanner, n)
	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		queries[i] = readNNums(scanner, 2)
	}

	res := solve(n, A, q, queries)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(n int, A []int, q int, queries [][]int) []int {
	//f[i][j] is the count of index k, between (i, j) that A[i] > A[k] < A[j]

	B := make([]int, n)
	copy(B, A)
	sort.Ints(B)

	var ii int
	for i := 1; i <= n; i++ {
		if i < n && B[i] == B[i-1] {
			continue
		}
		B[ii] = B[i-1]
		ii++
	}
	B = B[:ii]
	for i := 0; i < n; i++ {
		A[i] = sort.SearchInts(B, A[i])
	}

	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	bit_r := NewBit(ii)
	bit_k := NewBit(ii)
	bit_j := NewBit(ii)
	bit_i := NewBit(ii)

	// ii++

	for i := 0; i < n; i++ {
		bit_i.Reset()
		bit_j.Reset()
		bit_k.Reset()
		bit_r.Reset()
		bit_i.Insert(ii-A[i], 1)
		for j := i + 1; j < n; j++ {
			ans[i][j] = bit_r.Query(A[j] - 1)
			bit_r.Insert(A[j], bit_k.Query(ii-A[j]-1))
			bit_k.Insert(ii-A[j], bit_j.Query(A[j]-1))
			bit_j.Insert(A[j], bit_i.Query(ii-A[j]-1))
		}
	}

	res := make([]int, q)

	for i, query := range queries {
		l, r := query[0]-1, query[1]-1
		res[i] = ans[l][r]
	}

	return res
}

type BIT []int

func NewBit(n int) BIT {
	arr := make([]int, n+1)
	return BIT(arr)
}

func (bit BIT) Insert(pos int, val int) {
	n := len(bit) - 1
	pos++
	for pos <= n {
		bit[pos] += val
		pos += pos & (-pos)
	}
}

func (bit BIT) Query(pos int) int {
	var res int
	pos++
	for pos > 0 {
		res += bit[pos]
		pos -= pos & (-pos)
	}
	return res
}

func (bit BIT) Reset() {
	n := len(bit) - 1
	for i := 0; i <= n; i++ {
		bit[i] = 0
	}
}
