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

	t := readNum(scanner)

	for t > 0 {
		scanner.Scan()
		var X int64

		readInt64(scanner.Bytes(), 0, &X)

		scanner.Scan()

		var B int
		pos := readInt(scanner.Bytes(), 0, &B)
		dishes := make([][]int64, B)

		for i := 0; i < B; i++ {
			var x, y int64
			pos = readInt64(scanner.Bytes(), pos+1, &x)
			pos = readInt64(scanner.Bytes(), pos+1, &y)
			dishes[i] = []int64{x, y}
		}

		scanner.Scan()
		var C int

		pos = readInt(scanner.Bytes(), 0, &C)
		clans := make([][]int64, C)
		for i := 0; i < C; i++ {
			var p, q, r int64
			pos = readInt64(scanner.Bytes(), pos+1, &p)
			pos = readInt64(scanner.Bytes(), pos+1, &q)
			pos = readInt64(scanner.Bytes(), pos+1, &r)
			clans[i] = []int64{p, q, r}
		}

		fmt.Println(solve(X, B, dishes, C, clans))

		t--
	}
}

func solve(X int64, B int, dishes [][]int64, C int, clans [][]int64) uint64 {

	check := func(balance uint64) bool {
		var i, j int

		for i < B && j < C {
			x, y := dishes[i][0], dishes[i][1]
			p, q, r := clans[j][0], clans[j][1], clans[j][2]
			if x < p {
				if balance < uint64(y) {
					// can't afford y persons
					return false
				}
				balance -= uint64(y)
				i++
			} else {
				if balance+1 >= uint64(q) {
					balance += uint64(r)
				}
				j++
			}
		}

		for i < B {
			y := dishes[i][1]
			if balance < uint64(y) {
				return false
			}
			balance -= uint64(y)
			i++
		}
		// no need to process left clans
		return true
	}

	var left, right uint64
	right = 1000000000000000000

	for left < right {
		mid := left + (right-left)>>1
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left + 1
}

func solve1(X int64, B int, dishes [][]int64, C int, clans [][]int64) uint64 {
	var ans uint64 = 1
	// final need only Chef to arrive at X
	i, j := B-1, C-1

	pos := X

	for pos > 0 {
		if i >= 0 && j >= 0 {
			x, y := dishes[i][0], dishes[i][1]

			p, q, r := clans[j][0], clans[j][1], clans[j][2]

			if x > p {
				// first arrive at dish i
				ans += uint64(y)
				i--
				pos = x
			} else {
				var Q uint64
				if uint64(q+r) >= ans {
					Q = uint64(q)
				} else {
					Q = ans - uint64(r)
				}
				ans = min(Q, ans)
				j--
				pos = p
			}
		} else if i >= 0 {
			x, y := dishes[i][0], dishes[i][1]
			ans += uint64(y)
			pos = x
			i--
		} else if j >= 0 {
			p, q, r := clans[j][0], clans[j][1], clans[j][2]
			var Q uint64
			if uint64(q+r) >= ans {
				Q = uint64(q)
			} else {
				Q = ans - uint64(r)
			}
			ans = min(Q, ans)
			j--
			pos = p
		} else {
			pos = 0
		}
	}

	return ans
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
