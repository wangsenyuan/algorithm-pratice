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
	var N, Q, K uint64
	pos := readUint64(scanner.Bytes(), 0, &N)
	pos = readUint64(scanner.Bytes(), pos+1, &Q)
	readUint64(scanner.Bytes(), pos+1, &K)
	n := int(N)
	D := make([]int, n)
	X := make([]int64, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		pos = readInt(scanner.Bytes(), 0, &D[i])
		var x uint64
		readUint64(scanner.Bytes(), pos+1, &x)
		X[i] = int64(x)
	}

	T := make([]int64, int(Q))

	for i := 0; i < int(Q); i++ {
		var t uint64
		scanner.Scan()
		readUint64(scanner.Bytes(), 0, &t)
		T[i] = int64(t)
	}

	res := solve(n, int(Q), int64(K), D, X, T)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func solve(N, Q int, K int64, D []int, X []int64, T []int64) []int64 {
	// sort.Sort(Int64Nums(X))

	cars1 := make([]int64, 0, N)
	cars2 := make([]int64, 0, N)
	for i := 0; i < N; i++ {
		if D[i] == 1 {
			cars1 = append(cars1, X[i])
		} else {
			cars2 = append(cars2, X[i])
		}
	}

	sort.Sort(Int64Nums(cars1))
	sort.Sort(Int64Nums(cars2))
	cnt2 := len(cars2)
	for i := 0; i < cnt2; i++ {
		cars2 = append(cars2, cars2[i]+K)
	}

	cnt1 := len(cars1)

	ans := make([]int64, Q)
	for i := 0; i < Q; i++ {
		t := T[i]
		t *= 2

		ans[i] = (t / K) * int64(cnt1) * int64(cnt2)

		t %= K

		var l, r int

		for _, a := range cars1 {
			// a := cars1[i]

			for l < len(cars2) && cars2[l] < a {
				l++
			}
			for r < len(cars2) && cars2[r] <= a+t {
				r++
			}

			ans[i] += int64(r - l)
		}
	}
	return ans
}

type Int64Nums []int64

func (this Int64Nums) Len() int {
	return len(this)
}

func (this Int64Nums) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64Nums) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
