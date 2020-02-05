package main

import (
	"bufio"
	"fmt"
	"os"
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
		color := readNNums(scanner, n)
		fmt.Println(solve(n, color))
	}
}

func solve(n int, color []int) int {
	pos := make(map[int]int)
	pos[0] = -1

	var best int
	var cur int

	for i := 0; i < n; i++ {
		color[i]--
		cur ^= 1 << uint(color[i])

		for t := 0; t <= 30; t++ {
			v := 1 << uint(t)
			x := v ^ cur
			if j, found := pos[x]; found {
				best = max(best, i-j)
			}
		}

		if _, found := pos[cur]; !found {
			pos[cur] = i
		}
	}

	return best / 2
}

func solve1(n int, color []int) int {
	pos := make(map[int]int)
	pos[0] = -1

	var best int

	cnt := make([]int, 31)
	for i := 0; i < n; i++ {
		color[i]--
		cnt[color[i]] = (cnt[color[i]] + 1) & 1

		var flag int

		for j := 0; j <= 30; j++ {
			if cnt[j] > 0 {
				flag |= 1 << uint(j)
			}
		}

		for j := 0; j <= 30; j++ {
			var tmp int
			if cnt[j] > 0 {
				tmp = flag & ^(1 << uint(j))
				// use color j as roof
			} else {
				tmp = flag | (1 << uint(j))
			}
			if ii, found := pos[tmp]; found {
				best = max(best, i-ii)
			}
		}

		if _, found := pos[flag]; !found {
			pos[flag] = i
		}
	}

	return best / 2
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
