package p1999

func reversePrefix(word string, ch byte) string {
	buf := []byte(word)

	for i := 0; i < len(word); i++ {
		if buf[i] == ch {
			reverse(buf[:i+1])
			break
		}
	}
	return string(buf)
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
