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
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(a, b, k)
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

const inf = 1 << 40

type pair struct {
	first  int
	second int
}

func solve(a []int, b []int, k int) int {
	n := len(a)
	arr := make([]pair, n)
	best := -1
	for i := 0; i < n; i++ {
		arr[i] = pair{a[i], b[i]}
		if b[i] == 1 {
			if best < 0 || a[best] < a[i] {
				best = i
			}
		}
	}

	var nums []int
	for i := 0; i < n; i++ {
		if i != best {
			nums = append(nums, a[i])
		}
	}

	sort.Ints(nums)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first < arr[j].first
	})

	if best < 0 {
		return arr[(n-2)/2].first + arr[n-1].first
	}

	var good []int
	var prev []int
	var sum int

	for i := 0; i < n; i++ {
		if arr[i].second == 1 {
			good = append(good, i)
			sum += arr[i].first
			prev = append(prev, sum)
		}
	}

	med := nums[(len(nums)-1)/2]

	res := max(a[best]+k+med, arr[n-1].first+arr[(n-2)/2].first)

	check2 := func(x int, id int) bool {
		p := search(0, n, func(i int) bool {
			return arr[i].first >= x
		})
		cnt := n - p
		if p <= id {
			cnt--
		}
		if cnt >= n-1-(n-2)/2 {
			return true
		}
		cnt = n - 1 - (n-2)/2 - cnt
		// 在p前面，需要cnt能增加到x
		j := search(0, len(good), func(i int) bool {
			return good[i] >= p
		})
		if j < cnt {
			return false
		}
		sum := prev[j-1]
		if j > cnt {
			sum -= prev[j-cnt-1]
		}
		return cnt*x-sum <= k
	}

	check := func(id int) int {
		// b[id] = 0
		v := arr[id].first

		l, r := 0, inf

		for l < r {
			mid := (l + r) / 2
			if !check2(mid, id) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return r - 1 + v
	}

	for i := n - 1; i >= 0; i-- {
		if arr[i].second == 0 {
			res = max(res, check(i))
			break
		}
	}

	return res
}

func search(l, r int, f func(int) bool) int {
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
