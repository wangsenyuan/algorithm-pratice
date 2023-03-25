package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		P := readNNums(reader, n)
		res := solve(P)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
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

func readLenNums(r *bufio.Reader) []int {
	s, _ := r.ReadBytes('\n')
	var n int
	pos := readInt(s, 0, &n) + 1
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		pos = readInt(s, pos, &arr[i]) + 1
	}
	return arr
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const INF = 1e9

func solve(P []int) []int {
	n := len(P)
	dp := make([][]int, n)
	pr := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		pr[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = -INF
	dp[0][1] = -INF

	get := func(i int, j int) int {
		if j == 0 {
			return P[i]
		}
		return -P[i]
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < 2; j++ {
			if dp[i][j] < INF {
				x := get(i, j)
				y := dp[i][j]
				if x < y {
					x, y = y, x
				}
				for nj := 0; nj < 2; nj++ {
					z := get(i+1, nj)
					if z > x {
						if dp[i+1][nj] > y {
							dp[i+1][nj] = y
							pr[i+1][nj] = j
						}
					} else if z > y && dp[i+1][nj] > x {
						dp[i+1][nj] = x
						pr[i+1][nj] = j
					}
				}
			}
		}
	}

	for j := 0; j < 2; j++ {
		if dp[n-1][j] < INF {
			res := make([]int, n)
			for i := n - 1; i >= 0; i-- {
				res[i] = P[i]
				if j == 1 {
					res[i] = -P[i]
				}
				j = pr[i][j]
			}
			return res
		}
	}

	return nil
}

func solve1(P []int) []int {
	n := len(P)
	pr := make([][][]Pair, n)
	for i := 0; i < n; i++ {
		pr[i] = make([][]Pair, 2)
		for j := 0; j < 2; j++ {
			pr[i][j] = make([]Pair, 2)
		}
	}
	dp := make([]int, 4)
	fp := make([]int, 4)
	dp[0] = -INF
	dp[1] = -INF
	for i := 0; i < n-1; i++ {
		for j := 0; j < 4; j++ {
			fp[j] = INF
		}
		for pos := 0; pos < 2; pos++ {
			for sg := 0; sg < 2; sg++ {
				if dp[pos*2+sg] != INF {
					for nsg := 0; nsg < 2; nsg++ {
						x := P[i]
						if sg == 1 {
							x = -P[i]
						}
						y := dp[pos*2+sg]
						if pos == 1 {
							x, y = y, x
						}
						z := P[i+1]
						if nsg == 1 {
							z = -P[i+1]
						}
						if z > x {
							if fp[nsg] > y {
								fp[nsg] = y
								pr[i+1][0][nsg] = Pair{pos, sg}
							}
						} else if z > y {
							if fp[2+nsg] > y {
								fp[2+nsg] = x
								pr[i+1][1][nsg] = Pair{pos, sg}
							}
						}
					}
				}
			}
		}
		copy(dp, fp)
	}

	pos := -1
	sg := -1

	for j := 0; j < 2; j++ {
		for k := 0; k < 2; k++ {
			if dp[j*2+k] != INF {
				pos = j
				sg = k
			}
		}
	}

	if pos < 0 {
		return nil
	}

	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		res[i] = P[i]
		if sg == 1 {
			res[i] *= -1
		}
		x := pr[i][pos][sg]
		pos = x.first
		sg = x.second
	}

	return res
}

type Pair struct {
	first  int
	second int
}
