package stripmargin

import (
	"strings"
	"unicode"
)

// Accepts a string and a `marginChar` rune to strip margins
// from a multiline string similar to Scala's stripMargin method
func StripMarginWith(str string, marginChar rune) string {
	lines := strings.Split(str, "\n")

	for i, line := range lines {
		strippedLine := strings.TrimLeftFunc(line, unicode.IsSpace)
		if len(strippedLine) > 0 && strippedLine[0] == byte(marginChar) {
			strippedLine = strippedLine[1:]
		}

		lines[i] = strippedLine
	}

	return strings.Join(lines, "\n")
}

// Accepts a string and strips margins from a multiline string
// using `|` similar to Scala's stripMargin method
func StripMargin(str string) string {
	return StripMarginWith(str, '|')
}
