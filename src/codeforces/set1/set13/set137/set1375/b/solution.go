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
		G := make([][]int, n)
		for i := 0; i < n; i++ {
			G[i] = readNNums(reader, m)
		}
		good := solve(n, m, G)
		if good {
			buf.WriteString("YES\n")
			for i := 0; i < n; i++ {
				for j := 0; j < m; j++ {
					if j < m-1 {
						buf.WriteString(fmt.Sprintf("%d ", G[i][j]))
					} else {
						buf.WriteString(fmt.Sprintf("%d\n", G[i][j]))
					}
				}
			}
		} else {
			buf.WriteString("NO\n")
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

func solve(n, m int, G [][]int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var cnt int
			if i > 0 {
				cnt++
			}
			if i < n-1 {
				cnt++
			}
			if j > 0 {
				cnt++
			}
			if j < m-1 {
				cnt++
			}
			if G[i][j] > cnt {
				return false
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var cnt int
			if i > 0 {
				cnt++
			}
			if i < n-1 {
				cnt++
			}
			if j > 0 {
				cnt++
			}
			if j < m-1 {
				cnt++
			}
			G[i][j] = cnt
		}
	}
	return true
}
