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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		scanner.Scan()
		s := scanner.Text()
		r := solve(s)
		fmt.Printf("Case #%d: %s\n", x, r)
	}
}

func solve(s string) string {
	n := len(s)

	i := n - 1
	for i > 0 && s[i-1] >= s[i] {
		i--
	}
	i--
	// s[i-1] < s[i]

	if i < 0 {
		//no one
		s = "0" + s
		i = 0
		n++
	}

	j := n - 1
	for j > i && s[j] <= s[i] {
		j--
	}
	// s[j] > s[i]
	res := []byte(s)
	res[i], res[j] = res[j], res[i]
	i++
	j = n - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}
