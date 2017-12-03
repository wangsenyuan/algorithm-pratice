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

var r []int

func init() {
	r = make([]int, 1001)
	nu := make([]bool, 1001)
	di := make([]bool, 1001)
	mn, md := 0, 0
	for mn < 1001 && md < 1001 && mn+md < 1001 {
		for mn < 1001 && nu[mn] {
			mn++
		}
		if mn > 1000 {
			break
		}
		for md < 1001 && di[md] {
			md++
		}

		if md > 1000 || mn+md > 1000 {
			break
		}
		r[mn] = mn + md
		nu[mn] = true
		di[md] = true
		nu[mn+md] = true
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var m, n, p, q int
		scanner.Scan()
		bs := scanner.Bytes()
		x := readInt(bs, 0, &m)
		x = readInt(bs, x+1, &n)
		x = readInt(bs, x+1, &p)
		readInt(bs, x+1, &q)
		m -= p
		n -= q
		if m > n {
			m, n = n, m
		}
		if r[m] == n {
			fmt.Println("Bob")
		} else {
			fmt.Println("Alice")
		}
	}
}
