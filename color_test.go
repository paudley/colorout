// Copyright (C) 2022-2024 by Blackcat Informatics® Inc.
//nolint
package colorout_test

import (
	"testing"

	"github.com/paudley/colorout"
	. "github.com/smartystreets/goconvey/convey"
)

type s1 struct {
	m string
	i int
}
type s2 struct {
	m map[string]s1
}

func TestErrorCreation(t *testing.T) {
	t.Parallel()
	Convey("check output", t, func() {
		Convey("make sure expected keys in nested structures are present", func() {
			t1 := s2{
				m: make(map[string]s1),
			}
			t1.m["foo"] = s1{m: "bar", i: 10}
			out := colorout.SdumpColorSimple(t1)
			So(out, ShouldContainSubstring, "foo")
			So(out, ShouldContainSubstring, "bar")
			So(out, ShouldContainSubstring, "10")
		})
	})
}
