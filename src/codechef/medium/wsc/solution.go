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

func readTwoNum(scanner *bufio.Scanner) (a int, b int) {
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
		n, m := readTwoNum(scanner)
		relation := make([][]int, m)
		for i := 0; i < m; i++ {
			relation[i] = readNNums(scanner, 2)
		}
		res := solve(n, m, relation)

		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--
	}
}

func solve(n int, m int, A [][]int) bool {
	if n < 3 || m < 2 {
		return true
	}
	if m > 2 {
		return false
	}
	a, b := make([]int, 2), make([]int, 2)
	a[0] = A[0][0]
	a[1] = A[1][0]
	b[0] = A[0][1]
	b[1] = A[1][1]

	if a[0] != b[1] && a[1] != b[0] && b[0] != b[1] && a[0] != a[1] {
		return false
	}
	return true
}

func solve1(n int, m int, relation [][]int) bool {
	eaten := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		eaten[i] = make(map[int]bool)
	}
	left := make(map[int]bool)
	for i := 0; i < m; i++ {
		x, y := relation[i][0]-1, relation[i][1]-1
		eaten[x][y] = true
		left[x] = true
		left[y] = true
	}

	if len(left) < 3 {
		// no danger
		return true
	}

	if len(left) > 3 {
		return false
	}

	items := make([]int, 3)
	var i int

	for k := range left {
		items[i] = k
		i++
	}

	cache := make([][][]int, 8)
	for i := 0; i < 8; i++ {
		cache[i] = make([][]int, 8)
		for j := 0; j < 8; j++ {
			cache[i][j] = make([]int, 2)
		}
	}
	safe := func(x int) bool {
		if x == 7 {
			return false
		}
		if x == 3 {
			return !eaten[items[0]][items[1]] && !eaten[items[1]][items[0]]
		}
		if x == 5 {
			return !eaten[items[0]][items[2]] && !eaten[items[2]][items[0]]
		}
		if x == 6 {
			return !eaten[items[1]][items[2]] && !eaten[items[2]][items[1]]
		}
		return true
	}
	var dfs func(a, b, c int) bool

	dfs = func(a, b, c int) bool {
		if cache[a][b][c] != 0 {
			return cache[a][b][c] > 0
		}
		cache[a][b][c] = -1
		if c == 0 {
			if a > 0 {
				//in side a
				for i := 0; i < 3; i++ {
					if a&(1<<uint(i)) > 0 && safe(a^(1<<uint(i))) && dfs(a^(1<<uint(i)), b|(1<<uint(i)), 1) {
						// try move i
						cache[a][b][c] = 1
						break
					}
				}
			} else {
				cache[a][b][c] = 1
			}
		} else {
			if b < 7 {
				if safe(b) && dfs(a, b, 0) {
					cache[a][b][c] = 1
				} else {
					//in side b
					for i := 0; i < 3; i++ {
						if b&(1<<uint(i)) > 0 && safe(b^(1<<uint(i))) && dfs(a|(1<<uint(i)), b^(1<<uint(i)), 0) {
							// try move i
							cache[a][b][c] = 1
							break
						}
					}
				}
			} else {
				cache[a][b][c] = 1
			}
		}

		return cache[a][b][c] > 0
	}

	return dfs(7, 0, 0)
}
