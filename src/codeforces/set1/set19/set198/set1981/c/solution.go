package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		if len(res) == 0 {
			buf.WriteString("-1\n")
			continue
		}
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

const H = 30

func solve(a []int) []int {
	n := len(a)

	b := make([]int, n)
	copy(b, a)

	process := func(x int, y int, s int, t int, d int) {
		for i := s; i != t; i += d {
			if y == x {
				y = x * 2
			} else {
				y /= 2
			}
			b[i] = y
		}
	}

	for i := 0; i < n; {
		if b[i] == -1 {
			j := i
			for i < n && a[i] < 0 {
				i++
			}
			if j == 0 && i == n {
				process(1, 2, 0, n, 1)
			} else if j == 0 {
				process(1, b[i], i-1, -1, -1)
			} else if i == n {
				process(1, b[j-1], j, i, 1)
			} else {
				// in mid
				x, y := b[j-1], b[i]
				// 比如 1 -1, -1, -1, 4
				// 经过4次把1变成4
				ln := i - j + 1
				fx := get(x)
				fy := get(y)
				found := false
				for k, v := range fx {
					if w, ok := fy[k]; ok && ln >= v+w && (ln-v-w)%2 == 0 {
						// find one solution
						if v > 0 {
							process(k, x, j, j+v, 1)
						}
						if w > 0 {
							process(k, y, i-1, i-1-w, -1)
						}
						if ln > v+w {
							process(k, k, j+v, i-w, 1)
						}
						found = true
						break
					}
				}
				if !found {
					return nil
				}
			}

		} else {
			// a[i] >= 1
			if i > 0 && b[i] != b[i-1]/2 && b[i-1] != b[i]/2 {
				return nil
			}
			b[i] = a[i]
			i++
		}
	}

	return b
}

func get(num int) map[int]int {
	res := make(map[int]int)
	res[num] = 0
	for num > 1 {
		num /= 2
		res[num] = len(res)
	}
	return res
}
