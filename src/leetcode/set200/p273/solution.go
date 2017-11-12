package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numberToWords(123456789))
}

var dict = make(map[int]string)

func init() {
	dict[0] = "Zero"
	dict[1] = "One"
	dict[2] = "Two"
	dict[3] = "Three"
	dict[4] = "Four"
	dict[5] = "Five"
	dict[6] = "Six"
	dict[7] = "Seven"
	dict[8] = "Eight"
	dict[9] = "Nine"
	dict[10] = "Ten"
	dict[11] = "Eleven"
	dict[12] = "Twelve"
	dict[13] = "Thirteen"
	dict[14] = "Fourteen"
	dict[15] = "Fifteen"
	dict[16] = "Sixteen"
	dict[17] = "Seventeen"
	dict[18] = "Eighteen"
	dict[19] = "Nineteen"
	dict[20] = "Twenty"
	dict[30] = "Thirty"
	dict[40] = "Forty"
	dict[50] = "Fifty"
	dict[60] = "Sixty"
	dict[70] = "Seventy"
	dict[80] = "Eighty"
	dict[90] = "Ninety"
	dict[100] = "Hundred"
	dict[1000] = "Thousand"
	dict[1000000] = "Million"
	dict[1000000000] = "Billion"
}

func numberToWords(num int) string {
	steps := []int{1000000000, 1000000, 1000, 1}
	var res []string

	for _, step := range steps {
		if num/step > 0 {
			words := readWord(num / step)
			res = append(res, words...)
			if step > 1 {
				res = append(res, dict[step])
			}
		}
		num %= step
	}

	if len(res) == 0 {
		return "Zero"
	}

	return strings.Join(res, " ")
}

func readWord(n int) []string {
	var words []string
	if n/100 > 0 {
		words = append(words, dict[n/100])
		words = append(words, dict[100])
		n = n % 100
	}

	if n >= 20 {
		words = append(words, dict[n-n%10])
		n = n % 10
	}
	if n > 0 {
		words = append(words, dict[n])
	}
	return words
}
