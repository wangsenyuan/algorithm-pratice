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
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			points[i] = readNNums(reader, 2)
		}
		res := solve(points)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
		if s[i] == '\n' || s[i] == '\r' {
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

func index(x, y int) int {
	if x > 0 && y > 0 {
		return 0
	}
	if x > 0 && y < 0 {
		return 1
	}
	if x < 0 && y < 0 {
		return 2
	}
	return 3
}

func solve(points [][]int) int64 {
	cnt := make([][][]int64, 4)
	for i := 0; i < 4; i++ {
		cnt[i] = make([][]int64, 2)
		for j := 0; j < 2; j++ {
			cnt[i][j] = make([]int64, 2)
		}
	}

	for _, p := range points {
		x, y := p[0], p[1]
		if x != 0 && y != 0 {
			cnt[index(x, y)][abs(x) % 2][abs(y) % 2]++
		}
	}
	var res int64

	for ax := 0; ax < 2; ax++ {
		for bx := 0; bx < 2; bx++ {
			for cx := 0; cx < 2; cx++ {
				for dx := 0; dx < 2; dx++ {
					for ay := 0; ay < 2; ay++{
						for by := 0; by < 2; by++ {
							for cy := 0; cy < 2; cy++ {
								for dy := 0; dy < 2; dy++ {
									if abs(ax * (by - dy) + bx * (cy - ay) + cx * (dy - by) + dx * (ay - cy)) % 2 == 0 {
										res += cnt[0][ax][ay] * cnt[1][bx][by] * cnt[2][cx][cy] * cnt[3][dx][dy]
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}