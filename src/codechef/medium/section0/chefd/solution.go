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
	scanner.Scan()
	bs := scanner.Bytes()
	return getNNums(bs, n)
}

func getNNums(bs []byte, n int) []int {
	res := make([]int, n)
	x := -1
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	A := readNNums(scanner, n)

	m := readNum(scanner)
	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		scanner.Scan()
		if scanner.Bytes()[0] == '1' {
			queries[i] = getNNums(scanner.Bytes(), 4)
		} else {
			queries[i] = getNNums(scanner.Bytes(), 3)
		}
	}
	solve(n, A, queries)

	fmt.Printf("%d", A[0])

	for i := 1; i < n; i++ {
		fmt.Printf(" %d", A[i])
	}
	fmt.Println()
}

func solve(n int, A []int, queries [][]int) {
	bits := make([]*BIT, 6)
	bits[2] = CreateBIT(n)
	bits[3] = CreateBIT(n)
	bits[5] = CreateBIT(n)

	getAns := func(num int, pos int) int {
		for _, x := range []int{2, 3, 5} {
			var cnt int

			for num%x == 0 {
				cnt++
				num /= x
			}

			cnt -= bits[x].Read(pos)
			for j := 0; j < cnt; j++ {
				num *= x
			}
		}
		return num
	}

	B := make([]int, n)
	for i := len(queries) - 1; i >= 0; i-- {
		query := queries[i]
		if query[0] == 1 {
			l, r, p := query[1]-1, query[2]-1, query[3]
			bits[p].Update(l, 1)
			bits[p].Update(r+1, -1)
		} else {
			l, d := query[1]-1, query[2]
			if B[l] == 0 {
				B[l] = getAns(d, l)
			}
		}
	}

	for i := 0; i < n; i++ {
		if B[i] == 0 {
			B[i] = getAns(A[i], i)
		}
	}
	copy(A, B)
}

type BIT struct {
	arr []int
	n   int
}

func CreateBIT(n int) *BIT {
	arr := make([]int, n+1)
	return &BIT{arr, n}
}

func (bit *BIT) Update(i int, dx int) {
	i++
	for i <= bit.n {
		bit.arr[i] += dx
		i += i & (-i)
	}
}

func (bit BIT) Read(i int) int {
	i++
	var sum int

	for i > 0 {
		sum += bit.arr[i]
		i -= i & (-i)
	}

	return sum
}
