package p1923

const M1 = 1000000007
const M2 = 1000000009
const B1 = 1000007
const B2 = 1000009

func longestCommonSubpath(n int, paths [][]int) int {
	var m int = B1
	for i := 0; i < len(paths); i++ {
		if m > len(paths[i]) {
			m = len(paths[i])
		}
	}
	P1 := make([][]int64, len(paths))
	P2 := make([][]int64, len(paths))
	A1 := make([]int64, B1)
	A2 := make([]int64, B1)

	for i := 0; i < len(paths); i++ {
		P1[i] = getHash(paths[i], M1, B1)
		P2[i] = getHash(paths[i], M2, B2)
	}
	A1[0] = 1
	A2[0] = 1
	for i := 1; i < B1; i++ {
		A1[i] = A1[i-1] * B1 % M1
		A2[i] = A2[i-1] * B2 % M2
	}
	check := func(m int) bool {
		if m == 0 {
			return true
		}
		cnt1 := make(map[int64]bool)
		cnt2 := make(map[int64]bool)
		for r := 0; r < len(paths); r++ {
			cnt3 := make(map[int64]bool)
			cnt4 := make(map[int64]bool)

			for i := m - 1; i < len(paths[r]); i++ {
				tmp1 := (P1[r][i+1] + M1 - P1[r][i+1-m]*A1[m]%M1) % M1
				tmp2 := (P2[r][i+1] + M2 - P2[r][i+1-m]*A2[m]%M2) % M2
				if r == 0 {
					cnt3[tmp1] = true
					cnt4[tmp2] = true
				} else if cnt1[tmp1] && cnt2[tmp2] {
					cnt3[tmp1] = true
					cnt4[tmp2] = true
				}
			}
			cnt1, cnt2 = cnt3, cnt4
		}
		return len(cnt1) > 0
	}

	ans := search(m+1, func(i int) bool {
		return !check(i)
	})

	return ans - 1
}

func getHash(arr []int, mod int64, base int64) []int64 {
	res := make([]int64, len(arr)+1)
	res[0] = 0
	for i := 0; i < len(arr); i++ {
		res[i+1] = (res[i]*base%mod + int64(arr[i])) % mod
	}
	return res
}

func search(n int, fn func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
