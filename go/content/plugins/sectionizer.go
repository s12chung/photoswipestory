package plugins

import (
	"html/template"
	"regexp"
	"strings"
)

type Sectionizer struct{}

func NewSectionzier() *Sectionizer {
	return &Sectionizer{}
}

var imgRegex = regexp.MustCompile(`<img[^>]+>`)

func (sectionizer *Sectionizer) Sectionize(html string) string {
	hasImage := false
	sectionStart := `<section class="text">`
	sectionEnd := `</section>`

	replaced := imgRegex.ReplaceAllStringFunc(html, func(imgTag string) string {
		var parts []string
		if hasImage {
			parts = []string{sectionEnd}
		}
		hasImage = true
		parts = append(parts, []string{imgTag, sectionStart}...)
		return strings.Join(parts, "")
	})

	replaced = strings.Replace(replaced, "<p>", "", 1)
	replaced = strings.Replace(replaced, "</p>", "", 1)
	if hasImage {
		replaced += sectionEnd
	}
	return replaced
}

func (sectionizer *Sectionizer) TemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"sectionize": sectionizer.Sectionize,
	}
}
