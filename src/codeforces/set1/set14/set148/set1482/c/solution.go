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
		n, m := readTwoNums(reader)
		days := make([][]int, m)
		for i := 0; i < m; i++ {
			var k int
			s, _ := reader.ReadBytes('\n')
			pos := readInt(s, 0, &k)
			days[i] = make([]int, k)
			for j := 0; j < k; j++ {
				pos = readInt(s, pos+1, &days[i][j])
			}
		}
		res := solve(n, m, days)

		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for i := 0; i < m; i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
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

func solve(n int, m int, days [][]int) []int {
	// 如果每天都可以至少有两个人，那么肯定有答案
	// 应该优先安排参加最多的人，如果这样的人， 则应该先安排那些参与人数少的天数
	day_frient_cnt := make([]int, m)

	type Friend struct {
		id   int
		days []int
	}

	friends := make([]*Friend, n)
	for i := 0; i < n; i++ {
		friends[i] = &Friend{i, make([]int, 0, 1)}
	}

	for i, cur := range days {
		day_frient_cnt[i] = len(cur)
		for j := 0; j < len(cur); j++ {
			cur[j]--
			friends[cur[j]].days = append(friends[cur[j]].days, i)
		}
	}

	sort.Slice(friends, func(i, j int) bool {
		return len(friends[i].days) > len(friends[j].days)
	})

	ans := make([]int, m)

	tmp := make([]int, m)
	for _, cur := range friends {
		var it int
		for _, i := range cur.days {
			if ans[i] > 0 {
				// already ok
				continue
			}
			tmp[it] = i
			it++
		}

		sort.Slice(tmp[:it], func(i, j int) bool {
			a := tmp[i]
			b := tmp[j]
			return day_frient_cnt[a] < day_frient_cnt[b]
		})
		end := min(it, (m+1)/2)
		for i := 0; i < end; i++ {
			ans[tmp[i]] = cur.id + 1
		}
		for i := end; i < it; i++ {
			day_frient_cnt[tmp[i]]--
		}
	}

	for i := 0; i < m; i++ {
		if ans[i] <= 0 {
			return nil
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
