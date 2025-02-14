package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0], res[1], res[2])
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	firstRow := readNNums(reader, n)
	firstCol := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		firstCol[i] = readNum(reader)
	}
	return solve(n, firstRow, firstCol)
}

func solve(n int, firstRow []int, firstCol []int) []int {
	if n <= 4 {
		return bruteForce(n, firstRow, firstCol)
	}
	cnt := make([]int, 3)
	a := make([][]int, 4)
	for i := 0; i < 4; i++ {
		a[i] = make([]int, n)
		if i == 0 {
			copy(a[i], firstRow)
		} else {
			a[i][0] = firstCol[i-1]
			for j := 1; j < n; j++ {
				a[i][j] = 0
				if a[i-1][j] == 0 || a[i][j-1] == 0 {
					a[i][j]++
				}
				if a[i][j] == 1 && (a[i-1][j] == 1 || a[i][j-1] == 1) {
					a[i][j]++
				}
			}
		}
		for j := 0; j < n; j++ {
			cnt[a[i][j]]++
			if i == 3 && j > 3 {
				cnt[a[i][j]] += n - j - 1
			}
		}
	}

	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, 4)
		if i == 0 {
			copy(b[i], firstRow)
		} else {
			b[i][0] = firstCol[i-1]
			for j := 1; j < 4; j++ {
				b[i][j] = 0
				if b[i-1][j] == 0 || b[i][j-1] == 0 {
					b[i][j]++
				}
				if b[i][j] == 1 && (b[i-1][j] == 1 || b[i][j-1] == 1) {
					b[i][j]++
				}
			}
		}
		if i >= 4 {
			for j := 0; j < 4; j++ {
				cnt[b[i][j]]++
				if j == 3 {
					cnt[b[i][j]] += n - i - 1
				}
			}
		}
	}

	cnt[b[3][3]] += n - 4

	return cnt
}

func bruteForce(n int, firstRow []int, firstCol []int) []int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}

	copy(a[0], firstRow)
	for i := 1; i < n; i++ {
		a[i][0] = firstCol[i-1]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			a[i][j] = 0
			if a[i-1][j] == 0 || a[i][j-1] == 0 {
				a[i][j]++
			}
			if a[i][j] == 1 && (a[i-1][j] == 1 || a[i][j-1] == 1) {
				a[i][j]++
			}
		}
	}

	res := make([]int, 3)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[a[i][j]]++
		}
	}
	return res
}
