package main

func minCut(s string) int {
	cut := make([]int, len(s)+1)

	for i := range cut {
		cut[i] = i - 1
	}

	for i := range s {
		for j := 0; i-j >= 0 && i+j < len(s) && s[i+j] == s[i-j]; j++ {
			if cut[i-j]+1 < cut[i+j+1] {
				cut[i+j+1] = cut[i-j] + 1
			}
		}

		for j := 1; i-j+1 >= 0 && i+j < len(s) && s[i+j] == s[i-j+1]; j++ {
			if cut[i-j+1]+1 < cut[i+j+1] {
				cut[i+j+1] = cut[i-j+1] + 1
			}
		}
	}

	return cut[len(s)]
}
