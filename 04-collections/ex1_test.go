package strings

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniqueInts(t *testing.T) {
	Convey("it works", t, func() {
		results := UniqueInts([]int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5})
		expectedResults := []int{9, 1, 5, 4, 2, 8, 3, 6, 7}
		So(results, ShouldResemble, expectedResults)
	})
}
