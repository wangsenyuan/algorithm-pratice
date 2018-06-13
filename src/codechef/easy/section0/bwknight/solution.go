package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
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
		var n, m int
		scanner.Scan()
		bs := scanner.Bytes()
		j := readInt(bs, 0, &n)
		readInt(bs, j+1, &m)
		res := solve(n, m)
		fmt.Println(res)
	}
}

func solve(n, m int) int64 {
	if n < m {
		return solve(m, n)
	}
	a, b := int64(n), int64(m)
	if b == 1 {
		return a * (a - 1)
	}

	return (a*b)*(a*b-1) - 4*((a-2)*(b-1)+(a-1)*(b-2))
}
