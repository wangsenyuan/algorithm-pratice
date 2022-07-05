package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	// hint(105)
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		king := readNNums(reader, 2)
		rocks := make([][]int, 2)
		for i := 0; i < 2; i++ {
			rocks[i] = readNNums(reader, 2)
		}
		res := solve(king, rocks)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(king []int, rocks [][]int) bool {

	var first_rock_move func(king []int, first []int, second []int) bool
	var king_move func(king []int, first []int, second []int) bool
	var second_rock_move func(king []int, first []int, second []int) bool

	first_rock_move = func(king, first, second []int) bool {

		if canMoveTo(first, second, []int{king[0], first[1]}) &&
			king_move(king, []int{king[0], first[1]}, second) {
			return true
		}

		if canMoveTo(first, second, []int{first[0], king[1]}) &&
			king_move(king, []int{first[0], king[1]}, second) {
			return true
		}

		return false
	}

	king_move = func(king []int, first []int, second []int) bool {
		r, c := king[0], king[1]

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				u, v := r+i, c+j
				if u < 1 || u > 8 || v < 1 || v > 8 {
					continue
				}
				if u == first[0] && v == first[1] {
					// beat first rock
					if !second_rock_move([]int{u, v}, nil, second) {
						return false
					}
				}
				if u == second[0] && v == second[1] {
					if !second_rock_move([]int{u, v}, first, nil) {
						return false
					}
				}

				if !second_rock_move([]int{u, v}, first, second) {
					return false
				}
			}
		}
		return true
	}

	second_rock_move = func(king []int, first []int, second []int) bool {
		if first != nil {
			if first[0] == king[0] || first[1] == king[1] {
				return true
			}
		}
		if second != nil {
			if second[0] == king[0] || second[1] == king[1] {
				return true
			}
		}
		return false
	}

	return first_rock_move(king, rocks[0], rocks[1]) || first_rock_move(king, rocks[1], rocks[0])
}

func canMoveTo(a, b, c []int) bool {
	// can a go to c
	if a[0] == c[0] {
		// same column
		if b[0] == a[0] {
			return b[1] < min(a[1], c[1]) || max(a[1], c[1]) < b[1]
		}
	}
	if a[1] == c[1] {
		// same row
		if b[1] == a[1] {
			return b[0] < min(a[0], c[0]) || max(a[0], c[0]) < b[0]
		}
	}

	return true
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
