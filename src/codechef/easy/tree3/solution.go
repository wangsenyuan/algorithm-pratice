package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = sign * tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readPair(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	bs := scanner.Bytes()
	x := readInt(bs, 0, &a)
	readInt(bs, x+1, &b)
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
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}

		res := solve(n, edges)
		if len(res) == 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			for _, tmp := range res {
				fmt.Printf("%d %d %d %d\n", tmp[0], tmp[1], tmp[2], tmp[3])
			}
		}

		t--
	}
}

func solve(n int, edges [][]int) [][]int {
	if n < 4 {
		return nil
	}
	neighbors := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		neighbors[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		neighbors[a][b] = true
		neighbors[b][a] = true
	}

	res := make([][]int, 0, (n-1)/3)

	var dfs func(p, u int) int

	dfs = func(p, u int) int {
		arr := make([]int, 3)
		var i int
		for v := range neighbors[u] {
			if p == v {
				continue
			}

			tmp := dfs(u, v)
			if tmp == 0 {
				// can't find a solution for child v
				return 0
			} else if tmp == 1 {
				arr[i] = v
				i++
			}
			// tmp == 2, means v preceeds u
			if i == 3 {
				res = append(res, []int{u + 1, arr[0] + 1, arr[1] + 1, arr[2] + 1})
				i = 0
			}
		}
		if i == 1 {
			return 0
		}
		if i == 2 {
			// has to use parent
			if p < 0 {
				return 0
			}
			arr[i] = p
			res = append(res, []int{u + 1, arr[0] + 1, arr[1] + 1, arr[2] + 1})
			return 2
		}
		return 1
	}

	ans := dfs(-1, 0)
	if ans == 0 {
		return nil
	}

	return res
}
