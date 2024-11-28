package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	d, s := readTwoNums(reader)

	res := solve(d, s)

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

func solve(D int, S int) string {

	hd := bits.Len(uint(D))
	hs := bits.Len(uint(S))
	mask1 := int32(1<<hs - 1)

	// rem < d < 500 使用最高8位
	// ds < 5000 < 1e20 使用中间19位
	// num < 10, 使用最低的4位
	encode := func(rem int32, ds int32) int32 {
		return rem<<hs | (ds & mask1)
	}

	decode := func(state int32) (rem int32, ds int32) {
		rem = (state >> hs)
		ds = state & mask1
		return
	}

	s := int32(S)
	d := int32(D)

	from := make([]int32, 1<<(hs+hd))
	for i := range len(from) {
		from[i] = -2
	}

	var que []int32

	for i := int32(1); i <= 9; i++ {
		if i <= s {
			x := encode(i%d, i)
			que = append(que, x)
			from[x] = -1
		}
	}

	update := func(pos int32) string {
		var buf []byte

		for pos >= 0 {
			x := que[pos]
			_, num := decode(x)
			if from[x] >= 0 {
				y := que[from[x]]
				_, ds2 := decode(y)
				num -= ds2
			}
			buf = append(buf, byte(num+'0'))
			pos = from[x]
		}
		reverse(buf)
		return string(buf)
	}

	var tail int

	for tail < len(que) {
		pos := len(que)

		for tail < pos {
			cur := que[tail]
			tail++

			rem, ds := decode(cur)

			if rem == 0 && ds == s {
				return update(int32(tail - 1))
			}

			for i := int32(0); i < 10; i++ {
				nextRem := (rem*10 + i) % d
				nextDs := ds + i
				if nextDs <= s {
					nextState := encode(nextRem, nextDs)
					if from[nextState] == -2 {
						from[nextState] = int32(tail - 1)
						que = append(que, nextState)
					}
				}
			}
		}
	}

	return "-1"
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
