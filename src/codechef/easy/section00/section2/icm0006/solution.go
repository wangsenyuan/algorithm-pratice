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
		bb, _ := reader.ReadBytes('\n')
		var n int
		pos := readInt(bb, 0, &n) + 1
		var k uint64
		readUint64(bb, pos, &k)
		S, _ := reader.ReadString('\n')
		S = S[:len(S)-1]
		res := solve(S, n, int64(k))
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(S string, n int, k int64) int {
	m := len(S)
	if n == 1 {
		// k has to be 0
		return 0
	}
	// 将n个S连接在一起，在结果字符串中， 获得至少k个等于A的字串
	// 假设S有i段相等的字串组成比如S = abcabc, i = 2
	// 那么S重复n次后，正好可以有 n + (n - 1) * (i - 1) 个S = (n - 1) * i + 1
	// 其中第一个n表示,下班0, m, 2 * m, ... (n - 1) * m 的子串
	// 从下一段开始的字串有（n-1）次重复，且只有（i-1）个
	L := (k - 1) / int64(n-1)
	if L*int64(n-1)+1 < k {
		L++
	}
	cnt := make([][]int, m)
	for i := 0; i < m; i++ {
		cnt[i] = make([]int, 26)
	}
	find := func(l int) int {
		// each segments is length l
		for i := 0; i < m; i++ {
			x := int(S[i] - 'a')
			cnt[i%l][x]++
		}
		var res int
		tot := m / l
		for i := 0; i < l; i++ {
			tmp := tot
			for j := 0; j < 26; j++ {
				tmp = min(tmp, tot-cnt[i][j])
			}
			res += tmp
		}
		for i := 0; i < m; i++ {
			x := int(S[i] - 'a')
			cnt[i%l][x]--
		}
		return res
	}

	l := int(L)
	// l <= n
	best := m
	for x := 1; x <= m/x; x++ {
		if m%x == 0 {
			if x >= l {
				best = min(best, find(m/x))
			}
			y := m / x
			if y >= l && y != x {
				best = min(best, find(m/y))
			}
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
