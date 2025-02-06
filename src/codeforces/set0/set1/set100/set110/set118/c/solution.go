package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	cost, res := process(reader)

	fmt.Println(cost)
	fmt.Println(res)
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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) (cost int, res string) {
	n, k := readTwoNums(reader)
	s := readString(reader)[:n]
	cost, res = solve(s, k)
	return
}

const inf = 1 << 60

func solve(s string, k int) (cost int, res string) {
	// n := len(s)

	arr := make([][]int, 10)
	for i, x := range []byte(s) {
		y := int(x - '0')
		arr[y] = append(arr[y], i)
	}

	cost = inf

	var best []int

	for x := 0; x < 10; x++ {
		if len(arr[x]) >= k {
			return 0, s
		}
		cnt := len(arr[x])
		d := 1
		var tmp int
		for cnt < k {
			if x-d >= 0 {
				w := min(k-cnt, len(arr[x-d]))
				cnt += w
				tmp += w * d
			}

			if x+d < 10 {
				w := min(k-cnt, len(arr[x+d]))
				cnt += w
				tmp += w * d
			}

			d++
		}
		if tmp < cost {
			cost = tmp
			best = []int{x}
		} else if tmp == cost {
			best = append(best, x)
		}
	}
	n := len(s)

	buf := make([]byte, n)
	get := func(x int) string {
		copy(buf, s)

		for cnt, d := len(arr[x]), 1; cnt < k; d++ {
			if x+d < 10 {
				w := min(k-cnt, len(arr[x+d]))
				for i := 0; i < w; i++ {
					buf[arr[x+d][i]] = byte(x + '0')
				}
				cnt += w
			}
			if x-d >= 0 {
				w := min(k-cnt, len(arr[x-d]))
				for i := len(arr[x-d]) - w; i < len(arr[x-d]); i++ {
					buf[arr[x-d][i]] = byte(x + '0')
				}
				cnt += w
			}
		}

		return string(buf)
	}

	for _, x := range best {
		tmp := get(x)
		if len(res) == 0 || tmp < res {
			res = tmp
		}
	}

	return
}
