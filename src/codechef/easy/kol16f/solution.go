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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		chain := scanner.Bytes()
		res := solve(chain)
		fmt.Printf("Case %d: %d\n", i+1, res)
	}
}

func solve(chain []byte) int {
	n := len(chain)
	a, b := 0, 0
	hasOneSeg := false
	var j = 0
	for i := 1; i < n; i++ {
		if chain[i] == chain[j] {
			continue
		}
		c := i - j
		if c > a {
			a, b = c, a
		} else if c > b {
			b = c
		}
		if c == 1 {
			hasOneSeg = true
		}
		j = i
	}

	k := 0
	for k < j && chain[k] == chain[j] {
		k++
	}

	c := k + n - j
	if c > a {
		a, b = c, a
	} else if c > b {
		b = c
	}

	if a == n {
		if n == 1 {
			return 1
		}
		return n - 1
	}

	if a == 1 {
		if n == 2 {
			return 2
		}
		return 3
	}

	if a == 2 {
		if hasOneSeg || b == 1 {
			return 2
		}
		return 3
	}

	ha := a / 2
	if ha > b {
		return ha
	}
	return b
}
