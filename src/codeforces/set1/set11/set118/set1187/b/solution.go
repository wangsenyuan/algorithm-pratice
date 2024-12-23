package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readNum(reader)
	s := readString(reader)
	m := readNum(reader)
	words := make([]string, m)
	for i := 0; i < m; i++ {
		words[i] = readString(reader)
	}

	res := solve(s, words)

	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func solve(s string, words []string) []int {
	n := len(s)
	pref := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		pref[i] = make([]int, 26)
	}

	for i := 0; i < n; i++ {
		copy(pref[i+1], pref[i])
		pref[i+1][int(s[i]-'a')]++
	}

	check := func(i int, cnt []int) bool {
		for x := 0; x < 26; x++ {
			if pref[i][x] < cnt[x] {
				return false
			}
		}
		return true
	}

	ans := make([]int, len(words))
	for i, w := range words {
		cnt := getFreq(w)
		ans[i] = sort.Search(n+1, func(j int) bool {
			return check(j, cnt)
		})
	}

	return ans
}

func getFreq(w string) []int {
	freq := make([]int, 26)
	for _, x := range []byte(w) {
		freq[int(x-'a')]++
	}
	return freq
}
