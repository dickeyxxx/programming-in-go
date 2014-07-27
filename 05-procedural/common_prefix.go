package procedural

import "strings"

func CommonPrefix(strings []string) string {
	if len(strings) == 0 {
		return ""
	}
	result := []rune(strings[0])
	for _, s := range strings {
		for y, c := range []rune(s) {
			if y < len(result) && result[y] != c {
				result = result[:y]
			}
		}
	}
	return string(result)
}

func CommonPathPrefix(input []string) string {
	if len(input) == 0 {
		return ""
	}
	paths := make([][]string, len(input))
	for i, s := range input {
		paths[i] = strings.Split(s, "/")
	}
	result := paths[0]
	for _, path := range paths {
		for i := range path {
			if i < len(result) && path[i] != result[i] {
				result = result[:i]
			}
		}
	}
	return strings.Join(result, "/")
}
