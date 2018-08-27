package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		var N, K int
		pos := readInt(scanner.Bytes(), 0, &N)
		pos = readInt(scanner.Bytes(), pos+1, &K)
		var x, D int64
		pos = readInt64(scanner.Bytes(), pos+1, &x)
		readInt64(scanner.Bytes(), pos+1, &D)
		P := readNInt64Nums(scanner, K)
		res := solve(N, K, x, D, P)
		fmt.Println(res)
	}
}

func solve(N, K int, x int64, D int64, P []int64) int64 {
	if x-int64(N) < 0 {
		// can't have N distinct postive numbers
		return -1
	}

	sort.Sort(Int64Slice(P))

	for i := 1; i < K; i++ {
		if P[i] == P[i-1] {
			// not distinct
			return -1
		}
	}

	if P[K-1] > x {
		return -1
	}
	var prev int64 = -D
	A := make([]int64, 0, K)
	for j := 0; j < K; j++ {
		A = append(A, P[j])
		if P[j]-prev > D {
			//can't reach previous one
			if j+1 == K || P[j]+D < P[j+1] {
				//can not reach next
				prev = min(P[j]+D, x)
				A = append(A, prev)
			} else {
				prev = P[j]
			}
		} else {
			prev = P[j]
		}
	}
	if len(A) > N {
		return -1
	}
	// then add left (N - K) numbers bewteen [1, x]
	//very least number
	num := x - int64(N) + 1
	sum := (num + x) * int64(N) / 2
	for i := 0; i < len(A) && num > A[i]; i, num = i+1, num+1 {
		sum -= num - A[i]
	}
	return sum
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Int64Slice []int64

func (slice Int64Slice) Len() int {
	return len(slice)
}

func (slice Int64Slice) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice Int64Slice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
