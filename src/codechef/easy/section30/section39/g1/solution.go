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
		n, k := readTwoNums(reader)
		s := readString(reader)
		res := solve(s[:n], k)
		if len(res) == 0 {
			buf.WriteString("Johnny wins\n")
		} else {
			buf.WriteString("Mary wins\n")
			buf.WriteString(fmt.Sprintf("%s\n", res))
		}
		buf.WriteByte('\n')
		if tc > 0 {
			readString(reader)
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

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

func solve(s string, k int) string {
	// let dp(x) = 如果字符串是x, 当前用户是否可以赢得比赛
	// dp(x) = 把x分成a,b,c三段，其中len(b)且它是连续的1 <= k
	//       = dp(a) && dp(c) ， 无论另外一方如何选择，剩余的那个当前用户都可以获胜
	//       如果只有其中一个字符串能获胜，另外一个不能获胜， 那么对方就能先选择获胜的字符串,把不能获胜的那方留下来
	//       如果两个都不能获胜, 则当前用户获胜，因为另外一方不能获胜，意味这最后一次消除的人是当前用户，可以逼迫对方开始第二个子串
	// 所以 dp[x] = dp[a] && dp[c] || !dp[a] && !dp[c]
	// 现在有个n**3 * k 的算法
	// dp[i][j] 表示dp(x)
	// 然后迭代u from i to j, 然后迭代v from 1 to k,
	// too slow
	n := len(s)
	table := make([]int, n+1)
	var count []int
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			x := table[j] ^ table[i-(j+1)]
			if x >= len(count) {
				count = addValues(count, x-len(count)+1)
			}
			count[x]++
		}
		for j := 0; j+1 <= i-k; j++ {
			x := table[j] ^ table[i-k-(j+1)]
			if x >= len(count) {
				count = addValues(count, x-len(count)+1)
			}
			count[x]--
		}
		table[i] = 0
		for table[i] < len(count) && count[table[i]] != 0 {
			table[i]++
		}
	}

	var sum int

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			continue
		}
		j := i
		for i < n && s[i] == '1' {
			i++
		}
		sum ^= table[i-j]
		i--
	}

	if sum == 0 {
		return ""
	}
	buf := []byte(s)

	var pos, L int

	for i := 0; i < n; i++ {
		if s[i] == '0' {
			continue
		}
		j := i
		for i < n && s[i] == '1' {
			i++
		}

		sum2 := sum ^ table[i-j]

		for kk := k; kk > L; kk-- {
			for p := j; p+kk <= i; p++ {
				if (sum2 ^ table[p-j] ^ table[i-(p+kk)]) == 0 {
					L = kk
					pos = p
					break
				}
			}
		}

		i--
	}

	for i := pos; i < pos+L; i++ {
		buf[i] = '0'
	}

	return string(buf)
}

func addValues(arr []int, x int) []int {
	res := make([]int, len(arr)+x)
	copy(res, arr)
	return res
}
