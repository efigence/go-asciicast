package asciicast

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestUtil(t *testing.T) {
	var zeroTime time.Time
	unixTime := time.Unix(0,0)
	ts0 := JSONTimestamp(zeroTime)
	tsUnix := JSONTimestamp(unixTime)
	// just a sanity check if the override method still functions as time.Time correctly
	Convey("JSONTimestamp", t, func() {
		So(ts0.String(), ShouldContainSubstring, "0001-01-01 00:00:00 +0000" )
		So(tsUnix.String(), ShouldContainSubstring, "1970-01-01" )
		So(ts0.Time().Unix(), ShouldEqual , -62135596800)
		So(tsUnix.Time().Unix(), ShouldEqual , 0)
	})
	var jsTSBad JSONTimestamp
	var jsTSGood JSONTimestamp
	jsDecodeErr := jsTSBad.UnmarshalJSON([]byte(`let them eat cake`))
	jsDecodeOk := jsTSGood.UnmarshalJSON([]byte(`12345678`))
	Convey("JSON Timestamp Decode",t,func() {
		So(jsDecodeErr,ShouldNotBeNil)
		So(jsDecodeOk,ShouldBeNil)
		So(jsTSBad.String(),ShouldEqual,"0001-01-01 00:00:00 +0000 UTC")
		So(jsTSGood.Time().Unix(),ShouldEqual,12345678)
	})



}
