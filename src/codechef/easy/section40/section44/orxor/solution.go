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
		n, x := readTwoNums(reader)
		res := solve(n, x)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for _, ans := range res {
				for i := 0; i < len(ans); i++ {
					buf.WriteString(fmt.Sprintf("%d ", ans[i]))
				}
				buf.WriteByte('\n')
			}
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, x int) [][]int {
	if n == 2 {
		if x == 3 {
			return [][]int{{1, 1, 2}}
		}
		return nil
	}
	msb := make([]int, 2)
	pw := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		pw = append(pw, 1<<i)
		if (n>>i)&1 == 1 {
			msb[0] = i
		}
		if (x>>i)&1 == 1 {
			msb[1] = i
		}
	}

	if msb[0] < msb[1] {
		return nil
	}
	if (n&(n-1)) == 0 && x&n == 0 {
		return nil
	}
	var val int
	var res [][]int
	for i, j := 1, 0; i <= n; i++ {
		for j < len(pw) && pw[j] < i {
			j++
		}
		if j == len(pw) || pw[j] != i {
			if i > 3 {
				res = append(res, []int{1, i, val})
			}
			val |= i
		}
	}

	for j := 0; j < len(pw) && pw[j] <= n; j++ {
		if (x & pw[j]) == pw[j] {
			res = append(res, []int{1, pw[j], val})
			val |= pw[j]
		} else {
			res = append(res, []int{2, pw[j], val})
			val ^= pw[j]
		}
	}

	return res
}
