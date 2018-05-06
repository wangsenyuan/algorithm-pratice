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
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(scanner, 3)
	}
	res := solve(n, m, A, queries)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func solve(n int, m int, A []int, queries [][]int) []int {
	sln := CreateSolution(n, A)

	res := make([]int, m)
	m = 0
	for _, query := range queries {
		a, b, c := query[0], query[1], query[2]
		if a == 1 {
			sln.Update(b-1, c)
		} else {
			res[m] = sln.Query(b-1, c)
			m++
		}
	}
	return res[:m]
}

type Solution struct {
	arr       []int
	t         []int
	p         []int
	f         [][]int
	blockSize int
	blockCnt  int
	size      int
}

func CreateSolution(n int, A []int) *Solution {
	b := 1
	for b*b < n {
		b++
	}

	size := (n + b - 1) / b
	t := make([]int, size)
	p := make([]int, n)
	f := make([][]int, size)
	arr := make([]int, n)

	for i := 0; i < size; i++ {
		f[i] = make([]int, 1000000)
	}

	for i := 0; i < n; i++ {
		arr[i] = A[i]
		p[i] = arr[i]
		if i > 0 {
			p[i] = p[i-1] ^ arr[i]
		}
		f[i/b][p[i]]++
	}
	return &Solution{arr, t, p, f, b, size, n}
}

func (sln *Solution) Update(i int, x int) {
	c := sln.arr[i] ^ x
	sln.arr[i] = x
	block := i / sln.blockSize
	end := (block + 1) * sln.blockSize
	if end > sln.size {
		end = sln.size
	}
	for j := i; j < end; j++ {
		sln.f[block][sln.p[j]]--
		sln.p[j] ^= c
		sln.f[block][sln.p[j]]++
	}
	start := block + 1
	for j := start; j < sln.blockCnt; j++ {
		sln.t[j] ^= c
	}
}

func (sln Solution) Query(i int, x int) int {
	block := i / sln.blockSize
	var ans int
	start := block * sln.blockSize

	for j := start; j <= i; j++ {
		if sln.p[j]^sln.t[block] == x {
			ans++
		}
	}
	for j := 0; j < block; j++ {
		ans += sln.f[j][x^sln.t[j]]
	}
	return ans
}
