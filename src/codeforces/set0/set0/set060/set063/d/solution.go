package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := process(reader)
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, row := range res {
		buf.WriteString(row)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (res []string, island []int, a []int) {
	island = readNNums(reader, 5)
	a = readNNums(reader, island[4])
	res = solve(island, a)
	return
}

func solve(island []int, parties []int) []string {
	a, b, c, d := island[0], island[1], island[2], island[3]

	if b == d {
		return case1(a+c, b, parties)
	}

	if b > d {
		return case2(a, b, c, d, parties)
	}

	return case3(a, b, c, d, parties)
}

func case3(a int, b int, c int, d int, p []int) []string {
	// b < d
	buf := make([][]byte, d)
	for i := range d {
		buf[i] = make([]byte, a+c)
		for j := range a + c {
			buf[i][j] = '.'
		}
	}
	x, y, dy := d-1, a, 1
	if (d-b)&1 == 0 {
		y = a + c - 1
		dy = -1
	}
	for i, v := range p {
		for v > 0 {
			buf[x][y] = byte(i + 'a')
			y += dy
			if x >= b {
				if y == a+c {
					y = a + c - 1
					dy *= -1
					x--
				} else if y < a {
					y = a
					dy *= -1
					x--
				}
			} else {
				if y == a+c {
					y = a + c - 1
					dy *= -1
					x--
				} else if y < 0 {
					y = 0
					dy *= -1
					x--
				}
			}

			v--
		}
	}
	return convert(buf)
}

func case2(a int, b int, c int, d int, p []int) []string {
	// b > d
	buf := make([][]byte, b)
	for i := range b {
		buf[i] = make([]byte, a+c)
		for j := range a + c {
			buf[i][j] = '.'
		}
	}

	x, y, dy := b-1, 0, 1
	if (b-d)&1 == 1 {
		y = a - 1
		dy = -1
	}

	for i, v := range p {
		for v > 0 {
			buf[x][y] = byte(i + 'a')
			y += dy
			if x >= d {
				// 在突出部
				if y == a {
					x--
					y = a - 1
					dy *= -1
				} else if y < 0 {
					x--
					y = 0
					dy *= -1
				}
			} else {
				if y == a+c {
					x--
					y = a + c - 1
					dy *= -1
				} else if y < 0 {
					x--
					y = 0
					dy *= -1
				}
			}
			v--
		}
	}

	return convert(buf)
}

func case1(n int, m int, a []int) []string {
	buf := make([][]byte, m)
	for i := range m {
		buf[i] = make([]byte, n)
	}
	var x, y int
	dy := 1
	for i, v := range a {
		for v > 0 {
			buf[x][y] = byte(i + 'a')
			y += dy
			if y == n {
				x++
				y = n - 1
				dy = -1
			} else if y < 0 {
				x++
				y = 0
				dy = 1
			}
			v--
		}
	}
	return convert(buf)
}

func convert(buf [][]byte) []string {
	ans := make([]string, len(buf))
	for i, cur := range buf {
		ans[i] = string(cur)
	}
	return ans
}
