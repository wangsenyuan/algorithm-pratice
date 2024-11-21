package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res, _ := process(reader)

	if len(res) == 0 {
		fmt.Println("NIE")
		return
	}
	fmt.Println("TAK")
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) ([]int, [][]int) {
	n, m := readTwoNums(reader)
	a := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, m)
	}
	return solve(a), a
}

func solve(a [][]int) []int {
	n := len(a)
	m := len(a[0])

	construct := func(last int) []int {
		res := make([]int, n)
		var sum int
		for i := 0; i < n; i++ {
			res[i] = 1
			if i != last {
				sum ^= a[i][0]

			}
		}
		if last < 0 {
			// 随便选
			return res
		}
		// 除了last，其他的随便选
		for j := 0; j < m; j++ {
			if sum^a[last][j] != 0 {
				res[last] = j + 1
				break
			}
		}

		return res
	}

	for d := 0; d < 10; d++ {
		var cnt2 int

		for i := 0; i < n; i++ {
			var cnt int
			for j := 0; j < m; j++ {
				x := (a[i][j] >> d) & 1
				cnt += x
			}
			if cnt > 0 && cnt < m {
				// 这一位时个关键
				return construct(i)
			}
			if cnt == m {
				cnt2++
			}
		}
		if cnt2%2 == 1 {
			// 有一些是全0的，有一些全1的
			return construct(-1)
		}
	}

	return nil
}
