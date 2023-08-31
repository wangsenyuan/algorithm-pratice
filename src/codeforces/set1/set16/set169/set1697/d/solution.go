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

	ask := func(arr []int) string {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("? %d", len(arr)))
		for i := 0; i < len(arr); i++ {
			buf.WriteString(fmt.Sprintf(" %d", arr[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		return readString(reader)
	}

	n := readNum(reader)

	res := solve(n, ask)

	fmt.Printf("! %s\n", res)
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

func solve(n int, ask func([]int) string) string {
	// 如果 n <= 26, 可以直接用操作1
	// 考虑操作2，它提供的是在区间 l..r中间有多少不同的字符
	// 这个信息有什么用？
	// 对于每一个位置，可以通过2分查询，找到，左边和它一致的一段

	ask_ch := func(i int) string {
		return ask([]int{i + 1})
	}

	ask_cnt := func(l int, r int) int {
		res := ask([]int{l + 1, r + 1})
		var cnt int
		readInt([]byte(res), 0, &cnt)
		return cnt
	}

	var buf bytes.Buffer

	cnt := make([][]int, n+1)

	for i := 0; i < n; i++ {
		if i == 0 {
			buf.WriteString(ask_ch(i))
			cnt[0] = []int{1}
			continue
		}
		cur := ask_cnt(0, i)
		if cur > cnt[i-1][0] {
			buf.WriteString(ask_ch(i))
		} else {
			last := make(map[byte]int)
			s := buf.Bytes()
			for j := 0; j < len(s); j++ {
				last[s[j]] = j
			}
			var lasts []int
			for _, j := range last {
				lasts = append(lasts, j)
			}
			sort.Ints(lasts)
			l, r := 0, len(lasts)
			for r-l > 1 {
				mid := (l + r) / 2
				tmp := ask_cnt(lasts[mid], i)
				if tmp == cnt[i-1][lasts[mid]] {
					l = mid
				} else {
					r = mid
				}
			}
			buf.WriteByte(s[lasts[l]])
		}

		s := buf.Bytes()
		cnt[i] = make([]int, i+1)
		set := make(map[byte]int)
		for j := i; j >= 0; j-- {
			set[s[j]]++
			cnt[i][j] = len(set)
		}
	}

	return buf.String()
}
