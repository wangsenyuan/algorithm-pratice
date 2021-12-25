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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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
		var n int64
		var pos int
		scanner.Scan()
		pos = readInt64(scanner.Bytes(), 0, &n) + 1
		var q int
		readInt(scanner.Bytes(), pos, &q)
		queries := make([][]int64, q)
		for i := 0; i < q; i++ {
			scanner.Scan()
			var t int64
			pos = readInt64(scanner.Bytes(), 0, &t)
			if t == 1 {
				var x int64
				readInt64(scanner.Bytes(), pos+1, &x)
				queries[i] = []int64{t, x}
			} else {
				var l, r int64
				pos = readInt64(scanner.Bytes(), pos+1, &l)
				readInt64(scanner.Bytes(), pos+1, &r)
				queries[i] = []int64{t, l, r}
			}
		}

		res := solve(n, q, queries)
		for _, num := range res {
			fmt.Println(num)
		}
	}
}

func solve(n int64, q int, queries [][]int64) []int64 {
	// fmt.Fprintf(os.Stderr, "%v\n", queries)
	set := make(map[int64]int64)

	var find func(x int64) int64
	find = func(x int64) int64 {
		if set[x] == 0 {
			set[x] = x
			return x
		}
		if set[x] == x {
			return x
		}
		set[x] = find(x - 1)
		return set[x]
	}
	var s int64

	res := make([]int64, q)
	var j int
	for i := 0; i < q; i++ {
		query := queries[i]
		if query[0] == 1 {
			y := query[1]
			x := y + s
			set[x] = find(x - 1)
		} else {
			p := query[1]
			q := query[2]
			l := p + s
			r := q + s
			x := find(r)
			if x < l {
				res[j] = 0
			} else {
				res[j] = x
			}
			s += res[j]
			if s >= n {
				s -= n
			}
			j++
		}
	}
	return res[:j]
}
