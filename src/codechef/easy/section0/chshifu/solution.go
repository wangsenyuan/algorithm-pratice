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
		tmp := readNNums(scanner, 2)
		n, m := tmp[0], tmp[1]
		edges := make([][]int, m)
		for j := 0; j < m; j++ {
			edges[j] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, m, edges))
	}
}

func solve(n int, m int, edges [][]int) int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0, m/2)
	}

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	N := 1 << uint(n)

	check := func(k int) bool {
		v := 1<<uint(k) - 1
		for v < N {
			var informed int
			for i := 0; i < n; i++ {
				if v&(1<<uint(i)) > 0 {
					// inform
					informed |= 1 << uint(i)
					for _, v := range graph[i] {
						informed |= 1 << uint(v)
					}
					if informed == N-1 {
						return true
					}
				}
			}

			t := (v | (v - 1)) + 1
			v = t | ((((t & -t) / (v & -v)) >> 1) - 1)
		}
		return false
	}

	left, right := 1, n
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
