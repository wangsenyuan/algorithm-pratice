package main

import (
	"bufio"
	"os"
	"fmt"
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

func readOne(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readFour(scanner *bufio.Scanner) (a int, b int, c int, d int) {
	scanner.Scan()
	bs := scanner.Bytes()
	x := readInt(bs, 0, &a)
	x = readInt(bs, x+1, &b)
	x = readInt(bs, x+1, &c)
	readInt(bs, x+1, &d)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readOne(scanner)

	for i := 0; i < t; i++ {
		n, m, r, k := readFour(scanner)
		fmt.Println(solve(n, m, r, k))
	}

}

const MOD = 1000000007

func solve(n, m, r, k int) int64 {
	if r == 0 {
		return int64(m)
	}

	if m <= r || n == 1 {
		return 0
	}

	a := pow(r+1, n)
	b := pow(r, n)
	c := pow(r-1, n)
	d := int64(m - r)

	return (d * (a - 2*b + c)) % MOD
}

func pow(a, b int) int64 {
	if b == 0 {
		return 1
	}
	c := pow(a, b/2)
	d := (c * c) % MOD
	if b%2 == 1 {
		return (int64(a) * d) % MOD
	}
	return d % MOD
}
