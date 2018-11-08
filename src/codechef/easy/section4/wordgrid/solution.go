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

	if len(ss) > 8 {
		// at most 8 words can appear in the grid
		return fail
	}

	ws := make([]string, len(ss)+4)
	var p int
	for s := range ss {
		ws[p] = s
		p++
	}

	if p <= 4 {
		j := 4 - p
		for i := 0; i < j; i++ {
			ws[p] = "AAAA"
			p++
		}
	}
	ws = ws[:p]

	// then pick 4 rows, and check left can fit in the cols
	res := make([]string, 0, 10)
	var loop func(mask int, cnt int, putInMem func(string, int), find func(string, string) bool)

	check := func(mask int, str string, find func(string, string) bool) bool {
		for i := 0; i < p; i++ {
			if mask&(1<<uint(i)) == 0 {
				// not used
				if !find(ws[i], str) {
					return false
				}
			}
		}
		return true
	}

	mem := make([]byte, 16)

	loop = func(mask int, cnt int, putInMem func(string, int), find func(string, string) bool) {
		if cnt == 16 {
			str := string(mem)
			if check(mask, str, find) {
				res = append(res, str)
			}
			return
		}

		for i := 0; i < p; i++ {
			if mask&(1<<uint(i)) == 0 {
				putInMem(ws[i], cnt/4)
				loop(mask|(1<<uint(i)), cnt+4, putInMem, find)
				putInMem(reverse(ws[i]), cnt/4)
				loop(mask|(1<<uint(i)), cnt+4, putInMem, find)
			}
		}
	}

	findInCol := func(word, str string) bool {
		for j := 0; j < 4; j++ {
			can := true
			for i := 0; i < 4 && can; i++ {
				can = word[i] == str[i*4+j]
			}
			if can {
				return true
			}
			can = true
			for i := 0; i < 4 && can; i++ {
				can = word[3-i] == str[i*4+j]
			}
			if can {
				return true
			}
		}
		return false
	}

	findInRow := func(word, str string) bool {
		for i := 0; i < 4; i++ {
			can := true
			for j := 0; j < 4 && can; j++ {
				can = word[j] == str[i*4+j]
			}
			if can {
				return true
			}
			can = true
			for j := 0; j < 4 && can; j++ {
				can = word[3-j] == str[i*4+j]
			}
			if can {
				return true
			}
		}
		return false
	}

	putInRow := func(word string, row int) {
		for j := 0; j < 4; j++ {
			mem[row*4+j] = word[j]
		}
	}

	putInCol := func(word string, col int) {
		for i := 0; i < 4; i++ {
			mem[i*4+col] = word[i]
		}
	}

	loop(0, 0, putInRow, findInCol)

	loop(0, 0, putInCol, findInRow)

	if len(res) == 0 {
		return fail
	}

	sort.Strings(res)

	return res[0]
}

func reverse(word string) string {
	bs := []byte(word)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
