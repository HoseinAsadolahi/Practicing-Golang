package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16

	x := fullJustify(words, maxWidth)
	for _, line := range x {
		fmt.Println(line)
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var justifiedLines []string
	var currentLine []string
	var currentLineLength int

	startNewLine := func() {
		if len(currentLine) > 0 {
			justifiedLines = append(justifiedLines, justifyLine(currentLine, maxWidth, currentLineLength))
		}
		currentLine = []string{}
		currentLineLength = 0
	}

	for _, word := range words {
		if currentLineLength+len(currentLine)+len(word) <= maxWidth {
			currentLine = append(currentLine, word)
			currentLineLength += len(word)
		} else {
			startNewLine()
			currentLine = append(currentLine, word)
			currentLineLength += len(word)
		}
	}
	leftJustifyLine := strings.Join(currentLine, " ") + strings.Repeat(" ", maxWidth-len(strings.Join(currentLine, " ")))

	justifiedLines = append(justifiedLines, leftJustifyLine)

	return justifiedLines
}

func justifyLine(line []string, maxWidth int, currentLineLength int) string {
	if len(line) == 1 || len(line) == 0 {
		return line[0] + strings.Repeat(" ", maxWidth-currentLineLength)
	}

	neededSpaces := maxWidth - currentLineLength
	baseSpaces := neededSpaces/len(line) - 1
	extraSpaces := neededSpaces%len(line) - 1

	space := strings.Repeat(" ", baseSpaces)

	justifiedLine := line[0]

	for i := 1; i < len(line); i++ {
		if i <= extraSpaces {
			justifiedLine += space + " " + line[i]
		} else {
			justifiedLine += space + line[i]
		}
	}

	return justifiedLine
}
