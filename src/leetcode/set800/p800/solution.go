package p800

import (
	"strconv"
)

func similarRGB(color string) string {
	color = color[1:]
	dist := -(1 << 30)
	var ans string
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				ss := str(x) + str(x) + str(y) + str(y) + str(z) + str(z)

				tmp := similarity(color, ss)
				//fmt.Printf("[debug] %s => similairty %d\n", ss, tmp)
				if tmp > dist {
					ans = "#" + ss
					dist = tmp
				}
			}
		}
	}
	return ans
}

func similarity(color string, color1 string) int {
	a := 16*value(color[0]) + value(color[1])
	b := 16*value(color[2]) + value(color[3])
	c := 16*value(color[4]) + value(color[5])

	x := 16*value(color1[0]) + value(color1[1])
	y := 16*value(color1[2]) + value(color1[3])
	z := 16*value(color1[4]) + value(color1[5])
	tmp := (a-x)*(a-x) + (b-y)*(b-y) + (c-z)*(c-z)
	return -tmp
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func str(x int) string {
	var y string
	if x >= 0 && x <= 9 {
		y = strconv.Itoa(x)
	} else {
		y = string(byte(x - 10 + 'a'))
	}
	return y
}

func value(b byte) int {
	if b >= '0' && b <= '9' {
		return int(b - '0')
	}
	return int(b-'a') + 10
}
