package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')

	s = s[:len(s)-1]

	fmt.Println(solve(s))
}

func solve(s string) int64 {
	n := len(s)
	pf := make([]int, n)
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, 3)
	}

	for i := 0; i < n; i++ {
		pf[i] = int(s[i]-'0') % 3

		if i > 0 {
			pf[i] += pf[i-1]
			pf[i] %= 3
			copy(sum[i], sum[i-1])
		}

		if !(i+1 < n && s[i+1] == '0') {
			sum[i][pf[i]]++
		}
	}

	count := func(lf, rg int, val int) int {
		if lf > rg {
			return 0
		}
		res := sum[rg][val]
		if lf > 0 {
			res -= sum[lf-1][val]
		}
		return res
	}

	var ans int64

	l := 0
	r := -1
	z := make([]int, n)
	for i := 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[l+(r-i)])
		}
		for i-z[i] >= 0 && i+z[i] < n && s[i-z[i]] == s[i+z[i]] {
			z[i]++
		}

		cur := int(s[i]-'0') % 3
		want := ((3-cur)*2 + cur) % 3

		lf := i - z[i] + 1

		ans += int64(count(max(0, lf-1), i-1, (pf[i]-want+3)%3))

		if lf == 0 && s[0] != '0' && pf[i] == want {
			ans++
		}
		if s[i] == '0' {
			ans++
		}
		if i+z[i]-1 >= r {
			l = i - z[i] + 1
			r = i + z[i] - 1
		}
	}

	for i := 0; i < n; i++ {
		z[i] = 0
	}

	l = 0
	r = -1

	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[l+(r-i)+1])
		}

		for i-z[i]-1 >= 0 && i+z[i] < n && s[i-z[i]-1] == s[i+z[i]] {
			z[i]++
		}

		if z[i] > 0 {
			var want int
			lf := i - z[i]
			ans += int64(count(max(0, lf-1), i-2, (pf[i-1]-want+3)%3))

			if lf == 0 && s[0] != '0' && pf[i-1] == want {
				ans++
			}
		}

		if i+z[i]-1 > r {
			l = i - z[i]
			r = i + z[i] - 1
		}
	}

	return ans
}

func max(a, b int) int {
	return a + b - min(a, b)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(s string) int64 {
	buf := addBoundaries(s)

	n := len(buf)

	cnt := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		cnt[i] = make([]int, 3)
	}

	var sum int
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			cnt[i+1][j] = cnt[i][j]
		}
		if buf[i] == '0' || buf[i] == '$' {
			continue
		}
		x := int(buf[i]-'0') % 3
		sum += x
		sum %= 3
		cnt[i+1][sum]++
	}

	p := make([]int, len(buf))

	var center int
	var right int

	for i := 1; i < len(buf); i++ {
		var a, b int
		if i > right {
			p[i] = 0
			a = i - 1
			b = i + 1
		} else {
			i2 := 2*center - i
			if p[i2] < right-i-1 {
				p[i] = p[i2]
				a = -1
			} else {
				p[i] = right - i
				b = right + 1
				a = 2*i - b
			}
		}

		for a >= 0 && b < len(buf) && buf[a] == buf[b] {
			p[i]++
			a--
			b++
		}
		if i+p[i] > right {
			center = i
			right = i + p[i]
		}
	}

	var res int64
	sum = 0
	for i := 0; i < len(buf); i++ {
		if buf[i] != '$' {
			sum += int(buf[i] - '0')
			sum %= 3
		}

		j := i + p[i]

		if buf[i] == '$' {
			res += int64(cnt[j][sum] - cnt[i][sum])
			continue
		}

		x := int(buf[i]-'0') % 3

		if x == 0 {
			res++
		}

		for y := 0; y < 3; y++ {
			if (2*y+x)%3 == 0 {
				tmp := (sum + y) % 3
				res += int64(cnt[j][tmp] - cnt[i+1][tmp])
				break
			}
		}
	}

	return res
}

func addBoundaries(s string) []byte {
	n := len(s)
	buf := make([]byte, 2*n+1)

	for i := 0; i < n; i++ {
		buf[2*i] = '$'
		buf[2*i+1] = s[i]
	}
	buf[2*n] = '$'

	return buf
}
