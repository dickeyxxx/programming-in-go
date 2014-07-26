package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWeb(t *testing.T) {
	Convey("GET homePage", t, func() {
		req, _ := http.NewRequest("GET", "http://localhost:3000/", nil)
		w := httptest.NewRecorder()
		homePage(w, req)

		Convey("it renders the html", func() {
			fmt.Println(w.Body.String())
			So(w.Body.String(), ShouldStartWith, "<!DOCTYPE HTML>")
		})

		Convey("it has no results", func() {
			fmt.Println(w.Body.String())
			So(w.Body.String(), ShouldNotContainSubstring, "Results")
		})

		Convey("it has status code 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})

	Convey("POST homePage", t, func() {
		body := strings.NewReader("numbers=1+2+3+4")
		req, _ := http.NewRequest("POST", "http://localhost:3000/", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		w := httptest.NewRecorder()
		homePage(w, req)

		Convey("it renders the html", func() {
			fmt.Println(w.Body.String())
			So(w.Body.String(), ShouldStartWith, "<!DOCTYPE HTML>")
		})

		Convey("it has results", func() {
			fmt.Println(w.Body.String())
			So(w.Body.String(), ShouldContainSubstring, "Results")
		})

		Convey("it has status code 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
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
