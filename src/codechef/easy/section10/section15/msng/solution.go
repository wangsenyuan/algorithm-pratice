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
		n := readNum(scanner)
		ss := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			ss[i] = scanner.Text()
		}
		fmt.Println(solve(n, ss))
	}
}

const INF = int64(1000000000000)

func solve(n int, strs []string) int64 {
	cand := int64(-1)

	var unknowns []string

	for i := 0; i < n; i++ {
		base, left := parse(strs[i])
		if base < 0 {
			unknowns = append(unknowns, left)
		} else {
			tmp := parseAsNum(left, base)
			if tmp > INF {
				return -1
			}
			if cand >= 0 && tmp != cand {
				return -1
			}
			cand = tmp
		}
	}

	if len(unknowns) == 0 {
		return cand
	}

	if cand >= 0 {
		return check(cand, unknowns)
	}
	return find(unknowns)
}

func check(cand int64, strs []string) int64 {
	for _, str := range strs {
		xx := tryBases(str)
		if !xx[cand] {
			return -1
		}
	}
	return cand
}

func find(strs []string) int64 {
	xx := tryBases(strs[0])

	for i := 1; i < len(strs); i++ {
		yy := tryBases(strs[i])

		zz := make(map[int64]bool)

		for x := range xx {
			if yy[x] {
				zz[x] = true
			}
		}
		if len(zz) == 0 {
			return -1
		}
		xx = zz
	}
	if len(xx) == 0 {
		return -1
	}
	res := INF

	for x := range xx {
		if x < res {
			res = x
		}
	}
	return res
}

func tryBases(str string) map[int64]bool {
	from := 1
	for i := 0; i < len(str); i++ {
		x := int(parseDigit(str[i]))
		if x > from {
			from = x
		}
	}
	from++

	res := make(map[int64]bool)

	for from <= 36 {
		tmp := parseAsNum(str, from)
		if tmp > INF {
			break
		}
		res[tmp] = true
		from++
	}

	return res
}

func parse(str string) (int, string) {
	if str[0] == '-' {
		return -1, str[3:]
	}

	var base int

	var i int

	for i < len(str) && str[i] != ' ' {
		base = base*10 + int(str[i]-'0')
		i++
	}

	return base, str[i+1:]
}

func parseAsNum(str string, base int) int64 {
	var res int64

	B := int64(base)

	for i := 0; i < len(str); i++ {
		x := parseDigit(str[i])

		res = res*B + x
		if res > INF {
			return INF + 1
		}
	}

	return res
}

func parseDigit(c byte) int64 {
	if c >= '0' && c <= '9' {
		return int64(c - '0')
	}
	return 10 + int64(c-'A')
}
