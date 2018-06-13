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
		var n, a, b int
		scanner.Scan()
		bs := scanner.Bytes()
		x := readInt(bs, 0, &a)
		x = readInt(bs, x+1, &b)
		readInt(bs, x+1, &n)
		res := solve(n, a, b)
		for i := 0; i < n; i++ {
			fmt.Printf("%d", res[i])
			if i < n-1 {
				fmt.Printf(" ")
			} else {
				fmt.Println()
			}
		}
	}
}

func solve(n, a, b int) []int {
	res := make([]int, n)
	xs := make([]bool, n*n-1)

	x := 0
	for i := 0; i < n; i++ {
		x++
		for x < len(xs) && xs[x] {
			x++
		}
		res[i] = x
		for j := 0; j <= i; j++ {
			z := a*res[j] - b*res[i]
			if z >= 0 && z < len(xs) {
				xs[z] = true
			}
			z = a*res[i] - b*res[j]
			if z >= 0 && z < len(xs) {
				xs[z] = true
			}
		}

	}

	return res
}
