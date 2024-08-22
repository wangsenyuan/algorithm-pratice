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
		s := readString(reader)[:n]
		res := solve(s, k)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%s\n", res))
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

func solve(s string, k int) string {
	n := len(s)

	if n%k != 0 {
		return ""
	}

	freq := make([]int, 26)

	check := func(i int) int {
		// 如果在s[i]处放置
		if s[i] == 'z' {
			return -1
		}
		// change to some thing larger at position i
		m := n - i - 1

		var expect int

		for x := 0; x < 26; x++ {
			r := freq[x] % k
			if r != 0 {
				expect += k - r
			}
		}

		for j := int(s[i]-'a') + 1; j < 26; j++ {
			// 如果在这里放置j
			r := freq[j] % k
			tmp := expect
			if r != 0 {
				tmp -= k - r
			}
			r = (freq[j] + 1) % k
			if r != 0 {
				tmp += k - r
			}
			if tmp <= m {
				return j
			}
		}
		return -1
	}

	ans := []int{-1, -1}

	for i := 0; i < n; i++ {
		j := check(i)
		if j > 0 {
			ans = []int{i, j}
		}
		freq[int(s[i]-'a')]++
	}
	ok := true
	for _, v := range freq {
		if v%k != 0 {
			ok = false
			break
		}
	}
	if ok {
		return s
	}
	if ans[0] < 0 {
		return ""
	}
	for i := 0; i < 26; i++ {
		freq[i] = 0
	}
	for i := 0; i < ans[0]; i++ {
		freq[int(s[i]-'a')]++
	}

	buf := make([]byte, n)
	copy(buf, s[:ans[0]])
	buf[ans[0]] = byte(ans[1] + 'a')
	freq[ans[1]]++
	// 额外需要的字符数
	cnt := make([]int, 26)

	cnt[0] = n - ans[0] - 1

	for x := 1; x < 26; x++ {
		cnt[x] = 0
		r := freq[x] % k
		if r != 0 {
			cnt[x] = k - r
			cnt[0] -= k - r
		}
	}

	for x, i := 0, ans[0]+1; x < 26; x++ {
		for cnt[x] > 0 {
			buf[i] = byte(x + 'a')
			cnt[x]--
			i++
		}
	}

	return string(buf)
}
