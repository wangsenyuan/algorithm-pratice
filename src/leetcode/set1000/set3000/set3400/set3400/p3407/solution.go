package p3407

func hasMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	var dfs func(s string, p string) bool

	dfs = func(s string, p string) bool {
		if len(p) == 0 || len(s) == 0 && p == "*" {
			return true
		}
		if len(s) == 0 {
			return false
		}
		if s[0] == p[0] {
			return dfs(s[1:], p[1:])
		}
		if p[0] == '*' {
			return dfs(s[1:], p) || dfs(s[1:], p[1:]) || dfs(s, p[1:])
		}
		return false
	}

	for i := 0; n-i >= m-1; i++ {
		if dfs(s[i:], p) {
			return true
		}
	}

	return false
}
