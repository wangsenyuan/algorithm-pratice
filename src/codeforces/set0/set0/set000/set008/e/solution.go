package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadBytes('\n')
	var n, k uint64
	pos := readUint64(s, 0, &n)
	readUint64(s, pos+1, &k)

	fmt.Println(solve(int(n), int64(k)))
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

const N = 50

var s [N + 1]int

var dp [N + 1][2][2]int64

func solve(n int, k int64) string {
	nPrefix := n>>1 + n&1

	for i := 0; i <= n; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				dp[i][j][k] = -1
			}
		}
		s[i] = 0
	}

	rem := k + 1
	var found bool

	for prefix := 0; prefix < 1<<nPrefix; prefix++ {
		if prefix > 0 {
			i := nPrefix - 1
			for i >= 0 {
				if s[i] == 0 {
					s[i] = 1
					break
				}
				s[i] = 0
				i--
			}

			for j := n - nPrefix; j <= n-1-i; j++ {
				dp[j][0][0] = -1
				dp[j][0][1] = -1
				dp[j][1][0] = -1
				dp[j][1][1] = -1
			}
		}

		//firstPos := nPrefix
		less, lessRev := 1, 1
		if n&1 == 1 && s[nPrefix-1] == 1 {
			lessRev = 0
		}
		tmp := f(n, nPrefix, less, lessRev)

		if tmp < rem {
			rem -= tmp
		} else {
			found = true
			break
		}
	}

	if !found {
		return "-1"
	}

	less, lessRev := 1, 1
	if n&1 == 1 && s[nPrefix-1] == 1 {
		lessRev = 0
	}

	for pos := nPrefix; pos < n; pos++ {
		nextLess, nextLessRev := less, lessRev
		if 0 < s[n-1-pos] {
			nextLess = 0
		}
		if 1 > s[n-1-pos] {
			nextLessRev = 1
		}

		tmp := f(n, pos+1, nextLess, nextLessRev)

		if tmp >= rem {
			s[pos] = 0
			less, lessRev = nextLess, nextLessRev
			continue
		}

		rem -= tmp
		s[pos] = 1
		if 1 > s[n-1-pos] {
			less = 1
		}
		if 0 < s[n-1-pos] {
			lessRev = 0
		}
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if s[i] == 0 {
			buf.WriteByte('0')
		} else {
			buf.WriteByte('1')
		}
	}
	return buf.String()
}

func f(n int, pos int, less int, lessRev int) int64 {
	if dp[pos][less][lessRev] != -1 {
		return dp[pos][less][lessRev]
	}
	if pos >= n {
		dp[pos][less][lessRev] = int64(less & lessRev)

		return dp[pos][less][lessRev]
	}

	var ret int64

	for i := 0; i <= 1; i++ {
		nextLess := less
		nextLessRev := lessRev
		if i > s[n-1-pos] {
			nextLess = 1
		} else if i < s[n-1-pos] {
			nextLess = 0
		}
		if 1-i > s[n-1-pos] {
			nextLessRev = 1
		} else if 1-i < s[n-1-pos] {
			nextLessRev = 0
		}
		ret += f(n, pos+1, nextLess, nextLessRev)
	}

	dp[pos][less][lessRev] = ret
	return ret
}
