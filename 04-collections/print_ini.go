package strings

import (
	"fmt"
	"sort"
	"strings"
)

// Given an object representing an ini file, returns a string for saving to a .ini file.
func PrintIni(doc map[string]map[string]string) string {
	lines := []string{}
	var groups []string
	for group, _ := range doc {
		groups = append(groups, group)
	}
	sort.Strings(groups)
	for _, group := range groups {
		lines = append(lines, fmt.Sprintf("[%s]", group))
		lines = append(lines, printProperties(doc[group])...)
	}
	lines = append(lines, "")
	return strings.Join(lines, "\n")
}

func printProperties(properties map[string]string) (lines []string) {
	var names []string
	for name := range properties {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		lines = append(lines, fmt.Sprintf("%s=%s", name, properties[name]))
	}
	return
}
