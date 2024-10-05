package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	pos := readNNums(reader, 6)

	res := solve(pos)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func area(w int, h int) int {
	return w * h
}

func solve(dims []int) []string {
	var tot int
	for i := 0; i < 3; i++ {
		tot += area(dims[2*i], dims[2*i+1])
	}

	side := int(math.Sqrt(float64(tot)))

	if side*side != tot {
		return nil
	}
	// 边长确定了
	// 只有3种可能性，stacked， 一个在顶部，另外两个竖直放置
	for i := 0; i < 3; i++ {
		if dims[2*i] > dims[2*i+1] {
			dims[2*i], dims[2*i+1] = dims[2*i+1], dims[2*i]
		}
	}

	res := make([]string, side)

	if dims[1] == dims[3] && dims[3] == dims[5] && dims[1] == side && dims[0]+dims[2]+dims[4] == side {
		// stacked
		for i := 0; i < dims[0]; i++ {
			res[i] = strings.Repeat("A", side)
		}
		for i := dims[0]; i < dims[0]+dims[2]; i++ {
			res[i] = strings.Repeat("B", side)
		}
		for i := dims[0] + dims[2]; i < side; i++ {
			res[i] = strings.Repeat("C", side)
		}
		return res
	}

	construct := func(a, ah, b, bw, c int) []string {
		res := make([][]byte, side)
		for i := 0; i < side; i++ {
			res[i] = make([]byte, side)
			for j := 0; j < side; j++ {
				res[i][j] = byte('A' + c)
			}
		}

		for i := 0; i < ah; i++ {
			for j := 0; j < side; j++ {
				res[i][j] = byte('A' + a)
			}
		}

		for i := ah; i < side; i++ {
			for j := 0; j < bw; j++ {
				res[i][j] = byte('A' + b)
			}
		}

		ans := make([]string, side)
		for i := 0; i < side; i++ {
			ans[i] = string(res[i])
		}

		return ans
	}

	for i := 0; i < 3; i++ {
		// 这个是否可以放在顶部
		if dims[2*i+1] == side {
			height := side - dims[2*i]
			for j := 0; j < 3; j++ {
				if i == j {
					continue
				}
				k := 3 - i - j

				for jx := 0; jx < 2; jx++ {
					for kx := 0; kx < 2; kx++ {
						if dims[2*j+jx]+dims[2*k+kx] == side && dims[2*j+1-jx] == height && dims[2*k+1-kx] == height {
							return construct(i, dims[2*i], j, dims[2*j+jx], k)
						}
					}
				}
			}
		}
	}

	return nil
}
