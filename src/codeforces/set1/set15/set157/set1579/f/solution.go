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
		n, d := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, d)
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

func solve(a []int, d int) int {
	n := len(a)

	if d == n {
		for i := 0; i < n; i++ {
			if a[i] != 0 {
				return -1
			}
		}
		return 0
	}

	var ans int

	pos := make([]int, 20)
	vis := make([]bool, n)
	for it := 0; it < n; it++ {
		if vis[it] {
			continue
		}
		var arr []int
		tmp := a[it]
		for i := it; !vis[i]; i = (i + d) % n {
			vis[i] = true
			arr = append(arr, a[i])
			tmp &= a[i]
		}

		if tmp != 0 {
			return -1
		}
		//m := len(arr)
		arr = double(arr)
		// pos[i]表示最近的0在哪里
		for i := 0; i < 20; i++ {
			pos[i] = -1
		}

		for i := 0; i < len(arr); i++ {
			for j := 0; j < 20; j++ {
				if (arr[i]>>j)&1 == 0 {
					pos[j] = i
				} else if pos[j] >= 0 {
					ans = max(ans, i-pos[j])
				}
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func double(arr []int) []int {
	res := make([]int, len(arr)*2)
	copy(res, arr)
	copy(res[len(arr):], arr)
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
