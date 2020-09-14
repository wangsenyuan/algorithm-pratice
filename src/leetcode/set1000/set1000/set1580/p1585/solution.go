package p1585

func isTransformable(s string, t string) bool {
	cnt := make([]int, 10)

	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'0')]++
		cnt[int(t[i]-'0')]--

		if cnt[int(t[i]-'0')] == 0 {
			continue
		}

		for j := t[i] - 1; j >= '0'; j-- {
			if cnt[int(j-'0')] > 0 {
				return false
			}
		}
	}

	return true
}

func calc(s string) [][][]int {
	n := len(s)
	dp := make([][][]int, 10)
	for i := 0; i < 10; i++ {
		dp[i] = make([][]int, 0, n)
	}
	cnt := make([]int, 10)
	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - '0')
		tmp := make([]int, 10)
		copy(tmp, cnt)
		dp[x] = append(dp[x], tmp)
		cnt[x]++
	}
	return dp
}

func isTransformable2(s string, t string) bool {
	dp := calc(s)
	fp := calc(t)

	for i := 0; i < 10; i++ {
		if len(dp[i]) != len(fp[i]) {
			return false
		}
		for j := 0; j < len(dp[i]); j++ {
			for k := i + 1; k < 10; k++ {
				if dp[i][j][k] > fp[i][j][k] {
					return false
				}
			}
		}
	}
	return true
}

func isTransformable1(s string, t string) bool {
	n := len(s)
	if len(t) != n {
		return false
	}
	pos := make([][]int, 10)
	for i := 0; i < 10; i++ {
		pos[i] = make([]int, 0, n/10)
	}

	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		pos[x] = append(pos[x], i)
	}

	arr := make([]int, 2*n)

	for i := n; i < len(arr); i++ {
		arr[i] = int(s[i-n] - '0')
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = max(arr[2*i], arr[2*i+1])
	}

	update := func(p int, v int) {
		p += n
		arr[p] = v
		for p > 0 {
			arr[p>>1] = max(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) int {
		l += n
		r += n
		var res = -1
		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	i, j := n-1, n-1

	for i >= 0 && j >= 0 {
		if get(j, j+1) < 0 {
			j--
			continue
		}
		if s[j] == t[i] {
			update(j, -1)
			j--
			i--
			continue
		}
		if s[j] > t[i] {
			return false
		}
		//s[j] < t[i]
		// need to find some k, s[k] == t[i]
		x := int(t[i] - '0')
		if len(pos[x]) == 0 {
			return false
		}
		m := len(pos[x])
		for m > 0 && pos[x][m-1] > j {
			m--
		}
		if m < 0 {
			return false
		}
		k := pos[x][m-1]
		// then there can't be any y > x between k....j
		y := get(k+1, j)
		if y > x {
			return false
		}
		update(k, -1)
		pos[x] = pos[x][:m]
		i--
	}

	return i < 0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
