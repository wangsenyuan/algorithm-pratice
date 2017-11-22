package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int

	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		var n, m int
		scanner.Scan()
		bs := scanner.Bytes()
		x := readInt(bs, 0, &n)
		readInt(bs, x+1, &m)
		ss := make([][]int, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			bs = scanner.Bytes()
			var cnt int
			x = readInt(bs, 0, &cnt)

			ss[j] = make([]int, cnt)
			k := 0
			for x < len(bs) {
				x = readInt(bs, x+1, &ss[j][k])
				k++
			}
			// fmt.Printf("echo: %v\n", ss[i])
		}
		res := solve(n, m, ss)
		fmt.Println(res)
	}
}

func solve(n, m int, ss [][]int) int {
	xs := make([]int, n)

	for i := 0; i < m; i++ {
		s := ss[i]
		// fmt.Printf("%v\n", s)
		for _, num := range s {
			xs[num] |= (1 << uint(i))
		}
	}

	y := make(map[int]bool)

	for _, x := range xs {
		y[x] = true
	}

	return len(y)
}

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
