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
		n := readNum(scanner)
		src := make([][]int, n)
		for i := 0; i < n; i++ {
			src[i] = readNNums(scanner, n)
		}
		dst := make([][]int, n)
		for i := 0; i < n; i++ {
			dst[i] = readNNums(scanner, n)
		}
		can, row, col := solve(n, src, dst)
		if !can {
			fmt.Println(-1)
		} else if len(row) == 0 && len(col) == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(len(row))
			for _, idx := range row {
				fmt.Println(idx)
			}
			fmt.Println(len(col))
			for _, idx := range col {
				fmt.Println(idx)
			}
		}
		t--
	}
}
func solve(n int, src, dst [][]int) (bool, []int, []int) {
	M := make([][]int, n)
	for i := 0; i < n; i++ {
		M[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if src[i][j] != dst[i][j] {
				M[i][j] = 1
			}
		}
	}
	rows := make([]int, n)
	cols := make([]int, n)

	rows[0] = 0
	for j := 0; j < n; j++ {
		if M[0][j] != rows[0] {
			cols[j] = 1
		}
	}
	for i := 0; i < n; i++ {
		if M[i][0] != cols[0] {
			rows[i] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			x := rows[i] != cols[j]
			y := M[i][j] > 0
			if x != y {
				return false, nil, nil
			}
		}
	}

	resultSet := 1
	var cr, cc int
	for i := 0; i < n; i++ {
		cr += rows[i]
		cc += cols[i]
	}

	if cr+cc > n {
		resultSet = 0
		cr = n - cr
		cc = n - cc
	}
	if cr+cc == 0 {
		return true, nil, nil
	}

	rr := make([]int, 0, cr)
	rc := make([]int, 0, cc)

	for i := 0; i < n; i++ {
		if rows[i] == resultSet {
			rr = append(rr, i)
		}
		if cols[i] == resultSet {
			rc = append(rc, i)
		}
	}

	return true, rr, rc
}

func solve1(n int, src, dst [][]int) (bool, []int, []int) {

	find := func(flippedCols []bool, rows []bool) int {
		var ans int

		for j := 0; j < n; j++ {
			if flippedCols[j] {
				ans++
			}
		}
		for i := 1; i < n; i++ {
			var j int
			for j < n {
				x := src[i][j]
				if flippedCols[j] {
					x = 1 - x
				}
				if x != dst[i][j] {
					break
				}
				j++
			}
			if j == n {
				continue
			}
			j = 0
			for j < n {
				x := 1 - src[i][j]
				if flippedCols[j] {
					x = 1 - x
				}
				if x != dst[i][j] {
					break
				}
				j++
			}
			if j < n {
				return -1
			}
			rows[i] = true
			ans++
		}

		return ans
	}

	cols0 := make([]bool, n)
	rows0 := make([]bool, n)
	for j := 0; j < n; j++ {
		if src[0][j] != dst[0][j] {
			cols0[j] = true
		}
	}
	row0NotFlip := find(cols0, rows0)

	rows1 := make([]bool, n)
	cols1 := make([]bool, n)
	for j := 0; j < n; j++ {
		if src[0][j] == dst[0][j] {
			cols1[j] = true
		}
	}
	rows1[0] = true
	row0Flip := find(cols1, rows1)

	if row0NotFlip < 0 && row0Flip < 0 {
		return false, nil, nil
	}
	if row0NotFlip < 0 {
		return true, compact(rows1), compact(cols1)
	}

	if row0Flip < 0 {
		return true, compact(rows0), compact(cols0)
	}

	if row0Flip+1 <= row0NotFlip {
		return true, compact(rows1), compact(cols1)
	}

	return true, compact(rows0), compact(cols0)
}

func compact(row []bool) []int {
	res := make([]int, len(row))
	var j int
	for i := 0; i < len(row); i++ {
		if row[i] {
			res[j] = i
			j++
		}
	}
	return res[:j]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
