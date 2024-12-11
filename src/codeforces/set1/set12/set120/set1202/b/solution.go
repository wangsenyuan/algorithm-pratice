package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)

	var buf bytes.Buffer
	for _, row := range res {
		for i := 0; i < 10; i++ {
			buf.WriteString(fmt.Sprintf("%d ", row[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

const inf = 1 << 30

func solve(s string) [][]int {

	freq := make([][]int, 10)
	for i := 0; i < 10; i++ {
		freq[i] = make([]int, 10)
	}

	for i := 0; i+1 < len(s); i++ {
		a := int(s[i] - '0')
		b := int(s[i+1] - '0')
		freq[a][b]++
	}
	res := make([][]int, 10)
	for i := 0; i < 10; i++ {
		res[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			res[i][j] = inf
		}
	}

	check := func(x int, y int) int {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				res[i][j] = inf
			}
		}

		for v := 0; v < 10; v++ {
			for cx := 0; cx < 10; cx++ {
				for cy := 0; cy < 10; cy++ {
					if cx+cy == 0 {
						continue
					}
					u := (v + cx*x + cy*y) % 10
					res[v][u] = min(res[v][u], cx+cy)
				}
			}
		}

		var sum int

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if freq[i][j] == 0 {
					continue
				}
				if res[i][j] == inf {
					return -1
				}
				sum += (res[i][j] - 1) * freq[i][j]
			}
		}

		return sum
	}

	table := make([][]int, 10)
	for i := 0; i < 10; i++ {
		table[i] = make([]int, 10)
	}
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			table[i][j] = check(i, j)
			table[j][i] = table[i][j]
		}
	}

	return table
}
