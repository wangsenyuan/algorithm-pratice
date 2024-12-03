package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		k, res, _ := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", k))
		if k < 0 {
			continue
		}
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func process(reader *bufio.Reader) (k int, ans []int, words []string) {
	n := readNum(reader)
	words = make([]string, n)
	for i := range n {
		words[i] = readString(reader)
	}
	k, ans = solve(words)
	return
}

func solve(words []string) (int, []int) {
	// 0,1, 10, 01
	cnt := make([]int, 4)

	for _, word := range words {
		x := int(word[0] - '0')

		if word[0] == word[len(word)-1] {
			cnt[x]++
		} else {
			y := int(word[len(word)-1] - '0')
			if x == 0 && y == 1 {
				cnt[3]++
			} else {
				cnt[(x<<1)|y]++
			}
		}
	}

	// 0  01    10
	// 只有0和1的时候是有问题的
	if cnt[0] > 0 && cnt[1] > 0 && cnt[2]+cnt[3] == 0 {
		return -1, nil
	}
	// 全部0， 全部1， 0和 01, 1和01， 1和10都是ok的
	sum := cnt[2] + cnt[3]

	p := sum & 1
	// 那么只能是 ababa 这样的结构
	x := (sum + p) / 2

	if cnt[2] >= x {
		return find(words, 2, 1, cnt[2]-x)
	}
	// cnt[1] >= x
	return find(words, 1, 2, cnt[3]-x)
}

func find(words []string, u int, v int, cnt int) (int, []int) {
	if cnt == 0 {
		return 0, nil
	}
	// 将words中的u改变成v， 改变成数量为cnt
	occ := make(map[string]bool)
	for _, w := range words {
		if w[0] == w[len(w)-1] {
			continue
		}
		x := int(w[0] - '0')
		y := int(w[len(w)-1] - '0')
		if (x<<1)|y == v {
			occ[w] = true
		}
	}

	var res []int

	for i, w := range words {
		if w[0] == w[len(w)-1] {
			continue
		}
		x := int(w[0] - '0')
		y := int(w[len(w)-1] - '0')
		if (x<<1)|y == u && !occ[reverse(w)] {
			res = append(res, i+1)
			cnt--
			if cnt == 0 {
				break
			}
		}
	}
	if cnt > 0 {
		return -1, nil
	}

	return len(res), res
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func abs(num int) int {
	return max(num, -num)
}
