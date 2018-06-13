package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt64(bs []byte, from int, val *int64) int {
	var tmp int64
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int64(bs[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int64
	scanner.Scan()
	readInt64(scanner.Bytes(), 0, &t)
	for i := 0; int64(i) < t; i++ {
		var a, n, m int64
		scanner.Scan()
		x := readInt64(scanner.Bytes(), 0, &a)
		x = readInt64(scanner.Bytes(), x+1, &n)
		readInt64(scanner.Bytes(), x+1, &m)
		fmt.Println(solve(a, n, m))
	}
}

func solve(a, n, m int64) int64 {
	return (a * pow(a, n-1, m)) % m
}

func pow(a, n, m int64) int64 {
	b := int64(1)
	for a > 0 {
		b *= 10
		a /= 10
	}

	var tmp, ans int64
	tmp = b
	for n > 0 {
		if n&1 == 1 {
			ans = (ans*b + tmp) % m
		}
		tmp = (tmp + b*tmp) % m
		b = (b * b) % m
		n >>= 1
	}
	return (ans + 1) % m
}
