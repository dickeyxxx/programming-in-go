package strings

import (
	"fmt"
	"sort"
	"strings"
)

// Given an object representing an ini file, returns a string for saving to a .ini file.
func PrintIni(doc map[string]map[string]string) string {
	lines := []string{}
	for _, group := range orderedGroupNames(doc) {
		lines = append(lines, fmt.Sprintf("[%s]", group))
		lines = append(lines, printProperties(doc[group])...)
	}
	lines = append(lines, "")
	return strings.Join(lines, "\n")
}

func printProperties(properties map[string]string) (lines []string) {
	for _, name := range orderedPropertyNames(properties) {
		lines = append(lines, fmt.Sprintf("%s=%s", name, properties[name]))
	}
	return
}

func orderedGroupNames(doc map[string]map[string]string) (names []string) {
	for name := range doc {
		names = append(names, name)
	}
	sort.Strings(names)
	return
}

func orderedPropertyNames(properties map[string]string) (names []string) {
	for name := range properties {
		names = append(names, name)
	}
	sort.Strings(names)
	return
}
