package p916

func wordSubsets(A []string, B []string) []string {
	wc := make([]int, 26)
	tc := make([]int, 26)
	for i := 0; i < len(B); i++ {
		for j := 0; j < 26; j++ {
			tc[j] = 0
		}
		for j := 0; j < len(B[i]); j++ {
			tc[int(B[i][j]-'a')]++
		}
		for j := 0; j < 26; j++ {
			if tc[j] > wc[j] {
				wc[j] = tc[j]
			}
		}
	}
	res := make([]string, len(A))
	var p int
	for i := 0; i < len(A); i++ {
		for j := 0; j < 26; j++ {
			tc[j] = 0
		}
		for j := 0; j < len(A[i]); j++ {
			tc[int(A[i][j]-'a')]++
		}
		cand := true
		for j := 0; j < 26; j++ {
			if tc[j] < wc[j] {
				cand = false
				break
			}
		}
		if cand {
			res[p] = A[i]
			p++
		}
	}
	return res[:p]
}
