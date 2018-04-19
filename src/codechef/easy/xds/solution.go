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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		res := solve(n)

		fmt.Println(string(res))
		t--
	}
}

func solve(n int) []byte {
	N := int64(n)
	a := sort.Search(n, func(a int) bool {
		b := int64(a)
		return b*b >= N
	})

	if a*a == n {
		return solve1(n, a)
	}

	if a*(a-1) >= n {
		return solve2(n, a)
	}

	return solve3(n, a)
}

func solve1(n, a int) []byte {
	res := make([]byte, 2*a)
	for i := 0; i < 2*a; i++ {
		if i < a {
			res[i] = 'X'
		} else {
			res[i] = 'D'
		}
	}
	return res
}

func solve2(n, a int) []byte {
	res := make([]byte, 2*a-1)
	for i := 0; i < len(res); i++ {
		if i < a {
			res[i] = 'X'
		} else {
			res[i] = 'D'
		}
	}

	m := a * (a - 1)
	j := a - 1
	for m > n {
		m--
		j++
	}
	if j >= a {
		res[j], res[a-1] = 'X', 'D'
	}
	return res
}

func solve3(n, a int) []byte {
	res := make([]byte, 2*a)
	for i := 0; i < 2*a; i++ {
		if i < a {
			res[i] = 'X'
		} else {
			res[i] = 'D'
		}
	}
	j := a - 1
	m := a * a
	for m > n {
		m--
		j++
	}
	if j >= a {
		res[j], res[a-1] = 'X', 'D'
	}
	return res
}
