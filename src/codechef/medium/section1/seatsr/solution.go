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
		scanner.Scan()
		x := scanner.Bytes()
		scanner.Scan()
		y := scanner.Bytes()
		z := readNNums(scanner, 3)
		a, b, k := z[0], z[1], z[2]
		fmt.Println(solve(x, y, a, b, k))
		t--
	}
}

func solve(s, t []byte, a, b, k int) int {
	INF := k + 1

	m := len(s)
	n := len(t)

	cache := make(map[int]map[int]int)

	var dfs func(x, y, c int) int
	dfs = func(x, y, c int) int {
		if tmp, foundx := cache[x]; foundx {
			if res, foundy := tmp[y]; foundy {
				return res
			}
		} else {
			cache[x] = make(map[int]int)
		}

		if x == 0 {
			return a*y + c
		}
		if y == 0 {
			return a*x + c
		}
		if c > k {
			// no need to process any more
			return INF
		}
		//same
		res := INF
		if s[x-1] == t[y-1] {
			res = dfs(x-1, y-1, c)
		} else {
			res = dfs(x-1, y, c+a)
			res = min(res, dfs(x, y-1, c+a))
			res = min(res, dfs(x-1, y-1, b))
		}
		cache[x][y] = res
		return res
	}

	res := dfs(m, n, 0)
	if res >= INF {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
