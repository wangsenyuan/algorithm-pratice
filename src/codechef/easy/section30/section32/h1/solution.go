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
		reader.ReadBytes('\n')

		grid := make([][]int, 3)
		for i := 0; i < 3; i++ {
			grid[i] = readNNums(reader, 3)
		}
		res := solve(grid)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const N = 362880

var dp []int

var primes map[int]bool
var pos map[int64]int

func init() {
	primes = make(map[int]bool)
	set := make([]bool, 20)
	for i := 2; i < 20; i++ {
		if !set[i] {
			primes[i] = true
			for j := i + i; j < 20; j += i {
				set[j] = true
			}
		}
	}

	arr := make([]int64, 9)

	next := func(num int64) int64 {
		for i := 8; i >= 0; i-- {
			arr[i] = num % 10
			num /= 10
		}
		i := len(arr) - 2
		for i >= 0 && arr[i] > arr[i+1] {
			i--
		}
		if i < 0 {
			return -1
		}
		// arr[i] < arr[i+1]
		j := len(arr) - 1
		for arr[i] > arr[j] {
			j--
		}
		// j will at least be i+1
		arr[i], arr[j] = arr[j], arr[i]

		for a, b := i+1, len(arr)-1; a < b; a, b = a+1, b-1 {
			arr[a], arr[b] = arr[b], arr[a]
		}

		var ans int64

		for i := 0; i < len(arr); i++ {
			ans = ans*10 + arr[i]
		}
		return ans
	}

	nums := make([]int64, N)
	var cur int64 = 123456789
	dp = make([]int, N)
	pos = make(map[int64]int)
	for i := 0; i < N; i++ {
		nums[i] = cur
		pos[cur] = i
		cur = next(cur)
		dp[i] = -1
	}

	dp[0] = 0

	que := make([]int, N)
	var front, end int
	que[end] = 0
	end++

	for front < end {
		cur := que[front]
		front++
		num := nums[cur]
		neighors := move(num, arr)

		for _, x := range neighors {
			i := pos[x]
			if dp[i] < 0 {
				dp[i] = dp[cur] + 1
				que[end] = i
				end++
			}
		}
	}

}

func move(num int64, buf []int64) []int64 {
	// len(buf) = 9
	for i := 8; i >= 0; i-- {
		buf[i] = num % 10
		num /= 10
	}

	var res []int64

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			x := buf[i*3+j]
			if j+1 < 3 && primes[int(buf[i*3+j+1]+x)] {
				buf[i*3+j], buf[i*3+j+1] = buf[i*3+j+1], buf[i*3+j]
				res = append(res, toNum(buf))
				buf[i*3+j], buf[i*3+j+1] = buf[i*3+j+1], buf[i*3+j]
			}
			if i+1 < 3 && primes[int(buf[(i+1)*3+j]+x)] {
				buf[i*3+j], buf[i*3+j+3] = buf[i*3+j+3], buf[i*3+j]
				res = append(res, toNum(buf))
				buf[i*3+j], buf[i*3+j+3] = buf[i*3+j+3], buf[i*3+j]
			}
		}
	}

	return res
}

func toNum(arr []int64) int64 {
	var res int64
	for i := 0; i < len(arr); i++ {
		res = res*10 + arr[i]
	}
	return res
}

func solve(grid [][]int) int {
	buf := make([]int64, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			buf[i*3+j] = int64(grid[i][j])
		}
	}
	num := toNum(buf)

	i := pos[num]

	return dp[i]
}
