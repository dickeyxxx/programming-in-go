package procedural

import "unicode/utf8"

// returns true if palindrome, false otherwise
func IsPalindrome(s string) bool {
	runes := runes(s)
	for i := range runes {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}

func runes(s string) []rune {
	t := make([]rune, utf8.RuneCountInString(s))
	i := 0
	for _, r := range s {
		t[i] = r
		i++
	}
	return t
}
