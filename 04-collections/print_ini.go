package strings

import (
	"fmt"
	"sort"
)

// Given an object representing an ini file, returns a string for saving to a .ini file.
func PrintIni(doc map[string]map[string]string) (ini string) {
	var groups []string
	for group, _ := range doc {
		groups = append(groups, group)
	}
	sort.Strings(groups)
	for _, group := range groups {
		ini += fmt.Sprintf("[%s]\n", group)
		ini += printProperties(doc[group])
	}
	return
}

func printProperties(properties map[string]string) (ini string) {
	var names []string
	for name := range properties {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		ini += fmt.Sprintf("%s=%s\n", name, properties[name])
	}
	return
}
