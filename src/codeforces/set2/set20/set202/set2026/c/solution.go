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
	for range tc {
		res := process(reader)
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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	s := readString(reader)
	return solve(n, s)
}

func solve(n int, s string) int {
	// 还挺难实现的
	// 如果在第i天到达，如果此时有超过2件，那么就应该购买最便宜的那个，并节省最贵的那个
	// 第一天是0/1没有区别，到达与否，它都要使用1元
	// 总的花费是 (1 + n) * n / 2
	// 那就看节省多少。节省的正好就是到达的那些天
	if n == 1 {
		return 1
	}

	// que表示要去参加的天数
	que := make([]int, n)
	var head, tail int
	// cnt 表示s[i] = 0 的天数
	var cnt int

	for i := 1; i <= n; i++ {
		if s[i-1] == '0' {
			cnt++
		} else {
			if cnt > 0 {
				que[head] = i
				head++
				cnt--
			} else {
				// cnt == 0
				if tail < head {
					// 交换最前面的
					tail++
					que[head] = i
					head++
				}
				// 不管怎么样cnt肯定加1
				cnt++
			}
		}
	}
	res := (n + 1) * n / 2

	for tail < head {
		res -= que[tail]
		tail++
	}

	return res
}
