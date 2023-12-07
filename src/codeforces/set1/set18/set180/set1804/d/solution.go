package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	mat := make([]string, n)
	for i := 0; i < n; i++ {
		mat[i] = readString(reader)[:m]
	}
	res := solve(mat)

	fmt.Printf("%d %d\n", res[0], res[1])
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(mat []string) []int {
	n := len(mat)
	m := len(mat[0])

	calc := func(r int) []int {
		var bright int
		for i := 0; i < m; i++ {
			if mat[r][i] == '1' {
				bright++
			}
		}
		var t2 int
		for i := 0; i < m; {
			if mat[r][i] == '0' {
				i++
				continue
			}
			j := i
			for i < m && mat[r][i] == '1' {
				i++
			}
			t2 += (i - j) / 2
		}
		if t2 > m/4 {
			t2 = m / 4
		}
		tmp := []int{bright - t2, 0}

		t2 = 0
		for i := 0; i < m; {
			// 要找到下一个连续的两个亮的窗户
			j := i
			for i+1 < m && (mat[r][i] == '0' || mat[r][i+1] == '0') {
				i++
			}
			if i == j {
				i++
				continue
			}
			i++
			t2 += (i - j) / 2
		}
		t2 = m/4 - t2
		if t2 < 0 {
			t2 = 0
		}
		tmp[1] = bright - t2
		return tmp
	}

	res := make([]int, 2)

	for i := 0; i < n; i++ {
		tmp := calc(i)
		res[0] += tmp[0]
		res[1] += tmp[1]
	}

	return res
}
