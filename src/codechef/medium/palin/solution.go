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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for i := 0; i < t; i++ {
		scanner.Scan()
		s := scanner.Bytes()
		res := solve(s)
		fmt.Println(res)
	}
}

func solve(s []byte) string {
	n := len(s)

	flag := true

	for i := 0; i < n; i++ {
		if s[i] != '9' {
			flag = false
			break
		}
	}
	if flag {
		for i := 0; i < n-1; i++ {
			s[i] = '0'
		}
		s[n-1] = '1'
		return "1" + string(s)
	}

	a, b := 0, n-1
	for a < b {
		if s[a] > s[b] {
			flag = true
		} else if s[a] < s[b] {
			flag = false
		}
		s[b] = s[a]
		b--
		a++
	}
	if a > b {
		a--
		b++
	}
	if !flag {
		for s[a] == '9' {
			s[a] = '0'
			s[b] = '0'
			a--
			b++
		}
		s[a]++
		s[b] = s[a]
	}
	return string(s)
}
