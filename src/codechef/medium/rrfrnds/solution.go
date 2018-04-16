package main

import (
	"bufio"
	"os"
	"fmt"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)
	fs := make([][]byte, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		fs[i] = scanner.Bytes()
	}

	fmt.Println(solve(n, fs))
}

func solve(n int, friendship [][]byte) int {
	adj := make([][]int, n)

	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if friendship[i][j] == '0' {
				continue
			}
			x, y := j/32, j%32
			adj[i][x] |= 1 << uint(y)
		}
	}

	p := (n + 31) / 32

	var ans int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			x, y := j/32, j%32
			if adj[i][x]&(1<<uint(y)) > 0 {
				continue
			}

			k := 0
			for k <= p {
				if adj[i][k]&adj[j][k] > 0 {
					break
				}
				k++
			}

			if k <= p {
				ans++
			}
		}
	}

	return ans
}
