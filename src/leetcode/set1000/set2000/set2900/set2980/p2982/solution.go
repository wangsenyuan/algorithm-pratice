package p2982

func canMakePalindromeQueries(s string, queries [][]int) []bool {
	n := len(s)
	first := s[:n/2]
	second := reverse(s[n/2:])
	n /= 2
	L := make([]int, n)

	pa := NewPref(n, 26)
	pb := NewPref(n, 26)

	for i := 0; i < n; i++ {
		if first[i] == second[i] {
			L[i] = i
			if i > 0 {
				L[i] = L[i-1]
			}
		} else {
			L[i] = i + 1
		}
		pa.Add(i, int(first[i]-'a'))
		pb.Add(i, int(second[i]-'a'))
	}

	possible := true

	for i := 0; i < 26; i++ {
		if pa.GetPrefSum(n-1, i) != pb.GetPrefSum(n-1, i) {
			possible = false
			break
		}
	}

	ha := make([]int, 26)
	hb := make([]int, 26)

	check := func(a, b, c, d int) bool {
		if !possible {
			return false
		}
		if L[n-1] == 0 {
			// palindrome always
			return true
		}
		if b < c || d < a {
			// no overlap at all
			x := min(a, c)
			if x > 0 && L[x-1] != 0 {
				return false
			}
			for i := 0; i < 26; i++ {
				diff := pa.GetRangeSum(a, b, i) - pb.GetRangeSum(a, b, i)
				if diff != 0 {
					return false
				}
			}
			for i := 0; i < 26; i++ {
				diff := pa.GetRangeSum(c, d, i) - pb.GetRangeSum(c, d, i)
				if diff != 0 {
					return false
				}
			}
			y := max(b, d)
			if y < n-1 && y+1 < L[n-1] {
				return false
			}
			return true
		}
		// if overlap totally
		if a <= c && d <= b || c <= a && b <= d {
			x := min(a, c)
			y := max(b, d)
			if x > 0 && L[x-1] != 0 {
				return false
			}
			for i := 0; i < 26; i++ {
				diff := pa.GetRangeSum(x, y, i) - pb.GetRangeSum(x, y, i)
				if diff != 0 {
					return false
				}
			}
			if y < n-1 && y+1 < L[n-1] {
				return false
			}
			return true
		}
		// have some overlap
		if a < c {
			if a > 0 && L[a-1] != 0 {
				return false
			}

			for i := 0; i < 26; i++ {
				ha[i] = pa.GetRangeSum(a, b, i)
				ha[i] -= pb.GetRangeSum(a, c-1, i)
				if ha[i] < 0 {
					return false
				}
				hb[i] = pb.GetRangeSum(c, d, i)
				hb[i] -= pa.GetRangeSum(b+1, d, i)
				if hb[i] < 0 {
					return false
				}
				if ha[i] != hb[i] {
					return false
				}
			}
			if d < n-1 && d+1 < L[n-1] {
				return false
			}
			return true
		}
		// c < a
		if c > 0 && L[c-1] > 0 {
			return false
		}
		for i := 0; i < 26; i++ {
			ha[i] = pb.GetRangeSum(c, d, i)
			ha[i] -= pa.GetRangeSum(c, a-1, i)
			if ha[i] < 0 {
				return false
			}
			hb[i] = pa.GetRangeSum(a, b, i)
			hb[i] -= pb.GetRangeSum(d+1, b, i)
			if hb[i] < 0 {
				return false
			}
			if ha[i] != hb[i] {
				return false
			}
		}
		if b+1 < n && L[n-1] > b+1 {
			return false
		}
		return true
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		a, b, c, d := cur[0], cur[1], cur[2], cur[3]
		c -= n
		d -= n
		c = n - c - 1
		d = n - d - 1
		ans[i] = check(a, b, d, c)
	}

	return ans
}

type Pref struct {
	arr [][]int
	dim int
}

func NewPref(n int, m int) *Pref {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	return &Pref{arr, m}
}

func (p *Pref) Add(i int, v int) {
	if i > 0 {
		copy(p.arr[i], p.arr[i-1])
	}
	p.arr[i][v]++
}

func (p *Pref) GetPrefSum(i int, j int) int {
	return p.arr[i][j]
}

func (p *Pref) GetRangeSum(l int, r int, j int) int {
	res := p.GetPrefSum(r, j)
	if l > 0 {
		res -= p.GetPrefSum(l-1, j)
	}
	return res
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
