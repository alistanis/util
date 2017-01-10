package util

import (
	_ "harvard/shared/log"
	"testing"

	"reflect"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConvertStringSliceToPointerSlice(t *testing.T) {
	Convey("We can pass a slice of strings and receive a slice of pointers to strings, because Amazon", t, func() {
		s := []string{"Why", "does", "Amazon", "use", "pointers", "for", "absolutely", "everything", "?"}
		ptrs := StringSliceToPointers(s)
		So(ptrs, ShouldNotBeNil)

		So(*ptrs[0], ShouldEqual, "Why")
		So(*ptrs[8], ShouldEqual, "?")

		s2 := PointerSliceToStrings(ptrs)
		So(reflect.DeepEqual(s, s2), ShouldBeTrue)
	})
}

