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
		G := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			G[i] = scanner.Text()
		}
		res := solve(n, m, G)
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func solve(n int, m int, G []string) bool {

	check := func(x, y int, dx, dy int) bool {
		u, v := x+dx, y+dy
		if u < 0 || u >= n || v < 0 || v >= m {
			//out of bounds
			return false
		}
		if G[u][y] == '#' || G[x][v] == '#' || G[u][v] == '#' {
			return false
		}
		return true
	}

	topLeft := func(x, y int) bool {
		return check(x, y, 1, 1)
	}

	topRight := func(x, y int) bool {
		return check(x, y, 1, -1)
	}

	bottomLeft := func(x, y int) bool {
		return check(x, y, -1, 1)
	}

	bottomRight := func(x, y int) bool {
		return check(x, y, -1, -1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] == '.' {
				// cell (i, j) at top-left
				can := topLeft(i, j) || topRight(i, j) || bottomLeft(i, j) || bottomRight(i, j)
				if !can {
					return false
				}
			}
		}
	}
	return true
}
