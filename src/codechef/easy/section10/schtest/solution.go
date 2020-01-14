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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		bs := scanner.Bytes()
		var N, R uint64
		pos := readUint64(bs, 0, &N)
		pos = readUint64(bs, pos+1, &R)
		var x, y int
		pos = readInt(bs, pos+1, &x)
		readInt(bs, pos+1, &y)

		X := make([]int64, x)

		if x > 0 {
			var num uint64
			scanner.Scan()
			pos = 0
			for j := 0; j < x; j++ {
				pos = readUint64(scanner.Bytes(), pos, &num) + 1
				X[j] = int64(num)
			}
		}
		Y := make([]int64, y)
		if y > 0 {
			var num uint64
			scanner.Scan()
			pos = 0
			for j := 0; j < y; j++ {
				pos = readUint64(scanner.Bytes(), pos, &num) + 1
				Y[j] = int64(num)
			}
		}

		fmt.Println(solve(int64(N), int64(R), X, Y))
	}
}

func solve(N int64, R int64, X []int64, Y []int64) int64 {
	if R == 0 {
		return 0
	}

	sort.Sort(Int64Nums(X))
	sort.Sort(Int64Nums(Y))

	var same int

	for i, j := 0, 0; i < len(X) && j < len(Y); {
		if X[i] == Y[j] {
			i++
			j++
			same++
		} else if X[i] < Y[j] {
			i++
		} else {
			j++
		}
	}

	ans := N - int64(len(X)) - int64(len(Y)) + int64(same)

	return min(R, ans)
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
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
