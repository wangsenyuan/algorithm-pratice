package main

import (
	"bufio"
	"fmt"
	"os"
)

var w map[int]map[int]bool

func init() {
	w = make(map[int]map[int]bool)
	for i := 1; i <= 100; i++ {
		w[i] = make(map[int]bool)
		iii := i * i * i
		for j := 0; j < iii; j++ {
			k := (j * j * j) % iii
			w[i][k] = true
		}
	}
}

func readInt(bs []byte, from int, val *int) int {
	var tmp int
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
		scanner.Scan()
		var k, c int
		x := readInt(scanner.Bytes(), 0, &k)
		readInt(scanner.Bytes(), x+1, &c)
		if solve(k, c) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func solve(k, c int) bool {
	if c == 0 {
		return true
	}

	return w[k][k*k*k-c]
}
