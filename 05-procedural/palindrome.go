package procedural

// returns true if palindrome, false otherwise
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}
