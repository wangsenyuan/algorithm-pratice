package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		readString(reader)
		s := readString(reader)
		ok, res := solve(s)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d\n", x))
			}
		}
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

var base = "LIT"

type pair struct {
	first  int
	second int
}

func cmp(a, b pair) int {
	return a.first - b.first
}

func solve(s string) (ok bool, ans []int) {
	n := len(s)
	if count(s, s[0]) == n {
		return false, nil
	}

	op := func(i int) {
		c := base[0]
		if s[i] == c || s[i+1] == c {
			c = base[1]
		}
		if s[i] == c || s[i+1] == c {
			c = base[2]
		}
		ans = append(ans, i)
		s = s[:i+1] + string(c) + s[i+1:]
	}

	for {
		var cnt []pair
		for i := range 3 {
			cnt = append(cnt, pair{count(s, base[i]), i})
		}

		slices.SortFunc(cnt, cmp)

		if cnt[0].first == cnt[2].first {
			break
		}
		done := false
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				continue
			}
			if s[i] != base[cnt[0].second] && s[i+1] != base[cnt[0].second] {
				op(i)
				done = true
				break
			}
		}
		if done {
			continue
		}
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				continue
			}
			if s[i] == base[cnt[2].second] {
				op(i)
				op(i + 1)
				op(i + 2)
				op(i)
				break
			}
			if s[i+1] == base[cnt[2].second] {
				op(i)
				op(i)
				op(i + 1)
				op(i + 3)
				break
			}
		}
	}
	for i := range ans {
		ans[i]++
	}

	return true, ans
}

func count(s string, c byte) int {
	var res int
	for _, x := range []byte(s) {
		if x == c {
			res++
		}
	}
	return res
}
