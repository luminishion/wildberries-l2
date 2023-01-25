package main

import (
	"testing"
)

func assert(t *testing.T, this string, want string) {
	got := Unpack(this)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestUnpack(t *testing.T) {
	assert(t, "a4bc2d5e", "aaaabccddddde")
	assert(t, "abcd", "abcd")
	assert(t, "45", "")
	assert(t, "", "")
	assert(t, `qwe\4\5`, `qwe45`)
	assert(t, `qwe\45`, `qwe44444`)
	assert(t, `qwe\\5`, `qwe\\\\\`)
}
