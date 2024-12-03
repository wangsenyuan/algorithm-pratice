package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	l := readString(reader)
	r := readString(reader)
	res := solve(l, r)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
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

func solve(l string, r string) [][]int {
	count := func(s string) [][]int {
		cnt := make([][]int, 27)
		for i, x := range []byte(s) {
			if x == '?' {
				cnt[26] = append(cnt[26], i)
			} else {
				cnt[int(x-'a')] = append(cnt[int(x-'a')], i)
			}
		}
		return cnt
	}

	cnt1 := count(l)
	cnt2 := count(r)
	var res [][]int

	for i := 0; i < 26; i++ {
		var j int
		for ; j < min(len(cnt1[i]), len(cnt2[i])); j++ {
			a, b := cnt1[i][j], cnt2[i][j]
			res = append(res, []int{a + 1, b + 1})
		}
		cnt1[i] = cnt1[i][j:]
		cnt2[i] = cnt2[i][j:]
	}

	for i := 0; i < 26; i++ {
		var j int
		for ; j < min(len(cnt1[i]), len(cnt2[26])); j++ {
			a, b := cnt1[i][j], cnt2[26][j]
			res = append(res, []int{a + 1, b + 1})
		}
		cnt2[26] = cnt2[26][j:]
	}

	for i := 0; i < 26; i++ {
		var j int
		for ; j < min(len(cnt2[i]), len(cnt1[26])); j++ {
			a, b := cnt2[i][j], cnt1[26][j]
			res = append(res, []int{b + 1, a + 1})
		}
		cnt1[26] = cnt1[26][j:]
	}

	for i := 0; i < len(cnt1[26]) && i < len(cnt2[26]); i++ {
		x, y := cnt1[26][i], cnt2[26][i]
		res = append(res, []int{x + 1, y + 1})
	}

	return res
}
