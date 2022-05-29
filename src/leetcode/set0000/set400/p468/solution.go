package p468

import "strings"

func validIPAddress(IP string) string {
	if validIPv4(IP) {
		return "IPv4"
	}

	if validIPv6(IP) {
		return "IPv6"
	}

	return "Neither"
}

func validIPv6(IP string) bool {
	parts := strings.Split(IP, ":")

	if len(parts) != 8 {
		return false
	}

	for _, part := range parts {
		if !validIPv6Part(part) {
			return false
		}
	}
	return true
}

func validIPv6Part(part string) bool {
	if len(part) == 0 || len(part) > 4 {
		return false
	}

	for i := 0; i < len(part); i++ {
		c := part[i]
		if c >= '0' && c <= '9' {
			continue
		}
		if c >= 'a' && c <= 'f' {
			continue
		}

		if c >= 'A' && c <= 'F' {
			continue
		}
		return false
	}

	return true
}

func validIPv4(IP string) bool {
	parts := strings.Split(IP, ".")

	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		if !validIPv4Part(part) {
			return false
		}
	}

	return true
}

func validIPv4Part(part string) bool {
	if len(part) == 0 {
		return false
	}
	num := 0
	for i := 0; i < len(part); i++ {
		c := part[i]
		if c < '0' || c > '9' {
			return false
		}
		if i == 0 && c == '0' && i < len(part)-1 {
			return false
		}
		num = num*10 + int(c-'0')
	}

	return num < 256
}
