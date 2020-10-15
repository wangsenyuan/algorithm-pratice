package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	B := make([]string, 8)
	for i := 0; i < 8; i++ {
		B[i], _ = reader.ReadString('\n')
	}
	fmt.Println(solve(B))
}

func solve(B []string) int {
	var cnt int
	row := make([]bool, 8)
	col := make([]bool, 8)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if B[i][j] == 'B' {
				cnt++
			} else {
				row[i] = true
				col[j] = true
			}
		}
	}
	if cnt == 0 {
		return 0
	}
	if cnt == 64 {
		return 8
	}

	var res int
	for i := 0; i < 8; i++ {
		if !row[i] {
			res++
		}
		if !col[i] {
			res++
		}
	}
	return res
}
