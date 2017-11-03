package asciicast

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
	"encoding/json"
)

func TestV1Header(t *testing.T) {
	_ = time.Now()
	testV1, newErr := NewCastV1 (CastMetadata{
	})

	intB, jsonErr := json.Marshal(testV1)
	Convey("Basic header", t, func() {
		So(newErr,ShouldBeNil)
		So(jsonErr,ShouldBeNil)
		So(string(intB), ShouldContainSubstring, `"version":1`)
		So(string(intB), ShouldContainSubstring, `"width"`)
		So(string(intB), ShouldContainSubstring, `"height"`)
	})
		ts, _ :=  time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	testV1, newErr = NewCastV1 (CastMetadata{
		Width: 80,
		Height: 24,
		Title: "Terminal test",
		Timestamp: ts,
		Duration: time.Minute,
		Command: "/bin/bash",
	})
	intB, jsonErr = json.Marshal(testV1)
	Convey("Full header", t, func() {
		So(newErr,ShouldBeNil)
		So(jsonErr,ShouldBeNil)
		So(string(intB), ShouldContainSubstring, `"version":1`)
		So(string(intB), ShouldContainSubstring, `"duration":60`)
		So(string(intB), ShouldContainSubstring, `"Terminal test"`)
		So(string(intB), ShouldContainSubstring, `"width":80`)
		So(string(intB), ShouldContainSubstring, `"height":24`)
		So(string(intB), ShouldContainSubstring, `"/bin/bash`)
		// that is not supported in v1
		So(string(intB), ShouldNotContainSubstring, `"timestamp"`)
	})
}
