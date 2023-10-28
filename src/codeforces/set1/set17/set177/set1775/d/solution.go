package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)
	s, x := readTwoNums(reader)
	res := solve(a, s, x)
	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(a []int, s int, t int) []int {
	// 每个数最多有lg(n)个质因数，每个数都移动到它的质因数，然后移动到其他地方
	x := maxOf(a)
	var primes []int
	lpf := make([]int, x+1)
	for i := 2; i <= x; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes); j++ {
			if i*primes[j] > x {
				break
			}
			lpf[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}
	n := len(a)

	pos := make([]int, x+1)
	for i := 0; i < len(primes); i++ {
		pos[primes[i]] = i + n
	}

	ids := make([][]int, x+1)
	for i := 0; i < n; i++ {
		if ids[a[i]] == nil {
			ids[a[i]] = make([]int, 0, 1)
		}
		ids[a[i]] = append(ids[a[i]], i)
	}

	b := make([]int, n+len(primes))
	copy(b, a)
	copy(b[n:], primes)

	m := len(b)
	dist := make([]int, m)
	for i := 0; i < m; i++ {
		dist[i] = -1
	}
	que := make([]int, m)
	var front, tail int

	pushFront := func(x int) {
		que[front] = x
		front++
	}
	popTail := func() int {
		u := que[tail]
		tail++
		return u
	}
	prev := make([]int, m)
	prev[s-1] = -1
	dist[s-1] = 0
	pushFront(s - 1)

	for tail < front {
		u := popTail()
		num := b[u]
		if u < n {
			for num > 1 {
				f := lpf[num]
				if dist[pos[f]] < 0 {
					prev[pos[f]] = u
					dist[pos[f]] = dist[u] + 1
					pushFront(pos[f])
				}
				num /= f
			}
		} else {
			// a prime factor
			for j := num; j <= x; j += num {
				for _, k := range ids[j] {
					if dist[k] < 0 {
						prev[k] = u
						dist[k] = dist[u] + 1
						pushFront(k)
					}
				}
			}
		}
	}

	if dist[t-1] < 0 {
		return nil
	}

	var ans []int

	for i := t - 1; i >= 0; i = prev[i] {
		if i < n {
			ans = append(ans, i+1)
		}
	}

	reverse(ans)

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func maxOf(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
