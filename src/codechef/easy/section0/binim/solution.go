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
		var n int
		pos := readInt(scanner.Bytes(), 0, &n)
		firstPlayer := scanner.Bytes()[pos+1:]
		stones := make([][]byte, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			stones[i] = scanner.Bytes()
		}
		res := solve(n, stones, firstPlayer)

		if res {
			fmt.Println(string(firstPlayer))
		} else {
			fmt.Println(opponent(string(firstPlayer)))
		}
		t--
	}
}

func opponent(s string) string {
	if s == "Dee" {
		return "Dum"
	}
	return "Dee"
}

func solve(n int, stones [][]byte, firstPlayer []byte) bool {
	cnts := make([]int, 2)
	for i := 0; i < n; i++ {
		stone := stones[i]
		var c int
		for j := 0; j < len(stone); j++ {
			if stone[j] == stone[0] {
				c++
			}
		}
		cnts[int(stone[0]-'0')] = c
	}

	if firstPlayer[1] == 'e' {
		// Dee
		return cnts[0] > cnts[1]
	}
	return cnts[1] > cnts[0]
}
