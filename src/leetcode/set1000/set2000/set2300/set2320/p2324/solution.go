package p2324

func decodeMessage(key string, message string) string {
	mapping := make([]int, 26)
	for i := 0; i < 26; i++ {
		mapping[i] = -1
	}
	var p int
	for i := 0; i < len(key); i++ {
		if key[i] >= 'a' && key[i] <= 'z' {
			x := int(key[i] - 'a')
			if mapping[x] < 0 {
				mapping[x] = p
				p++
			}
		}
	}

	buf := make([]byte, len(message))

	for i := 0; i < len(message); i++ {
		if message[i] >= 'a' && message[i] <= 'z' {
			x := int(message[i] - 'a')
			y := mapping[x]
			buf[i] = byte(y + 'a')
		} else {
			buf[i] = message[i]
		}
	}

	return string(buf)
}
