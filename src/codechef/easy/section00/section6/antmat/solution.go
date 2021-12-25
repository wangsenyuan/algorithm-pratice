package main

import (
	"bufio"
	"fmt"
	"os"
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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, m, edges))
	}
}

func solve(n, m int, edges [][]int) int {
	degree := make([]int, n)

	for i := 0; i < m; i++ {
		a, b := edges[i][0]-1, edges[i][1]-1
		degree[a]++
		degree[b]++
	}

	var res int
	for i := 0; i < n; i++ {
		if degree[i] > res {
			res = degree[i]
		}
	}

	if res >= 3 {
		return res
	}

	// try to find a 3 edges cycle
	conn := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		conn[i] = make(map[int]bool)
	}

	for i := 0; i < m; i++ {
		a, b := edges[i][0]-1, edges[i][1]-1
		conn[a][b] = true
		conn[b][a] = true
	}

	//then all nodes will have less than 3 neighbors
	for i := 0; i < n; i++ {
		//find a, b that, a is connected to b
		if len(conn[i]) <= 1 {
			continue
		}
		for a := range edges[i] {
			for b := range edges[i] {
				if a == b {
					continue
				}
				if conn[a][b] {
					return 3
				}
			}
		}
	}
	return res
}
