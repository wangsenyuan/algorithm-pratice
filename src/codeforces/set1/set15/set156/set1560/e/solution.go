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
		x := readString(reader)
		ok, s, o := solve(x)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%s %s\n", s, o))
		}
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(x string) (bool, string, string) {
	freq := make([]int, 26)
	var disc int
	n := len(x)
	for i := 0; i < n; i++ {
		a := int(x[i] - 'a')
		freq[a]++
		if freq[a] == 1 {
			disc++
		}
	}
	// 要操作disc后，全部删除掉
	rem := make([]byte, disc)
	var flag int
	// flag[i] 表示i已经被删除了
	var cnt int
	pos := n - 1
	r := n - 1
	for pos >= 0 && disc > 0 {
		j := pos
		a := -1
		for j >= 0 {
			a = int(x[j] - 'a')
			if (flag>>a)&1 == 0 {
				break
			}
			// 当前删除的肯定是a
			j--
		}
		if a < 0 {
			return false, "", ""
		}

		if freq[a]%disc != 0 {
			return false, "", ""
		}

		rem[disc-1] = x[j]

		cnt += freq[a] / disc
		if pos+1-cnt < 0 {
			return false, "", ""
		}
		// 删除这么多
		flag |= 1 << a
		for i := pos; i > pos-cnt; i-- {
			b := int(x[i] - 'a')
			if (flag>>b)&1 == 0 {
				// 出现了未删除的字符
				return false, "", ""
			}
			if a != b {
				if b != int(x[r]-'a') {
					return false, "", ""
				}
				r--
			}
		}
		r = pos

		pos -= cnt
		disc--
	}

	return true, x[:r+1], string(rem)
}
