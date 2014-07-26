package strings

import (
	"regexp"
	"strings"
)

var groupRegex = regexp.MustCompile(`\[(\w+)\]`)
var propRegex = regexp.MustCompile(`(\w+) ?= ?(.+)`)

func ParseIni(ini []string) map[string]map[string]string {
	doc := map[string]map[string]string{}
	group := map[string]string{}
	doc["General"] = group
	for _, line := range ini {
		line = strings.TrimSpace(line)
		matches := groupRegex.FindStringSubmatch(line)
		if len(matches) > 0 {
			group = make(map[string]string)
			doc[matches[1]] = group
			continue
		}
		matches = propRegex.FindStringSubmatch(line)
		if len(matches) > 0 {
			group[matches[1]] = matches[2]
		}
	}
	if len(doc["General"]) == 0 {
		delete(doc, "General")
	}
	return doc
}
