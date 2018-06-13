package main

import (
	"sort"
	"bufio"
	"os"
	"fmt"
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
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

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	var X int64
	scanner.Scan()
	pos := readInt(scanner.Bytes(), 0, &n)
	readInt64(scanner.Bytes(), pos+1, &X)
	A := readNInt64Nums(scanner, n)
	sad, happy := solve(n, A, X)
	fmt.Printf("%d %d\n", sad, happy)
}

func solve(n int, A []int64, X int64) (int, int) {
	sort.Sort(Int64Slice(A))

	happy, sad := 0, n

	var completed int64
	for i := 0; i < n; i++ {
		x := int64((A[i] + 1) / 2)
		if completed+x > X {
			break
		}
		sad--
		completed += x
	}

	for i := 0; i < n; i++ {
		x := int64((A[i] + 1) / 2)
		y := int64(A[i]) - x
		if completed+y > X {
			break
		}
		happy++
		completed += y
	}

	return sad, happy
}

type Int64Slice []int64

func (s Int64Slice) Len() int {
	return len(s)
}

func (s Int64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Int64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func solve1(n int, A []int, X int64) (int, int) {
	S1 := make([]int64, n+1)
	S2 := make([]int64, n+1)
	sort.Ints(A)

	for i := 0; i < n; i++ {
		x := (A[i] + 1) / 2
		S1[i+1] = S1[i] + int64(x)
		S2[i+1] = S2[i] + int64(A[i]-x)
	}

	notSad := sort.Search(n+1, func(i int) bool {
		return S1[i] > X
	})

	notSad--

	sad := n - notSad

	Y := X - S1[notSad]

	happy := sort.Search(n+1, func(i int) bool {
		return S2[i] > Y
	})
	happy--

	return sad, happy
}
