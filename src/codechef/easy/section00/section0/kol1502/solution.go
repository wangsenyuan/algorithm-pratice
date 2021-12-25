package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)
	for i := 0; i < t; i++ {
		var n int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)
		ingredients := make([]int, n)
		scanner.Scan()
		bs := scanner.Bytes()
		pos := -1
		for j := 0; j < n; j++ {
			pos = readInt(bs, pos+1, &ingredients[j])
		}
		fmt.Println(solve(n, ingredients))
	}
}

const MOD = 10000000

func solve(n int, ingredients []int) int64 {
	var sum int
	for i := 0; i < n; i++ {
		sum += ingredients[i]
	}
	g := make([]int64, sum+1)

	g[0] = 1

	for i := 0; i < n; i++ {
		for state := sum - ingredients[i]; state >= 0; state-- {
			if g[state] > 0 {
				g[state+ingredients[i]] = (g[state+ingredients[i]] + g[state]) % MOD
			}
		}
	}

	var res int64
	for state := 0; state <= sum; state++ {
		res = (res + int64(abs(sum-2*state))*g[state]) % MOD
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
