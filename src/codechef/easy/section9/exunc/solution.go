package main

import (
	"bufio"
	"fmt"
	"os"
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

	n, q := readTwoNums(scanner)
	A := readNNums(scanner, n)

	Q := make([][]int, q)

	for i := 0; i < q; i++ {
		scanner.Scan()
		var t int
		pos := readInt(scanner.Bytes(), 0, &t) + 1
		if t == 1 {
			var j, v int
			pos = readInt(scanner.Bytes(), pos, &j) + 1
			readInt(scanner.Bytes(), pos, &v)
			Q[i] = []int{t, j, v}
		} else {
			var j int
			readInt(scanner.Bytes(), pos, &j)
			Q[i] = []int{t, j}
		}
	}

	res := solve(n, A, Q)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, A []int, Q [][]int) []int {
	arr := make([]int, n+1)

	update := func(pos int, v int) {
		pos++
		for pos <= n {
			arr[pos] += v
			pos += pos & -pos
		}
	}

	get := func(pos int) int {
		pos++
		var res int
		for pos > 0 {
			res += arr[pos]
			pos -= pos & -pos
		}
		return res
	}

	for i := 1; i < n; i++ {
		if A[i]%A[i-1] != 0 {
			update(i, 1)
		}
	}

	updateAt := func(i int, x int) {
		if i > 0 {
			if A[i]%A[i-1] == 0 {
				// same as prev
				if x%A[i-1] != 0 {
					// need to update
					update(i, 1)
				}
				// no change
			} else if x%A[i-1] == 0 {
				update(i, -1)
			}
		}

		if i < n-1 {
			if A[i+1]%A[i] == 0 {
				if A[i+1]%x != 0 {
					update(i+1, 1)
				}
			} else if A[i+1]%x == 0 {
				update(i+1, -1)
			}
		}

		A[i] = x
	}

	find := func(i int) int {
		x := get(i)

		var left int
		right := i

		for left < right {
			mid := (left + right) / 2
			y := get(mid)
			if x == y {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return right
	}

	var m int
	res := make([]int, len(Q))

	for i := 0; i < len(Q); i++ {
		q := Q[i]
		if q[0] == 1 {
			updateAt(q[1]-1, q[2])
		} else {
			res[m] = find(q[1]-1) + 1
			m++
		}
	}

	return res[:m]
}
