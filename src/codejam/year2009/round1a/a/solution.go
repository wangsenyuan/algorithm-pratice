package main

import (
	"bufio"
	"os"
	"fmt"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		var pos int
		var x int
		B := make([]int, 0, 10)
		scanner.Scan()
		bs := scanner.Bytes()
		for pos < len(bs) {
			pos = readInt(bs, pos, &x)
			B = append(B, x)
			pos++
		}
		res := solve(B)
		fmt.Printf("Case #%d: %d\n", i, res)
	}
}

const N = 20000000

var cache [][]int

func init() {
	cache = make([][]int, N+1)
	for i := 0; i <= N; i++ {
		cache[i] = make([]int, 11)
	}
}

func happy(x, b int) bool {
	if x == 1 {
		return true
	}
	if cache[x][b] == 2 {
		// visit it again
		return false
	}
	if cache[x][b] == 0 {
		cache[x][b] = 2
		y := x
		tmp := 0
		for y > 0 {
			k := y % b
			y /= b
			tmp += k * k
		}
		if happy(tmp, b) {
			cache[x][b] = 1
		} else {
			cache[x][b] = -1
		}
	}

	return cache[x][b] > 0
}

func solve(B []int) int {

	for x := 2; x <= N; x++ {
		can := true
		for j := 0; j < len(B) && can; j++ {
			can = happy(x, B[j])
		}
		if can {
			return x
		}
	}
	return -1
}
