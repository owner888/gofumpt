// Copyright (c) 2021, Daniel Martí <mvdan@>
// See LICENSE for licensing information

package format_test

import (
    "testing"

    qt "github.com/frankban/quicktest"

    "owner888/gofumpt/format"
)

func TestSourceIncludesSimplify(t *testing.T) {
	t.Parallel()

	in := []byte(`
package p

var ()

func f() {
	for _ = range v {
	}
}
`[1:])
	want := []byte(`
package p

func f() {
	for range v {
	}
}
`[1:])
	got, err := format.Source(in, format.Options{})
	qt.Assert(t, err, qt.IsNil)
	qt.Assert(t, string(got), qt.Equals, string(want))
}
