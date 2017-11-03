package asciicast

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"encoding/json"
	"time"
//	"fmt"
)

func TestV2Header(t *testing.T) {
	_ = time.Now()
	testV2, newErr := NewCastV2 (CastMetadata{
	})

	intB, jsonErr := json.Marshal(testV2)
	Convey("Basic header", t, func() {
		So(newErr,ShouldBeNil)
		So(jsonErr,ShouldBeNil)
		So(string(intB), ShouldContainSubstring, `"version":2`)
		So(string(intB), ShouldContainSubstring, `"width"`)
		So(string(intB), ShouldContainSubstring, `"height"`)
		So(string(intB), ShouldNotContainSubstring, `"env"`)
		So(string(intB), ShouldNotContainSubstring, `"timestamp"`)
		So(string(intB), ShouldNotContainSubstring, `"duration"`)

	})
	ts, _ :=  time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	testV2, newErr = NewCastV2 (CastMetadata{
		Width: 80,
		Height: 24,
		Title: "Terminal test",
		Timestamp: ts,
		Duration: time.Minute,
		Command: "/bin/bash",
	})
	intB, jsonErr = json.Marshal(testV2)
	Convey("Full header", t, func() {
		So(newErr,ShouldBeNil)
		So(jsonErr,ShouldBeNil)
		So(string(intB), ShouldContainSubstring, `"version":2`)
		So(string(intB), ShouldContainSubstring, `"timestamp":1136239445`)
		So(string(intB), ShouldContainSubstring, `"Terminal test"`)
		So(string(intB), ShouldContainSubstring, `"width":80`)
		So(string(intB), ShouldContainSubstring, `"height":24`)
		So(string(intB), ShouldContainSubstring, `"/bin/bash`)
		So(string(intB), ShouldContainSubstring, `"duration":60`)
	})


}
