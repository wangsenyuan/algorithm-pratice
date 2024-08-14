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
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		s := readString(reader)
		res := solve(s)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

type pair struct {
	first  int
	second int
}

func solve(s string) string {
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'a')]++
	}

	var arr []pair

	for i := 0; i < 26; i++ {
		if cnt[i] == 0 {
			continue
		}
		arr = append(arr, pair{cnt[i], i})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first < arr[j].first || arr[i].first == arr[j].first && arr[i].second < arr[j].second
	})

	if len(arr) == 1 {
		// same character, no chocie
		return s
	}
	n := len(s)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = ' '
	}
	if arr[0].first == 1 {
		// best, f(s) = 0
		buf[0] = byte(arr[0].second + 'a')
		it := 1
		for i := 0; i < 26; i++ {
			if i == arr[0].second {
				continue
			}
			for cnt[i] > 0 {
				buf[it] = byte(i + 'a')
				it++
				cnt[i]--
			}
		}
		return string(buf)
	}
	// f(s) = 1
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].second < arr[j].second
	})

	buf[0] = byte(arr[0].second + 'a')
	cnt[arr[0].second]--

	if cnt[arr[0].second] < n/2+1 {
		for i := 1; cnt[arr[0].second] > 0; i += 2 {
			buf[i] = byte(arr[0].second + 'a')
			cnt[arr[0].second]--
		}
		var it int
		for i := 0; i < 26; i++ {
			if i == arr[0].second {
				continue
			}
			for cnt[i] > 0 {
				for buf[it] != ' ' {
					it++
				}
				buf[it] = byte(i + 'a')
				cnt[i]--
			}
		}
		return string(buf)
	}
	if len(arr) == 2 {
		for i := 1; i < n; i++ {
			if cnt[arr[1].second] > 0 {
				buf[i] = byte(arr[1].second + 'a')
				cnt[arr[1].second]--
			} else {
				buf[i] = buf[0]
			}
		}
		return string(buf)
	}

	buf[1] = byte(arr[1].second + 'a')
	cnt[arr[1].second]--
	it := 2
	for cnt[arr[0].second] > 0 {
		buf[it] = buf[0]
		it++
		cnt[arr[0].second]--
	}
	buf[it] = byte(arr[2].second + 'a')
	it++
	cnt[arr[2].second]--

	for i := 0; i < 26; i++ {
		for cnt[i] > 0 {
			buf[it] = byte(i + 'a')
			it++
			cnt[i]--
		}
	}

	return string(buf)
}
