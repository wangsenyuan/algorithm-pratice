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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		scanner.Scan()
		a := scanner.Text()
		scanner.Scan()
		b := scanner.Text()
		scanner.Scan()
		c := scanner.Text()
		fmt.Println(solve(a, b, c))
	}
}

const MOD = 10000000007
const P = 31

func solve(a, b, c string) int {
	if search(a, c) || search(b, c) {
		return len(a) + len(b)
	}

	ca := getLps(c + "*" + a)
	bc := getLps(b + "*" + c)

	ans := len(c)
	ans -= ca[len(ca)-1]
	ans -= bc[len(bc)-1]
	return ans + len(a) + len(b)
}

func search(txt, pat string) bool {
	lps := getLps(pat)

	var i, j int

	for i < len(txt) {
		if pat[j] == txt[i] {
			j++
			i++
		}
		if j == len(pat) {
			return true
		}
		if i < len(txt) && pat[j] != txt[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return false
}

func getLps(pat string) []int {
	n := len(pat)
	res := make([]int, n)
	var i = 1
	var l = 0
	for i < n {
		if pat[i] == pat[l] {
			l++
			res[i] = l
			i++
		} else {
			if l != 0 {
				l = res[l-1]
			} else {
				res[i] = 0
				i++
			}
		}
	}
	return res
}
