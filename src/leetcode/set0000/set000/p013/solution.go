package p013

var anchors = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	var res int

	for i := 0; i < len(s); {
		if i+1 < len(s) {
			x := s[i : i+2]
			if y, ok := anchors[x]; ok {
				res += y
				i += 2
				continue
			}
		}
		y := anchors[s[i:i+1]]
		res += y
		i++
	}
	return res
}
