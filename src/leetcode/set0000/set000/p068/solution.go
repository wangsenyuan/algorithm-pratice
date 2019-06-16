package main

import "fmt"

func main() {
	//fmt.Println("aaa")
	words := []string{"What", "must", "be", "shall", "be."}
	rs := fullJustify(words, 12)
	for _, r := range rs {
		fmt.Println(r)
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var rs []string

	i := 0
	width := len(words[0])
	for j := 1; j < len(words); j++ {
		w := width + len(words[j]) + (j - i)
		if w > maxWidth {
			line := justifyLine(words[i:j], maxWidth, false)
			//fmt.Printf("got: %s\n", line)
			rs = append(rs, line)
			i = j
			width = len(words[j])
			continue
		}
		width += len(words[j])
	}

	if i < len(words) {
		rs = append(rs, justifyLine(words[i:], maxWidth, true))
	}

	return rs
}

func justifyLine(words []string, lineWidth int, leftPad bool) string {
	width := 0
	for _, word := range words {
		width += len(word)
	}

	allSpaceWidth := lineWidth - width
	if len(words) == 1 {
		return words[0] + padSpace(allSpaceWidth)
	}
	str := words[0]
	if !leftPad {
		spaceWidth := allSpaceWidth / (len(words) - 1)
		left := allSpaceWidth - spaceWidth * (len(words) - 1)

		for i := 1; i < len(words); i++ {
			spaceCount := spaceWidth
			if left > 0 {
				spaceCount++
				left--
			}
			str += padSpace(spaceCount) + words[i]
		}

		return str
	}

	for i := 1; i < len(words); i++ {
		str += " " + words[i];
		allSpaceWidth--
	}
	str += padSpace(allSpaceWidth)
	return str
}

func padSpace(count int) string {
	str := ""

	for i := 0; i < count; i++ {
		str += " "
	}

	return str
}
