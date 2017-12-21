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

func readNum(scanner *bufio.Scanner) (x int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &x)
	return
}

func readNums(scanner *bufio.Scanner) (n int, m int, x int, y int, l int) {
	scanner.Scan()
	bs := scanner.Bytes()
	i := readInt(bs, 0, &n)
	i = readInt(bs, i+1, &m)
	i = readInt(bs, i+1, &x)
	i = readInt(bs, i+1, &y)
	readInt(bs, i+1, &l)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for i := 0; i < t; i++ {
		n, m, x, y, l := readNums(scanner)
		scanner.Scan()
		ds := scanner.Bytes()
		which, steps := solve(n, m, x, y, l, ds)
		fmt.Printf("%s %d", which, steps)
	}
}

func solve(n, m, x, y, l int, ds []byte) (string, int) {
	board := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		board[i] = make([]int, m+1)
	}
	board[x][y] = 1
	dx, dy := 0, 0
	for i := 0; i < len(ds); i++ {
		if ds[i] == 'U' {
			dy = 1
			dx = 0
		} else if ds[i] == 'R' {
			dy = 0
			dx = 1
		} else if ds[i] == 'D' {
			dy = -1
			dx = 0
		} else {
			dy = 0
			dx = -1
		}
		x += dx
		y += dy
		board[x][y] = i + 2
	}
	k := 1
	for {
		x += dx
		y += dy
		if x < 0 || x >= n || y < 0 || y >= m {
			return "WALL", k - 1
		}
		if board[x][y] > k {
			return "BODY", k - 1
		}
		k++
	}
}
