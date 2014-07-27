package procedural

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsPalindrom(t *testing.T) {
	Convey("with a word that is not a palindrome", t, func() {
		word := "pulrup"

		Convey("returns false", func() {
			result := IsPalindrome(word)
			So(result, ShouldBeFalse)
		})
	})

	Convey("with a word that is a palindrome", t, func() {
		word := "pullup"

		Convey("returns true", func() {
			result := IsPalindrome(word)
			So(result, ShouldBeTrue)
		})
	})

	Convey("with a unicode word that is a palindrome", t, func() {
		word := "mørrrøm"

		Convey("returns true", func() {
			result := IsPalindrome(word)
			So(result, ShouldBeTrue)
		})
	})
}
