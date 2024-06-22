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
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, k)
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

const D = 30

func solve(a []int, x int) int {
	n := len(a)
	// arr[i].second是每个分组的开始位置
	// arr[i].first 是到d为止每组的xor值
	arr := make([]int, n)
	copy(arr, a)

	var best int = -1

	for d := D - 1; d >= 0; d-- {
		expect := (x >> d) & 1
		var or int
		var cnt int

		for i := 0; i < len(arr); i++ {
			or |= (arr[i] >> d)
			if (arr[i]>>d)&1 == 1 {
				cnt++
			}
		}
		if expect == 1 {
			// 查看当前的分组情况是否能满足条件
			if or < x>>d {
				// 后续的结果不可能比len(arr)更好
				return max(best, len(arr))
			}
			// or == x >> d
			// 看看能不能减少一个分组来获取到比x小的情况
			if cnt%2 == 0 {
				// 偶数个分组为1，把它们合并起来，就可以得到比x小的情况
				tmp := len(arr)
				prev := -1
				for i := 0; i < len(arr); i++ {
					if (arr[i]>>d)&1 == 1 {
						if prev >= 0 {
							tmp -= i - prev
							prev = -1
						} else {
							prev = i
						}
					}
				}
				best = max(best, tmp)
			}
		} else {
			// 必须把1合并掉
			if cnt%2 == 1 {
				// no way
				return max(-1, best)
			}
			var next []int
			var xor int
			for i := 0; i < len(arr); i++ {
				xor ^= arr[i]
				if (xor>>d)&1 == 0 {
					next = append(next, xor)
					xor = 0
				}
			}
			arr = next
		}
	}

	return max(len(arr), best)
}
