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
	for i := 0; i < t; i++ {
		n := readNum(scanner)
		scanner.Scan()
		bs := scanner.Bytes()
		ss := make([]string, n)
		pos := 0
		for i := 0; i < n; i++ {
			j := pos
			for j < len(bs) && bs[j] != ' ' {
				j++
			}
			ss[i] = string(bs[pos:j])
			pos = j + 1
		}
		fmt.Println(solve(n, ss))
	}
}

func solve(n int, ss []string) int {
	count := make([]int, 4)
	// 0: a, 1: b, 3: ab, b: ba

	for _, s := range ss {
		if s == "a" {
			count[0]++
		} else if s == "b" {
			count[1]++
		} else if s == "aa" {
			count[0]++
		} else if s == "bb" {
			count[1]++
		} else if s == "ab" {
			count[2]++
		} else if s == "ba" {
			count[3]++
		}
	}
	// concat n a's, to get on a
	if count[0] > 0 {
		count[0] = 1
	}
	if count[1] > 0 {
		count[1] = 1
	}
	if count[2] > 0 || count[3] > 0 {
		if count[2] == 0 {
			return 2 * count[3]
		}
		if count[3] == 0 {
			return 2 * count[2]
		}

		// ab & ba -> aba (1 -> 2 * 2 - 1)
		// aba & aba -> abaaba -> ababa (2 -> 2 * 2 * 2 - 3)
		// ababa & ababa -> ababaababa -> ababababa (4 -> 2 * 2 * 4 - 7)
		if count[2] == count[3] {
			return 2*count[2] + 1
		}
		return max(count[2], count[3]) * 2
	}
	return count[0] + count[1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
