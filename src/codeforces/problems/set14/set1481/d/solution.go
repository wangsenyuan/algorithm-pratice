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
		n, m := readTwoNums(reader)
		G := make([][]byte, n)
		for i := 0; i < n; i++ {
			G[i], _ = reader.ReadBytes('\n')
		}
		res := solve(n, m, G)

		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
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

func solve(n int, m int, G [][]byte) []int {
	res := make([]int, m+1)
	if m&1 == 1 {
		for i := 0; i <= m; i++ {
			res[i] = (i & 1) + 1
		}
		//res[m] = 1
		return res
	}
	// then m is even
	var ok bool
	var pivot [2]int

	has := make([][]int, n)
	for i := 0; i < n; i++ {
		has[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			has[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if G[i][j] == G[j][i] {
				pivot = [...]int{i, j}
				ok = true
			}
			has[i][int(G[i][j]-'a')] = j
		}
	}

	if ok {
		for i := 0; i <= m; i++ {
			res[i] = pivot[i&1] + 1
		}
		return res
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			cur := has[j][int(G[i][j]-'a')]
			if cur < 0 {
				continue
			}
			if (m/2)&1 == 1 {
				for k := 0; k <= m; k++ {
					if k%4 == 0 {
						res[k] = i + 1
					} else if k%4 == 2 {
						res[k] = cur + 1
					} else {
						res[k] = j + 1
					}
				}
			} else {
				var p int
				res[p] = j + 1
				p++
				for k := 0; k < m/2; k++ {
					if k&1 == 1 {
						res[p] = j + 1
					} else {
						res[p] = cur + 1
					}
					p++
				}
				for k := 0; k < m/2; k++ {
					if k&1 == 1 {
						res[p] = j + 1
					} else {
						res[p] = i + 1
					}
					p++
				}
				return res
			}

			return res
		}
	}

	return nil
}
