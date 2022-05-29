package p2287

import (
	"bytes"
	"fmt"
)

func discountPrices(sentence string, discount int) string {
	var buf bytes.Buffer

	for i := 0; i < len(sentence); {
		if sentence[i] == ' ' {
			buf.WriteByte(' ')
			i++
		} else if sentence[i] != '$' {
			for i < len(sentence) && sentence[i] != ' ' {
				buf.WriteByte(sentence[i])
				i++
			}
		} else {
			// a $
			j := i
			i++
			var price int64
			for i < len(sentence) && isNumber(sentence[i]) {
				price = price*10 + int64(sentence[i]-'0')
				i++
			}
			if i-j > 1 && (i == len(sentence) || sentence[i] == ' ') {
				// a price
				price *= int64(100 - discount)
				result := float64(price) / 100
				buf.WriteString(fmt.Sprintf("$%.2f", result))
			} else {
				i = j
				for i < len(sentence) && sentence[i] != ' ' {
					buf.WriteByte(sentence[i])
					i++
				}
			}
		}
	}
	return buf.String()
}

func isNumber(x byte) bool {
	return x >= '0' && x <= '9'
}
