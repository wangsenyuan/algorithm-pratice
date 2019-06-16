package main

import "fmt"

func main() {
	fmt.Println(fractionToDecimal(-50, 8))
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var buf []byte

	if sign(numerator) != sign(denominator) {
		buf = append(buf, '-')
	}

	divident := abs(int64(numerator))
	divisor := abs(int64(denominator))
	buf = append(buf, fmt.Sprintf("%d", divident/divisor)...)
	reminder := divident % divisor

	if reminder == 0 {
		return string(buf)
	}

	buf = append(buf, '.')

	pos := make(map[int64]int)
	for reminder > 0 {
		p := pos[reminder]
		if p > 0 {
			buf = insertLoopAt(buf, p)
			break
		}
		pos[reminder] = len(buf)
		reminder *= 10
		buf = append(buf, fmt.Sprintf("%d", reminder/divisor)...)
		reminder %= divisor
	}

	return string(buf)
}

func insertLoopAt(buf []byte, p int) []byte {
	sb := make([]byte, len(buf)+2)

	copy(sb, buf[:p])

	sb[p] = '('

	copy(sb[p+1:], buf[p:])

	sb[len(sb)-1] = ')'

	return sb
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	return 1
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
