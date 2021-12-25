package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		n := readNum(scanner)
		words := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			words[i] = scanner.Text()
		}
		res := solve(n, words)
		for i := 0; i < 4; i++ {
			fmt.Println(res[i*4 : (i+1)*4])
		}
		fmt.Println()
	}
}

const fail = "gridsnotpossible"

func solve(n int, words []string) string {
	ws := removeDuplicates(words)

	if len(ws) > 8 {
		return fail
	}

	res := make([]string, 0, 16)
	mem := make([]byte, 16)

	fillEmpty := func(rowMask, colMask int) {
		for i := 0; i < 4; i++ {
			if rowMask&(1<<uint(i)) == 0 {
				//this row is not filled
				for j := 0; j < 4; j++ {
					if colMask&(1<<uint(j)) == 0 {
						// this col is not filled
						mem[i*4+j] = 'A'
					}
				}
			}
		}
	}

	canPutInRow := func(word string, row int, colMask int) bool {
		for j := 0; j < 4; j++ {
			if colMask&(1<<uint(j)) > 0 {
				// this col has word
				k := row*4 + j
				if mem[k] != word[j] {
					//confict
					return false
				}
			}
		}
		return true
	}

	putInRow := func(word string, row int) {
		for j := 0; j < 4; j++ {
			k := row*4 + j
			mem[k] = word[j]
		}
	}

	canPutInCol := func(word string, col int, rowMask int) bool {
		for i := 0; i < 4; i++ {
			if rowMask&(1<<uint(i)) > 0 {
				k := i*4 + col
				if mem[k] != word[i] {
					return false
				}
			}
		}
		return true
	}

	putInCol := func(word string, col int) {
		for i := 0; i < 4; i++ {
			k := i*4 + col
			mem[k] = word[i]
		}
	}
	//
	var loop func(rowMask, colMask int, idx int)

	loop = func(rowMask, colMask int, idx int) {
		if idx == len(ws) {
			fillEmpty(rowMask, colMask)
			res = append(res, string(mem))
			return
		}

		if rowMask != 15 {
			//has empty row to try word[idx]
			for i := 0; i < 4; i++ {
				if rowMask&(1<<uint(i)) == 0 {
					word := ws[idx]
					if canPutInRow(word, i, colMask) {
						putInRow(word, i)
						loop(rowMask|(1<<uint(i)), colMask, idx+1)
					}
					rw := reverse(word)
					if canPutInRow(rw, i, colMask) {
						putInRow(rw, i)
						loop(rowMask|(1<<uint(i)), colMask, idx+1)
					}
				}
			}
		}

		if colMask != 15 {
			//has empty col to try word[idx]
			for j := 0; j < 4; j++ {
				if colMask&(1<<uint(j)) == 0 {
					word := ws[idx]

					if canPutInCol(word, j, rowMask) {
						putInCol(word, j)
						loop(rowMask, colMask|(1<<uint(j)), idx+1)
					}

					rw := reverse(word)

					if canPutInCol(rw, j, rowMask) {
						putInCol(rw, j)
						loop(rowMask, colMask|(1<<uint(j)), idx+1)
					}
				}
			}
		}
	}

	loop(0, 0, 0)

	if len(res) == 0 {
		return fail
	}

	sort.Strings(res)

	return res[0]
}

func removeDuplicates(words []string) []string {
	ss := make(map[string]bool)

	for _, word := range words {
		s := word
		rs := reverse(word)
		if rs < s {
			s = rs
		}
		if ss[s] {
			continue
		}
		ss[s] = true
	}

	ws := make([]string, 0, len(ss))
	for w := range ss {
		ws = append(ws, w)
	}

	return ws
}

func reverse(word string) string {
	bs := []byte(word)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
