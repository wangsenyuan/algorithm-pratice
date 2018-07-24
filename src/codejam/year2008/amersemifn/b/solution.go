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

func main() {
	f, err := os.Open("./B-large-practice.in")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		n := readNum(scanner)

		s := readNNums(scanner, n)

		res := solve(n, s)

		if res < 0 {
			fmt.Printf("Case #%d: UNKNOWN\n", i)
		} else {
			fmt.Printf("Case #%d: %d\n", i, res)
		}
	}
}

func solve(n int, s []int) int {
	d := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		d[n-1-i-1] = (s[i+1] - s[i] + 10007) % 10007
	}

	s1, s0, delta := 0, 1, 2
	for {
		if s1+delta >= len(d) {
			return -1
		}
		ok := true
		for i := s0 + delta; i < len(d); i += delta {
			if d[i] != d[s0] {
				ok = false
				break
			}
		}
		var ok2 bool
		for i := s1; i < len(d); i += delta {
			if d[i] != d[s1] {
				ok2 = true
				break
			}
		}

		if ok && ok2 {
			return (s[n-1] + d[s0]) % 10007
		}
		delta <<= 1
		s1 = s1<<1 | 1
		s0 = s0<<1 | 1
	}
}
