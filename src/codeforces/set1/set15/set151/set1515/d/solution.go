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
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, l, r := readThreeNums(reader)
		a := readNNums(reader, n)
		res := solve(n, l, r, a)
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

func solve(n int, l int, r int, c []int) int {
	// first l for left side
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	freq := make([]int, n)
	for i := 0; i < n; i++ {
		x := c[i] - 1
		if i < l {
			freq[x]++
		} else {
			freq[x]--
		}
	}

	expect := abs(l-r) / 2
	// expect * 2 == l - r
	var res int

	if l > r {
		// 需要把一部分l变成r
		sort.Slice(id, func(i, j int) bool {
			return freq[id[i]] > freq[id[j]]
		})
		for i := 0; i < n && expect > 0 && freq[id[i]] > 0; i++ {
			// 这些是左边比右边多的颜色
			x := freq[id[i]]
			x = min(x/2, expect)
			expect -= x
			res += x
			// 这边减去了，另外一边加上了，所以是2 * x
			freq[id[i]] -= 2 * x
		}
		if expect > 0 {
			sort.Slice(id, func(i, j int) bool {
				return freq[id[i]] > freq[id[j]]
			})
			// 相同颜色的没法搞定，只能处理不同颜色的, 把颜色多的改成掉
			for i := 0; i < n && expect > 0; i++ {
				x := min(freq[id[i]], expect)
				// x是左边比右边多的那部分， 全部转换,颜色没有变化
				expect -= x
				res += x
				freq[id[i]] -= 2 * x
			}
		}
	} else if l < r {
		sort.Slice(id, func(i, j int) bool {
			return freq[id[i]] < freq[id[j]]
		})

		for i := 0; i < n && expect > 0 && freq[id[i]] < 0; i++ {
			x := abs(freq[id[i]])
			x = min(x/2, expect)
			expect -= x
			res += x
			freq[id[i]] += 2 * x
		}
		if expect > 0 {
			sort.Slice(id, func(i, j int) bool {
				return freq[id[i]] < freq[id[j]]
			})
			for i := 0; i < n && expect > 0; i++ {
				x := min(abs(freq[id[i]]), expect)
				expect -= x
				res += x
				freq[id[i]] += 2 * x
			}
		}
	}
	var cost int
	// expect = 0, 现在把颜色搞平
	for i := 0; i < n; i++ {
		if freq[i] != 0 {
			cost += abs(freq[i])
		}
	}

	res += cost / 2
	return res
}

func abs(num int) int {
	return max(num, -num)
}
