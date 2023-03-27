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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(A []int) []int {
	// 假设当前的mode是x
	// a, a, a, b, b, b, c, c, c ....
	// 通过一次操作，可以很容易得到x-1, 即把其中的某个freq = mode-freq 的数字改成其他的数字，即可
	// 要得到 x - 2,一次是不是也足够? 假设把 a => b, 那么如果 等于freq(b)的个数 = x - 1, 即可
	// 否则的话，必须要至少进行两次操作
	// i from x to 1 似乎可以在O(c)时间内完成
	// i from x + 1 to n 呢？
	// 这里有两种操作，一种是从freq < x的部分中，通过改变其中的一部分到另一部分，凑出一个新的列表freq = x
	// 另外一个操作是从 freq = x 中的部分，通过拆分，使得新的 freq = y 的个数等于 i
	// 关键是要找出最少的次数
	//考虑那些不可能的i是 c = n / i (i个freq = c)的部分，r = n % i
	// if r == 0, c >= 1即可，else, r > 0, then c > 1
	// for i from 1 to n
	// if i is reachable
	//  然后找出 freq >= c 中的最多的那个 freq of freq, 剩下的就是需要改变的部分
	n := len(A)
	freq := make(map[int]int)
	for _, num := range A {
		freq[num]++
	}

	cnt := make([]int, n+1)
	for _, v := range freq {
		cnt[v]++
	}

	freqs := []int{0}
	for x := 1; x <= n; x++ {
		if cnt[x] > 0 {
			freqs = append(freqs, x)
		}
	}

	sufSum := make([]int, n+2)
	sufCnt := make([]int, n+2)

	for i := n; i >= 1; i-- {
		sufCnt[i], sufSum[i] = cnt[i], i*cnt[i]
		sufCnt[i] += sufCnt[i+1]
		sufSum[i] += sufSum[i+1]
	}

	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = -1
	}

	for mf := 1; mf <= n; mf++ {
		pt := search(len(freqs), func(i int) bool {
			return freqs[i] >= mf
		})
		pt--
		// freqs[pt] < mf
		rem := cnt[freqs[pt]]
		if pt == 0 {
			rem = n + 100
		}

		var sum int

		mc := 1

		for mc*mf <= n {
			// Everything >= mxfreq should be brought down -> S1 moves
			// If there are x >= ct such things, (x - ct) more moves
			// If there are x < ct such things, take the largest (ct - x) of them and bring them up to x -> S2 moves
			// ans = max(S1, S2)
			moresum, morect := sufSum[mf], sufCnt[mf]
			moves := moresum - mf*morect
			if morect >= mc {
				moves += morect - mc
			} else {
				sum += mf - freqs[pt]
				rem--
				if rem == 0 {
					pt--
					rem = cnt[freqs[pt]]
					if pt == 0 {
						rem = n + 100
					}
				}
				moves = max(moves, sum)
			}

			if 2*mc <= n || mc == n {
				if ans[mc] < 0 {
					ans[mc] = moves
				} else {
					ans[mc] = min(ans[mc], moves)
				}
			}
			mc++
		}
	}

	return ans[1:]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
