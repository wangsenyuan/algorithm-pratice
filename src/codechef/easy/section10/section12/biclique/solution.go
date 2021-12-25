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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	firstLine := readNNums(scanner, 3)
	n, m, k := firstLine[0], firstLine[1], firstLine[2]
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(scanner, 2)
	}
	res := solve(n, m, k, edges)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(n int, m int, k int, edges [][]int) bool {
	count := make([][]int, n)
	for i := 0; i < n; i++ {
		count[i] = make([]int, n)
	}

	neighbors := make([][]int, n)
	for i := 0; i < n; i++ {
		neighbors[i] = make([]int, 0, 10)
	}

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		neighbors[a] = append(neighbors[a], b)
		neighbors[b] = append(neighbors[b], a)
	}

	for y := 0; y < n; y++ {
		xx := neighbors[y]
		for i := 0; i < len(xx); i++ {
			x := xx[i]
			for j := 0; j < len(xx); j++ {
				z := xx[j]
				if x < z {
					count[x][z]++
					if count[x][z] == k {
						return true
					}
				}
			}
		}
	}

	return false
}
