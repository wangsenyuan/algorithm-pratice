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
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	dp[n] = make([]int, 4)
	suf := make([][]int, n+1)
	suf[n] = make([]int, 4)
	cnt := make([]int, 4)

	for i := n - 1; i >= 0; i-- {
		dp[i] = make([]int, 4)
		copy(dp[i], dp[i+1])
		if s[i] == 'W' {
			cnt[0]++
			cnt[1] = max(0, cnt[1]-1)
		} else if s[i] == 'S' {
			cnt[0] = max(0, cnt[0]-1)
			cnt[1]++
		} else if s[i] == 'A' {
			cnt[2]++
			cnt[3] = max(0, cnt[3]-1)
		} else {
			cnt[2] = max(0, cnt[2]-1)
			cnt[3]++
		}

		suf[i] = make([]int, 4)
		copy(suf[i], cnt)
		for j := 0; j < 4; j++ {
			dp[i][j] = max(dp[i][j], cnt[j])
		}
	}

	get := func(dp []int) int {
		h := max(dp[0], dp[1]) + 1
		w := max(dp[2], dp[3]) + 1
		return h * w
	}

	res := get(dp[0])

	// 如果在头部插入的话，会怎么样呢？
	rect := make([]int, 4)
	update := func(fp []int, dp []int, sf []int, x byte) {
		// 如果插入字符x
		clear(rect)
		for i := 0; i < 4; i++ {
			rect[i] = cnt[i] + sf[i]
		}
		if x == 'W' {
			rect[0]++
			rect[1] = max(0, cnt[1]+sf[1]-1)
		} else if x == 'S' {
			rect[0] = max(0, cnt[0]+sf[0]-1)
			rect[1]++
		} else if x == 'A' {
			rect[2]++
			rect[3] = max(0, cnt[3]+sf[3]-1)
		} else {
			rect[2] = max(0, cnt[2]+sf[2]-1)
			rect[3]++
		}

		for i := 0; i < 4; i++ {
			rect[i] = max(rect[i], fp[i], dp[i])
		}

		tmp := get(rect)

		res = min(res, tmp)
	}
	const cmd = "WSAD"

	update2 := func(fp []int, dp []int, sf []int) {
		for _, x := range []byte(cmd) {
			update(fp, dp, sf, x)
		}
	}

	clear(cnt)
	fp := make([]int, 4)
	update2(fp, dp[0], suf[0])

	for i := 0; i < n; i++ {
		if s[i] == 'W' {
			cnt[0]++
			cnt[1] = max(0, cnt[1]-1)
		} else if s[i] == 'S' {
			cnt[0] = max(0, cnt[0]-1)
			cnt[1]++
		} else if s[i] == 'A' {
			cnt[2]++
			cnt[3] = max(0, cnt[3]-1)
		} else {
			cnt[2] = max(0, cnt[2]-1)
			cnt[3]++
		}
		for j := 0; j < 4; j++ {
			fp[j] = max(fp[j], cnt[j])
		}
		update2(fp, dp[i+1], suf[i+1])
	}

	return res
}
