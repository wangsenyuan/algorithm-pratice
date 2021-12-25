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

	tc := readNum(scanner)

	for tc > 0 {
		scanner.Scan()
		bs := scanner.Bytes()

		pos := len(bs)

		for bs[pos-1] != ' ' {
			pos--
		}
		s := bs[:pos-1]
		var k int
		readInt(bs, pos, &k)
		fmt.Println(solve(string(s), k))
	}
}

func solve(s string, k int) string {
	buf := make(map[byte]bool)

	n := len(s)

	for i := 0; i < n; i++ {
		buf[s[i]] = true
	}

	res := make([]byte, n)

	var i int
	var x byte = 'a'

	for i < n && k > 0 {
		res[i] = x
		if buf[x] {
			k--
		}
		i++
		x++
	}

	// k == 0
	for i < n {
		for x <= 'z' && buf[x] {
			x++
		}
		if x > 'z' {
			break
		}
		res[i] = x
		i++
		x++
	}

	if i == n {
		return string(res)
	}
	return "NOPE"
}
