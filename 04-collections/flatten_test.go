package strings

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFlatten(t *testing.T) {
	Convey("it works", t, func() {
		results := Flatten([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}, {12, 13, 14, 15}, {16, 17, 18, 19, 20}})
		expectedResults := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		So(results, ShouldResemble, expectedResults)
	})
}
