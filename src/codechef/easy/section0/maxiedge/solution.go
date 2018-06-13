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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)
	for i := 0; i < t; i++ {
		var n int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)
		grid := make([]string, 2)
		scanner.Scan()
		grid[0] = scanner.Text()
		scanner.Scan()
		grid[1] = scanner.Text()
		res := solve(n, grid)
		fmt.Println(res[0])
		fmt.Println(res[1])
	}
}

func solve(n int, grid []string) []string {
	if n == 1 {
		return grid
	}
	cnt := make([]int, 26)
	for i := 0; i < 2; i++ {
		for j := 0; j < len(grid[i]); j++ {
			cnt[int(grid[i][j]-'a')]++
		}
	}
	res := make([][]byte, 2)
	res[0] = make([]byte, len(grid[0]))
	res[1] = make([]byte, len(grid[0]))

	var row, col int
	for i := 0; i < 26; i++ {
		if cnt[i] > 0 && cnt[i]%2 == 0 {
			for k := 0; k < cnt[i]; k++ {
				res[row][col] = byte('a' + i)
				row = (row + 1) % 2
				if row == 0 {
					col++
				}
			}
		}
	}

	for i := 0; i < 26; i++ {
		if cnt[i] > 0 && cnt[i]%2 == 1 {
			for k := 0; k < cnt[i]; k++ {
				res[row][col] = byte('a' + i)
				row = (row + 1) % 2
				if row == 0 {
					col++
				}
			}
		}
	}
	return []string{string(res[0]), string(res[1])}
}
