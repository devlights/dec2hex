package main

import (
	"testing"
)

type (
	doTestInfo struct {
		in  string
		out string
	}
)

var (
	doTestValues = []doTestInfo{
		doTestInfo{"255", "0xFF"},
		doTestInfo{"10", "0xA"},
		doTestInfo{"0", "0x0"},
		doTestInfo{"", ""},
	}
)

func TestDo(t *testing.T) {
	for _, d := range doTestValues {
		// Arrange
		// Act
		got, err := Do(d.in)

		// Assert
		if err != nil {
			t.Error(err)
		}

		if d.out != got {
			t.Errorf("want: %v, but got:%v\n", d.out, got)
		}
	}
}
