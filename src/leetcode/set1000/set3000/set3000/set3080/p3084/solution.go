package p3084

import "fmt"

func isSubstringPresent(s string) bool {
	vis := make(map[string]bool)

	for i := 0; i+2 <= len(s); i++ {
		vis[s[i:i+2]] = true
	}

	for i := len(s) - 1; i > 0; i-- {
		x := fmt.Sprintf("%c%c", s[i], s[i-1])
		if vis[x] {
			return true
		}
	}
	return false
}
