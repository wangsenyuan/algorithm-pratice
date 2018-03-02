package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for i := 0; i < t; i++ {
		kp, ks := readTwoNums(scanner)
		n := readNum(scanner)
		g := readNNums(scanner, n)
		s := readNNums(scanner, n)
		res := solve(n, g, s, kp, ks)
		fmt.Printf("Case %d: %d\n", i+1, res)
	}
}

const MOD = 1e9 + 7
const MAX = 1e6

func solve(n int, g []int, s []int, kp, ks int) int {

	goal := int64(kp) * int64(ks)

	count := make([]int, MAX+1)

	for _, x := range s {
		count[x]++
	}

	for x := int(MAX) - 1; x > 0; x-- {
		count[x] += count[x+1]
	}

	sort.Ints(g)
	ans := 1
	for i := 0; i < n; i++ {
		need := goal/int64(g[i]) + 1
		if need > MAX {
			ans = 0
			break
		}
		res := count[need] - i
		if res <= 0 {
			ans = 0
			break
		}
		ans = (ans * res) % MOD
	}

	return ans
}
