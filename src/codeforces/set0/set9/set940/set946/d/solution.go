package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	timetable := make([]string, n)
	for i := 0; i < n; i++ {
		timetable[i] = readString(reader)
	}

	res := solve(n, m, k, timetable)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

const oo = 1 << 60

func solve(n int, m int, k int, timetable []string) int {
	dp := make([]int, k+1)

	for i := 0; i < k; i++ {
		dp[i] = oo
	}
	dp[0] = 0

	pd := make([]int, k+1)

	fp := make([]int, k+1)

	getBest := func(day string) []int {
		// 一共这么多
		var sum int
		for i := 0; i < m; i++ {
			if day[i] == '1' {
				sum++
			}
		}
		for x := k; x >= 0; x-- {
			// 前面删除cnt[0]个1，后面删除cnt[1]个
			cnt := []int{0, sum}
			if x >= cnt[1] {
				// 删完了
				fp[x] = 0
				continue
			}
			r := 0

			for i := 0; i < m; i++ {
				if day[i] == '1' {
					// 前面删除cnt[0]个1
					// 后面要删除cnt[1]个1
					// 前面删除的多了，后面就删除的少了
					for r < m && cnt[0]+cnt[1] > x {
						if day[r] == '1' {
							cnt[1]--
						}
						r++
					}
					if cnt[0]+cnt[1] == x {
						fp[x] = min(fp[x], r-i)
					}

					cnt[0]++
				}
			}
		}

		return fp
	}

	for _, cur := range timetable {
		for i := 0; i <= k; i++ {
			pd[i] = oo
			fp[i] = m
		}

		fp = getBest(cur)

		for j := k; j >= 0; j-- {
			for x := j; x >= 0; x-- {
				pd[j] = min(pd[j], dp[j-x]+fp[x])
			}
		}
		copy(dp, pd)
	}

	return dp[k]
}
