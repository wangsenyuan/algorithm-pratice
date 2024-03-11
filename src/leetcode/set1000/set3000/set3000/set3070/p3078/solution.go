package p3078

import "strings"

func shortestSubstrings(arr []string) []string {
	ans := make([]string, len(arr))

	for i, cur := range arr {
		for l := 1; l <= len(cur); l++ {
			for k := 0; k+l <= len(cur); k++ {
				ok := true
				for j := 0; j < len(arr); j++ {
					if i == j {
						continue
					}
					if strings.Contains(arr[j], cur[k:k+l]) {
						ok = false
						break
					}
				}
				if ok && (len(ans[i]) == 0 || ans[i] > cur[k:k+l]) {
					ans[i] = cur[k : k+l]
				}
			}
			if len(ans[i]) > 0 {
				break
			}
		}
	}
	return ans
}
