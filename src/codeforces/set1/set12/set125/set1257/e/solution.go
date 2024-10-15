package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	c := readNNums(reader, k)

	res := solve(a, b, c)

	fmt.Println(res)
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

func solve(a, b, c []int) int {
	n := len(a) + len(b) + len(c)

	// 都放在一个集合中
	res := n - max(len(a), len(b), len(c))

	belong := make([]int, n)
	for _, x := range a {
		belong[x-1] = 0
	}
	for _, x := range b {
		belong[x-1] = 1
	}
	for _, x := range c {
		belong[x-1] = 2
	}

	// dp[i] 表示前i个数属于第一个集合时，需要移动多少个数进去
	var dp int
	p2 := make([]int, n)
	lp := make([]int, n)

	k2 := len(b)

	for i := 0; i < n; i++ {
		if i > 0 {
			p2[i] = p2[i-1]
		}
		if belong[i] != 0 {
			// 要从某个集合移动进来
			dp++
		}
		if belong[i] == 1 {
			p2[i]++
		}
		lp[i] = dp + p2[i] - i
		if i > 0 {
			lp[i] = min(lp[i], lp[i-1])
		}
		// 如果i前面的都是集合1的，i后面的都是集合2的
		// 那么 i + 1 - dp[i] 表示前面的不需要移动的部分 + 后面不需要移动的部分, 其他的都是要移动的
		res = min(res, n-(i+1-dp+k2-p2[i]))
	}
	// fp[r] + dp[l] 最小？
	// 不对，需要把l后面的所有的部分，移出第一个集合
	// dp[l] + x + fp[r] + y
	// x = 把集合1中的移动到集合2中，
	// y = 把集合3中的移动到集合2中
	// = r - l - size(set2) ? 不对，因为有一部分2中的移动了集合1或者3中
	// fp + dp[l] + r - l - 1 - (集合2在区间[l+1, r-1]中的数量, 这部分不用移动)
	// fp + dp[l] + r - l - 1 - (pref[r-1] - pref[l])
	// fp + dp[l] + r - l - 1 - pref[r-1] + pref[l]
	// fp + r - 1 - pref[r-1] + dp[l] + pref[l] - l 最小
	var fp int
	var p3 int
	for i := n - 1; i > 0; i-- {
		if belong[i] != 2 {
			fp++
		}
		if belong[i] == 1 {
			p3++
		}

		res = min(res, fp+i-1-p2[i-1]+lp[i])
		// 后面不需要移动的部分 n - i - fp
		// + 集合b中，处理前面部分， k2 - p3
		res = min(res, n-(n-i-fp+k2-p3))
	}

	// 如何处理只有集合1/和集合2的情况呢？

	return res
}
