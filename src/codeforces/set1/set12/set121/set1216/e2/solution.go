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
		n := readNum(reader)
		res := solve(n)
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

func solve(k int) int {

	calc := func(l int, r int) int {
		return (l + r) * (r - l + 1) / 2
	}

	count := func(num int) int {
		var sum int

		cur := 1
		d := 1
		for cur <= num {
			next := cur * 10

			var tmp int

			if next <= num {
				// num - cur + 1, num - cur, num - (next - 2)
				// num, num - 1, .... num - 8
				// 1, 2, ....9
				tmp = calc(num-(next-2), num-(cur-1))
			} else {
				tmp = calc(1, num-(cur-1))
			}
			if tmp*d < 0 || tmp*d >= k-sum {
				return k
			}
			sum += tmp * d
			d++
			cur = next
		}

		return sum
	}

	check := func(num int) bool {
		return count(num) >= k
	}

	num := sort.Search(k, check)
	k -= count(num - 1)
	// 1234

	count2 := func(num int) int {
		var sum int
		cur := 1
		d := 1
		for cur <= num {
			next := cur * 10
			if next <= num {
				sum += d * (next - cur)
			} else {
				sum += d * (num - cur + 1)
			}
			cur = next
			d++
		}
		return sum
	}

	num2 := sort.Search(k, func(num int) bool {
		return count2(num) >= k
	})

	k -= count2(num2 - 1)

	s := fmt.Sprintf("%d", num2)

	return int(s[k-1] - '0')
}
