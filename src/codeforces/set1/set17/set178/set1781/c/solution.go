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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		readNum(reader)
		s := readString(reader)
		diff, res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n%s\n", diff, res))
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

func solve(s string) (int, string) {
	n := len(s)

	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		cnt[int(s[i]-'a')]++
	}

	ps := make([]Pair, 26)
	for i := 0; i < 26; i++ {
		ps[i] = Pair{cnt[i], i}
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first > ps[j].first
	})

	// s % x == 0
	// 全部都换
	diff := n
	ans := 1
	for x := 1; x <= 26 && x <= n; x++ {
		if n%x == 0 {
			// 一共有x个字符，每个字符有 n / x 个
			y := n / x
			sum := []int{0, 0}
			for i := 0; i < x; i++ {
				// 多余y的肯定是需要变成其他的
				// 不在这x中的，肯定都是要变化的
				if ps[i].first >= y {
					sum[0] += ps[i].first - y
				}
				sum[1] += ps[i].first
			}
			tmp := sum[0] + n - sum[1]
			if diff > tmp {
				diff = tmp
				ans = x
			}
		}
	}
	rem := make([]bool, 26)
	for i := ans; i < 26; i++ {
		rem[ps[i].second] = true
	}

	buf := []byte(s)
	m := n / ans
	// 一共有ans个字符， 每个字符 n / ans 个
	// 先找出那些必须改变的下标
	change := make([]int, 0, diff)
	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		if cnt[x] > m || rem[x] {
			change = append(change, i)
			cnt[x]--
		}
	}

	for i, j := 0, 0; i < ans; i++ {
		if ps[i].first >= m {
			// already processed
			continue
		}
		// ps[i].first < m
		k := ps[i].first
		for j < len(change) && k < m {
			buf[change[j]] = byte(ps[i].second + 'a')
			j++
			k++
		}
	}

	return diff, string(buf)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Pair struct {
	first  int
	second int
}
