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
		var n, p int
		scanner.Scan()
		bs := scanner.Bytes()
		readInt(bs, readInt(bs, 0, &n)+1, &p)
		res, can := solve(n, p)
		if !can {
			fmt.Println("impossible")
		} else {
			fmt.Println(res)
		}
	}
}

func solve(n int, p int) (string, bool) {
	if p <= 2 {
		return "", false
	}

	q := n / p

	var seed string
	if p%2 == 1 {
		as := repeat("a", p/2)
		seed = as + "b" + as
	} else {
		as := repeat("a", p/2-1)
		seed = as + "bb" + as
	}

	return repeat(seed, q), true
}

func repeat(seed string, n int) string {
	var res string
	for i := 0; i < n; i++ {
		res += seed
	}
	return res
}
