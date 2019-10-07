package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

func renderItalic(markdown string) string {
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	return markdown
}

func renderBold(markdown string) string {
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	return markdown
}

func formatHeader(markdown string) (string, int) {
	pos := 0
	header := 0
	html := ""
	for {
		char := markdown[pos]
		if char == '#' {
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			html += fmt.Sprintf("<h%d>", header)
			pos++
			continue
		}
		if char == '\n' {
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
		}
		html += string(char)
		pos++
		if pos >= len(markdown) {
			break
		}
	}
	return html, header

}

func formatList(markdown string) (string, bool) {
	pos := 0
	list := 0
	html := ""
	for {
		char := markdown[pos]
		if char == '*' {
			if list == 0 {
				html += "<ul>"
			}
			html += "<li>"
			list++
			pos += 2
			continue
		}
		if char == '\n' {
			if list > 0 {
				html += "</li>"
			}
			pos++
			continue
		}
		html += string(char)
		pos++
		if pos >= len(markdown) {
			break
		}
	}

	return html, list > 0
}

func htmlize(markdown string) string {
	html := ""
	markdown, header := formatHeader(markdown)
	html, closeList := formatList(markdown)
	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}
	if closeList {
		return html + "</li></ul>"
	}
	return "<p>" + html + "</p>"

}

// Render translates markdown to HTML
func Render(markdown string) string {
	markdown = renderBold(markdown)
	markdown = renderItalic(markdown)

	return htmlize(markdown)
}
