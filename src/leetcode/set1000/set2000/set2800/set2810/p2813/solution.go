package p2813

func finalString(s string) string {
	buf := []byte(s)
	n := len(s)
	var j int
	for i := 0; i < n; i++ {
		if s[i] == 'i' {
			reverse(buf[:j])
			continue
		}
		buf[j] = s[i]
		j++
	}
	return string(buf[:j])
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
