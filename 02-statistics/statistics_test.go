package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStatistics(t *testing.T) {
	Convey("Sum", t, func() {
		result := Sum([]float64{2, 3, 4, 7, 9})
		So(result, ShouldAlmostEqual, 25)
	})

	Convey("Mean", t, func() {
		result := Mean([]float64{2, 3, 4, 7, 9})
		So(result, ShouldAlmostEqual, 5)
	})

	Convey("Median", t, func() {
		Convey("With even amount of numbers", func() {
			result := Median([]float64{2, 4, 5, 9})
			So(result, ShouldAlmostEqual, 4.5)
		})

		Convey("With odd amount of numbers", func() {
			result := Median([]float64{2, 4, 5, 8, 9})
			So(result, ShouldAlmostEqual, 5)
		})
	})

	Convey("StdDev", t, func() {
		result := StdDev([]float64{2, 3, 4, 7, 9})
		So(result, ShouldAlmostEqual, 2.9154759474)
	})

	Convey("getStats", t, func() {
		numbers := []float64{2, 3, 4, 7, 9}
		result := getStats(numbers)
		expectedResult := statistics{
			numbers: numbers,
			mean:    5,
			median:  4,
			stdDev:  2.9154759474226504,
		}
		So(result, ShouldResemble, expectedResult)
	})

	Convey("formatStats", t, func() {
		result := formatStats(statistics{
			numbers: []float64{2, 3, 4, 7, 9},
			mean:    5,
			median:  4,
			stdDev:  2.9154759474226504,
		})
		So(result, ShouldEqual, `<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>[2 3 4 7 9]</td></tr>
<tr><td>Count</td><td>5</td></tr>
<tr><td>Mean</td><td>5.000000</td></tr>
<tr><td>Median</td><td>4.000000</td></tr>
<tr><td>Std dev.</td><td>2.915476</td></tr>
</table>`)
	})
}
