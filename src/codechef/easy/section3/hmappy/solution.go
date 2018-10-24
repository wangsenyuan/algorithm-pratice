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
	scanner.Scan()
	var n int
	pos := readInt(scanner.Bytes(), 0, &n)

	var M uint64

	readUint64(scanner.Bytes(), pos+1, &M)
	A := readNNums(scanner, n)
	B := readNNums(scanner, n)

	fmt.Println(solve(n, int64(M), A, B))
}

func solve(n int, M int64, A []int, B []int) int64 {
	var S int64

	for i := 0; i < n; i++ {
		S += int64(A[i])
	}

	if S <= M {
		return 0
	}

	// for i := 0; i < n; i++ {
	// 	P[i] = Pair{A[i], B[i]}
	// }

	// sort.Sort(Pairs(P[:n]))

	check := func(C int64) bool {
		// C[i] = max(0, A[i] - X[i]) * B[i] <= C
		if C == 0 {
			return false
		}
		var S int64
		for i := 0; i < n; i++ {
			x := max(0, int64(A[i])-C/int64(B[i]))
			//x is the count to staisfy the condition
			S += x
		}
		return S <= M
	}

	var left int64
	var right int64

	for i := 0; i < n; i++ {
		tmp := int64(A[i]) * int64(B[i])
		right = max(right, tmp)
	}

	for left < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

// type Pair struct {
// 	first, second int
// }

// type Pairs []Pair

// func (ps Pairs) Len() int {
// 	return len(ps)
// }

// func (ps Pairs) Less(i, j int) bool {
// 	return ps[i].second > ps[j].second
// }

// func (ps Pairs) Swap(i, j int) {
// 	ps[i], ps[j] = ps[j], ps[i]
// }
