package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	room := make([]string, n)
	for i := 0; i < n; i++ {
		room[i] = readString(reader)[:m]
	}

	res := solve(room)

	fmt.Println(res)
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

func solve(room []string) int {
	n := len(room)
	m := len(room[0])

	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sum[i][j] = int(room[i][j] - '0')
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}

	get := func(a, b, c, d int) int {
		res := sum[c][d]
		if a > 0 {
			res -= sum[a-1][d]
		}
		if b > 0 {
			res -= sum[c][b-1]
		}
		if a > 0 && b > 0 {
			res += sum[a-1][b-1]
		}
		return res
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for u := i; u < n; u++ {
				for v := j; v < m; v++ {
					tmp := get(i, j, u, v)
					if tmp > 0 {
						break
					}
					res = max(res, 2*(u-i+1)+2*(v-j+1))
				}
			}
		}
	}
	return res
}
