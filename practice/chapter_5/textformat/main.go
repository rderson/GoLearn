package textformat

import (
	"strings"

	"golang.org/x/net/html"
)

func readAll(data string, n *html.Node) string {
	if n.Type == html.TextNode {
		data += n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		data = readAll(data, c)
	}
	return data
}

func removeTabs(s string) string {
    // Replace tabs '\t' with an empty string
    return strings.ReplaceAll(s, "\t", "")
}

func removeExtraNewlines(s string) string {
    var result strings.Builder
    newlineCount := 0

    for _, char := range s {
        if char == '\n' {
            newlineCount++
        } else {
            newlineCount = 0
        }

        // Append the character if less than or equal to two newlines in a row
        if newlineCount <= 2 {
            result.WriteRune(char)
        }
    }

    return result.String()
}

func FormatText(s string) string {
    s = removeTabs(s)
    s = removeExtraNewlines(s)
    return s
}

func ReadAndFormatTextElements(data string, n *html.Node) string {
	data = readAll(data, n)
	data = FormatText(data)
	return data
}