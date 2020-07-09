package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		removed := make([]int64, n)
		for i := 0; i < n; i++ {
			s, _ := reader.ReadBytes('\n')
			var res int64
			for j := 0; j < m; j++ {
				res = res<<1 | int64(s[j]-'0')
			}
			removed[i] = res
		}
		fmt.Println(solve(m, n, removed))
	}
}

func solve(m int, n int, removed []int64) string {
	var need int64 = int64((1<<m)-n-1)/2 + 1

	var cur int64 = int64(1<<(m-1)) - 1

	for {
		left := cur + 1

		var flag bool

		for i := 0; i < n; i++ {
			flag = flag || removed[i] == cur
			if removed[i] <= cur {
				left--
			}
		}

		if left == need && !flag {
			buf := make([]byte, m)

			for i := 0; i < m; i++ {
				buf[m-1-i] = byte(cur&1 + '0')
				cur >>= 1
			}

			return string(buf)
		}

		if left < need {
			cur++
		} else {
			cur--
		}
	}
}
