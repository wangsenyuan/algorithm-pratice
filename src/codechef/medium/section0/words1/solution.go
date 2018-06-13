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
		n := readNum(scanner)

		words := make([][]byte, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			words[i] = scanner.Bytes()
		}

		res := solve(n, words)

		if res {
			fmt.Println("Ordering is possible.")
		} else {
			fmt.Println("The door cannot be opened.")
		}

		t--
	}
}

func solve(n int, words [][]byte) bool {
	diff := make([]int, 26)

	grid := make([][]int, 26)

	for i := 0; i < 26; i++ {
		grid[i] = make([]int, 26)
	}

	mp := make(map[int]bool)
	for _, word := range words {
		a, b := int(word[0]-'a'), int(word[len(word)-1]-'a')
		diff[b]++
		diff[a]--
		grid[a][b]++
		mp[a] = true
		mp[b] = true
	}

	que := make([]int, 26)
	seen := make([]int, 26)
	bfs := func(s int, label int) int {
		front, tail := 0, 0
		que[tail] = s
		tail++
		seen[s] = label
		for front < tail {
			tt := tail
			for front < tt {
				u := que[front]
				front++
				for v := 0; v < 26; v++ {
					if seen[v] != label && grid[u][v] > 0 {
						seen[v] = label
						que[tail] = v
						tail++
					}
				}
			}
		}

		return tail
	}

	var connected bool
	for i := 0; i < 26 && !connected; i++ {
		connected = len(mp) == bfs(i, i+1)
	}

	if !connected {
		return false
	}

	a, b := -1, -1

	for i := 0; i < 26; i++ {
		if diff[i] == 0 {
			continue
		}
		if diff[i] > 1 || diff[i] < -1 {
			// stuck at i
			return false
		}

		if diff[i] == 1 {
			if a >= 0 {
				// two ends
				return false
			}
			a = i
		} else {
			if b >= 0 {
				// two begins
				return false
			}
			b = i
		}
	}

	if a < 0 && b < 0 {
		return true
	}

	if a >= 0 && b >= 0 {
		return true
	}

	return false
}
