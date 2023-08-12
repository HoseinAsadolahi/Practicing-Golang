package main

import (
	"fmt"
	"strings"
)

func main() {
	x := fullJustify([]string{"ask","not","what","your","country","can","do","for","you","ask","what","you","can","do","for","your","country"}, 16)
	for _, v := range x {
		fmt.Println(v, len(v))
	}
}

func fullJustify(words []string, maxWidth int) []string {
	n, currentLen := len(words), 0
	out := make([]string, 0, len(words))
	for i := 0; i < n; i++ {
		if currentLen+len(words[i]) > maxWidth {
			out = append(out, justify(words[:i], maxWidth, currentLen - 1, false))
			words = words[i:]
			i = 0
			currentLen = len(words[i]) + 1
			n = len(words)
		} else {
			currentLen += len(words[i]) + 1
		}
	}
	out = append(out, justify(words, maxWidth, currentLen - 1, true))
	return out
}

func justify(words []string, maxWidth, currentLen int, leftJustify bool) string {
	if len(words) == 1 {
		return words[0] + getSpaces(maxWidth - currentLen)
	}
	if leftJustify {
		return strings.Join(words, " ") + getSpaces(maxWidth - currentLen)
	}
	spacePlace := (len(words) - 1)
	totalSpaces := maxWidth - currentLen + spacePlace
	out := ""
	for i := 0; i < len(words); i++ {
		var spaces int
		if spacePlace != 0 {
			if totalSpaces % spacePlace != 0 {
				spaces = totalSpaces / spacePlace + 1
			} else {
				spaces = totalSpaces / spacePlace
			}
		}
		totalSpaces -= spaces
		spacePlace--
		out += words[i] + getSpaces(spaces)
	}
	return out
}

func getSpaces(n int) string {
	x := ""
	for i := 0; i < n; i++ {
		x += " "
	}
	return x
}