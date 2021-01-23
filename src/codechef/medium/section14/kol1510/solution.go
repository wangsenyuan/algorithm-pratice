package main

func main() {

}

func solve(k int, S []string) string {

	x := 0
	for i := 0; i < k; i++ {
		if x < len(S[i]) {
			x = len(S[i])
		}
	}

	contains := func(a []byte, b string) bool {
		var i, j int

		for i < len(a) && j < len(b) {
			if a[i] == b[j] {
				i++
				j++
				continue
			}
			i++
		}

		return j == len(b)
	}

	check := func(cur []byte) bool {
		for i := 0; i < k; i++ {
			if !contains(cur, S[i]) {
				return false
			}
		}

		return true
	}

	buf := make([]byte, 16)

	for l := x; l <= 16; l++ {
		for cur := 0; cur < 1<<uint(l); cur++ {
			decode(cur, l, buf)
			if check(buf[:l]) {
				return string(buf[:l])
			}
		}

	}

	return ""
}

func encode(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		x := s[i]
		var y int
		if x == 'G' {
			y = 1
		}
		res = (res << 1) | y
	}
	return res
}

func decode(x int, n int, buf []byte) {
	var i int
	for x > 0 {
		y := x & 1
		if y == 0 {
			buf[i] = 'B'
		} else {
			buf[i] = 'G'
		}
		x >>= 1
		i++
	}

	for i < n {
		buf[i] = 'B'
		i++
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
}
