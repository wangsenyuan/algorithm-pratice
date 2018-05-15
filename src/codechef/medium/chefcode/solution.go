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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var n int
	pos := readInt(scanner.Bytes(), 0, &n)

	var K int64
	readInt64(scanner.Bytes(), pos+1, &K)

	scanner.Scan()

	pos = -1

	A := make([]int64, n)
	for i := 0; i < n; i++ {
		pos = readInt64(scanner.Bytes(), pos+1, &A[i])
	}

	fmt.Println(solve(n, A, K))
}

func solve(n int, A []int64, K int64) int {
	B := A[:n/2]
	C := A[n/2:]

	bruteForce := func(arr []int64) []int64 {
		m := len(arr)
		M := 1 << uint(m)
		res := make([]int64, M)
		var u int
		for i := 0; i < M; i++ {
			var prod int64 = 1
			for j := 0; j < m; j++ {
				if i&(1<<uint(j)) > 0 {
					prod *= arr[j]
				}
				if prod > K || prod < 0 {
					break
				}
			}
			if prod > 0 && prod <= K {
				res[u] = prod
				u++
			}
		}
		return res[:u]
	}

	res1 := bruteForce(B)
	res2 := bruteForce(C)

	sort.Sort(Int64Slice(res1))
	sort.Sort(Int64Slice(res2))

	var ans int

	for _, a := range res1 {
		j := sort.Search(len(res2), func(j int) bool {
			return a*res2[j] > K
		})
		ans += j
	}

	return ans - 1
}

type Int64Slice []int64

func (this Int64Slice) Len() int {
	return len(this)
}

func (this Int64Slice) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64Slice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
