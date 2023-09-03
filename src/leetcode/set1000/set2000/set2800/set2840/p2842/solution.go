package p2842

import "sort"

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, mod-2)
}

func countKSubsequencesWithMaxBeauty(s string, k int) int {
	if k > 26 {
		return 0
	}
	var C [27][27]int
	C[0][0] = 1
	for i := 1; i <= 26; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = add(C[i-1][j-1], C[i-1][j])
		}
	}

	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[int(s[i]-'a')]++
	}
	var arr []int
	for _, v := range freq {
		arr = append(arr, v)
	}
	sort.Ints(arr)
	reverse(arr)

	if arr[k-1] == 0 {
		return 0
	}

	r := k - 1
	for r+1 < 26 && arr[r+1] == arr[k-1] {
		r++
	}
	var l int
	ans := 1
	for l < k-1 && arr[l] != arr[k-1] {
		ans = mul(ans, arr[l])
		l++
	}
	ans = mul(ans, mul(C[r-l+1][k-l], pow(arr[k-1], k-l)))

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
