package asciicast

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
	"fmt"
)

func TestV1Header(t *testing.T) {
	_ = time.Now()
	testV1, newErr := NewCastV1(CastMetadata{})
	testV1.PushFrame(0.1, []byte("frame1"))
	testV1.PushFrame(0.2, []byte("frame2"))

	outStr, jsonErr := json.Marshal(testV1)
	Convey("Basic header", t, func() {
		So(newErr, ShouldBeNil)
		So(jsonErr, ShouldBeNil)
		So(string(outStr), ShouldContainSubstring, `"version":1`)
		So(string(outStr), ShouldContainSubstring, `"width"`)
		So(string(outStr), ShouldContainSubstring, `"height"`)
	})
	Convey("Frame test", t, func() {
		So(string(outStr), ShouldContainSubstring, `[0.1,"frame1"]`)
		So(string(outStr), ShouldContainSubstring, `[0.2,"frame2"]`)

	})
	fmt.Println(string(outStr))
	ts, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	testV1, newErr = NewCastV1(CastMetadata{
		Width:     80,
		Height:    24,
		Title:     "Terminal test",
		Timestamp: ts,
		Duration:  time.Minute,
		Command:   "/bin/bash",
		Env: map[string]string{
			"PATH": "/bin:/usr/bin",
		},
	})

	outStr, jsonErr = json.Marshal(testV1)
	Convey("Full header", t, func() {
		So(newErr, ShouldBeNil)
		So(jsonErr, ShouldBeNil)
		So(string(outStr), ShouldContainSubstring, `"version":1`)
		So(string(outStr), ShouldContainSubstring, `"duration":60`)
		So(string(outStr), ShouldContainSubstring, `"Terminal test"`)
		So(string(outStr), ShouldContainSubstring, `"width":80`)
		So(string(outStr), ShouldContainSubstring, `"height":24`)
		So(string(outStr), ShouldContainSubstring, `"/bin/bash`)
		// that is not supported in v1
		So(string(outStr), ShouldNotContainSubstring, `"timestamp"`)
	})
	fmt.Println(string(outStr))

}
