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

	n, m := readTwoNums(scanner)
	rules := make([][]int, n)

	for i := 0; i < n; i++ {
		rules[i] = readNNums(scanner, 4)
	}

	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		queries[i] = readNNums(scanner, 2)
	}

	res := solve(n, rules, m, queries)

	for i := 0; i < m; i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, rules [][]int, m int, queries [][]int) []int {
	care := make(map[int]map[int]Pair)

	for _, rule := range rules {
		id, attr, val, pri := rule[0], rule[1], rule[2], rule[3]
		if _, found := care[id]; !found {
			care[id] = make(map[int]Pair)
		}

		if prev, found := care[id][attr]; !found || prev.pri <= pri {
			care[id][attr] = Pair{val, pri}
		}
	}

	res := make([]int, m)

	for i := 0; i < m; i++ {
		id, attr := queries[i][0], queries[i][1]
		res[i] = care[id][attr].val
	}
	return res
}

type Pair struct {
	val, pri int
}
