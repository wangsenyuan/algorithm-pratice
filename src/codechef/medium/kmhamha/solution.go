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

	for t > 0 {
		n := readNum(scanner)
		pairs := make([][]int, n)
		for i := 0; i < n; i++ {
			pairs[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, pairs))
		t--
	}
}

func solve(n int, pairs [][]int) int {
	rows := make(map[int]map[int]bool)

	for _, pair := range pairs {
		x, y := pair[0], pair[1]
		if _, found := rows[x]; !found {
			rows[x] = make(map[int]bool)
		}

		rows[x][y] = true
	}

	used := make(map[int]int)

	var dfs func(r int) bool
	var step int
	steps := make(map[int]int)
	dfs = func(r int) bool {
		for c := range rows[r] {
			if steps[c] == step {
				continue
			}
			steps[c] = step
			if rr, found := used[c]; !found || dfs(rr) {
				used[c] = r
				return true
			}
		}
		return false
	}
	var ans int
	for r := range rows {
		step++
		if dfs(r) {
			ans++
		}
	}
	return ans
}
