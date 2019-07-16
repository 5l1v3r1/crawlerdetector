package crawlerdetector

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPiwikMobilesList(t *testing.T) {
	response := PiwikMobilesList()
	Convey("Subject: Test PiwikMobilesList response\n", t, func() {
		Convey("Should NOT be Equal to ()", func() {
			So(len(response), ShouldNotEqual, 2)
		})
		Convey("Should Contain ELEMENT", func() {
			So(response, ShouldContainSubstring, "ELEMENT")
		})
		Convey("Should Contain HTC", func() {
			So(response, ShouldContainSubstring, "HTC")
		})
		Convey("Should Contain Nokia", func() {
			So(response, ShouldContainSubstring, "Nokia")
		})
		Convey("Should Contain BB10", func() {
			So(response, ShouldContainSubstring, "BB10")
		})
		Convey("Should Contain Apple", func() {
			So(response, ShouldContainSubstring, "Apple")
		})
		Convey("Should Contain SAMSUNG", func() {
			So(response, ShouldContainSubstring, "SAMSUNG")
		})
		Convey("Should Contain LG", func() {
			So(response, ShouldContainSubstring, "LG")
		})
		Convey("Should Contain Alcatel", func() {
			So(response, ShouldContainSubstring, "Alcatel")
		})
	})
}

//BenchmarkHandleRequest get a APIGatewayProxyResponse from a APIGatewayProxyRequest
func BenchmarkHandleRequest(b *testing.B) {

	// b.ReportAllocs()
	// b.ResetTimer()
	// for _, test := range testsOK {
	// 	b.Run(test, func(bb *testing.B) {
	// 		bb.ReportAllocs()
	// 		bb.ResetTimer()
	// 		for n := 0; n < bb.N; n++ {
	// 			// handleRequest(context.TODO(), event)
	// 		}
	// 	},
	// 	)
	// }
}
