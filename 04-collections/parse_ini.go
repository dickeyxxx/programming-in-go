package strings

import (
	"regexp"
	"strings"
)

var groupRegex = regexp.MustCompile(`\[(\w+)\]`)
var propRegex = regexp.MustCompile(`(\w+) ?= ?(.+)`)

// Given a slice of strings from an .ini file, parses the text into a document.
func ParseIni(ini []string) (doc map[string]map[string]string) {
	doc = map[string]map[string]string{}
	group := map[string]string{}
	doc["General"] = group
	for _, line := range ini {
		line = strings.TrimSpace(line)
		if matches := groupRegex.FindStringSubmatch(line); len(matches) > 0 {
			group = make(map[string]string)
			doc[matches[1]] = group
			continue
		}
		if matches := propRegex.FindStringSubmatch(line); len(matches) > 0 {
			group[matches[1]] = matches[2]
		}
	}
	if len(doc["General"]) == 0 {
		delete(doc, "General")
	}
	return
}
