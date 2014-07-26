package procedural

import "unicode/utf8"

func IsPalindrome(s string) bool {
	if utf8.RuneCountInString(s) <= 1 {
		return true
	}
	first, sizeOfFirst := utf8.DecodeRuneInString(s)
	last, sizeOfLast := utf8.DecodeLastRuneInString(s)
	if first != last {
		return false
	}
	return IsPalindrome(s[sizeOfFirst : len(s)-sizeOfLast])
}
