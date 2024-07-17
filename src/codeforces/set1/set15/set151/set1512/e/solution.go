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
		nums := readNNums(reader, 4)
		res := solve(nums[0], nums[1], nums[2], nums[3])
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
		}
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(n int, l int, r int, s int) []int {
	ms := (1 + r - l + 1) * (r - l + 1) / 2
	if s < ms {
		return nil
	}
	var sum int
	var arr []int
	for i := 0; i < r-l+1; i++ {
		arr = append(arr, n-i)
		sum += n - i
	}
	if s > sum {
		return nil
	}
	var other []int
	for i := 1; i < n-(r-l); i++ {
		other = append(other, i)
	}
	// 使用arr[i]去替换
	for i, j := 0, 0; sum > s; i++ {
		diff := sum - s
		// sum - arr[i] + other[j] = s
		if arr[i]-other[j] < diff {
			sum -= arr[i]
			sum += other[j]
			arr[i], other[j] = other[j], arr[i]
			j++
			continue
		}
		if arr[i]-other[len(other)-1] <= diff {
			v := arr[i] - diff
			v--
			sum -= arr[i]
			sum += other[v]
			arr[i], other[v] = other[v], arr[i]
			break
		}
	}
	res := make([]int, n)
	copy(res, other[:l-1])
	copy(res[l-1:], arr)
	copy(res[r:], other[l-1:])
	return res
}
