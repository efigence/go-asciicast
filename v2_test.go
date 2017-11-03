package asciicast

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
	"bytes"
	"fmt"
)

func TestV2Header(t *testing.T) {
	_ = time.Now()
	var buf1 bytes.Buffer
	testV2, newErr := NewCastV2(CastMetadata{},&buf1)
	testV2.PushFrame(0.1, []byte("frame1"))
	testV2.PushFrame(0.2, []byte("frame2"))
	outStr := buf1.String()
	Convey("Basic header", t, func() {
		So(newErr, ShouldBeNil)
		So(string(outStr), ShouldContainSubstring, `"version":2`)
		So(string(outStr), ShouldContainSubstring, `"width"`)
		So(string(outStr), ShouldContainSubstring, `"height"`)
		So(string(outStr), ShouldNotContainSubstring, `"env"`)
		So(string(outStr), ShouldNotContainSubstring, `"timestamp"`)
		So(string(outStr), ShouldNotContainSubstring, `"duration"`)

	})
	Convey("Frame test", t, func() {
		So(string(outStr), ShouldContainSubstring, `[0.1,"o","frame1"]`)
		So(string(outStr), ShouldContainSubstring, `[0.2,"o","frame2"]`)

	})
	fmt.Println(buf1.String())
	ts, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", "Mon Jan 2 15:04:05 -0700 MST 2006")
	var buf2 bytes.Buffer
	testV2, newErr = NewCastV2(CastMetadata{
		Width:     80,
		Height:    24,
		Title:     "Terminal test",
		Timestamp: ts,
		Duration:  time.Minute,
		Command:   "/bin/bash",
	},&buf2)
	testV2.PushFrame(0.3, []byte("frame3"))
	testV2.PushFrame(0.4, []byte("frame4"))

	outStr = buf2.String()
	fmt.Println(buf2.String())
	Convey("Full header", t, func() {
		So(newErr, ShouldBeNil)
		So(string(outStr), ShouldContainSubstring, `"version":2`)
		So(string(outStr), ShouldContainSubstring, `"timestamp":1136239445`)
		So(string(outStr), ShouldContainSubstring, `"Terminal test"`)
		So(string(outStr), ShouldContainSubstring, `"width":80`)
		So(string(outStr), ShouldContainSubstring, `"height":24`)
		So(string(outStr), ShouldContainSubstring, `"/bin/bash`)
		So(string(outStr), ShouldContainSubstring, `"duration":60`)
	})

}
