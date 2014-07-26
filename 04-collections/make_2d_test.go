package strings

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMake2D(t *testing.T) {
	Convey("it works", t, func() {
		Convey("3 columns", func() {
			results := Make2D([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 3)
			expectedResults := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}, {13, 14, 15}, {16, 17, 18}, {19, 20, 0}}
			So(results, ShouldResemble, expectedResults)
		})

		Convey("4 columns", func() {
			results := Make2D([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 4)
			expectedResults := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}}
			So(results, ShouldResemble, expectedResults)
		})

		Convey("5 columns", func() {
			results := Make2D([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 5)
			expectedResults := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}}
			So(results, ShouldResemble, expectedResults)
		})

		Convey("6 columns", func() {
			results := Make2D([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 6)
			expectedResults := [][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {13, 14, 15, 16, 17, 18}, {19, 20, 0, 0, 0, 0}}
			So(results, ShouldResemble, expectedResults)
		})
	})
}
