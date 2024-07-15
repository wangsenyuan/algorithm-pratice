package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
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
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func solve(s string) string {
	n := len(s)

	// 需要有这么一种判断，
	// 就是假设当前，我不选s[i]， 还能在后面找到更好的，否则就必须选择s[i]
	// 后面能不能找到更好的，就必须是某个比s[i]更大的字符（这个字符在前面还没有选中保留）
	// 且这个字符后面还有足够的字符
	// 假设这个字符是x，那么这个字符就是x离i最近的位置，然后在这个位置后面假设有w个distinct的，
	// 然后去掉前面的后，还剩余的，要 >= 需要的字符数即可
	fp := make([]int, n+1)
	var flag int

	pos := make([][]int, 26)
	whr := make([]int, 26)
	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - 'a')
		flag |= 1 << x
		fp[i] = flag
		pos[x] = append(pos[x], i)
		whr[x] = len(pos[x])
	}

	check := func(i int, prev int, expect int) bool {
		// 这里是不是一定要使用s[i]
		for x := int(s[i]-'a') + 1; x < 26; x++ {
			// 只要比s[i]大就可以
			// 要找x最近的位置
			for whr[x] > 0 && pos[x][whr[x]-1] < i {
				whr[x]--
			}
			// x 是领头的，所以它不能出现在前面
			if whr[x] == 0 || prev&(1<<x) > 0 {
				continue
			}
			j := pos[x][whr[x]-1]

			tmp := prev & fp[j]
			// 前面有的，后面也有的
			tmp = fp[j] ^ tmp
			// 只有后面有的
			w := bits.OnesCount32(uint32(tmp))
			if w >= expect {
				// 至少可以得到一个x的更好的结果，这里不需要放s[i]
				return false
			}
		}

		return true
	}

	k := bits.OnesCount32(uint32(flag))

	buf := make([]byte, k)

	flag = 0

	for i, j := 0, 0; i < n && j < k; i++ {
		x := int(s[i] - 'a')
		if flag&(1<<x) == 0 && check(i, flag, k-j) {
			buf[j] = s[i]
			j++
			flag |= 1 << x
		}
	}

	return string(buf)
}
