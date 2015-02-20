package command

import (
	"strings"
	"testing"
)

func TestTypeString(t *testing.T) {
	r := &Registry{}

	called := false
	r.Register("testing %", func(caller, arg string) {
		if caller != "gopher" || arg != "hello" {
			t.FailNow()
		}
		called = true
	})

	checkError(t, r.Execute("gopher", "testing hello"))
	if !called {
		t.FailNow()
	}
}

func TestTypeLimit(t *testing.T) {
	r := &Registry{}

	r.Register("testing %3", func(caller, arg string) {
		t.FailNow()
	})

	err := r.Execute("gopher", "testing hello")
	if err == nil || !strings.Contains(err.Error(), "too long") {
		t.FailNow()
	}
}

func TestTypeSub(t *testing.T) {
	r := &Registry{}

	called := 0
	r.Register("testing % a", func(caller, arg string) {
		if caller != "gopher" || arg != "hello" {
			t.FailNow()
		}
		if called != 0 {
			t.FailNow()
		}
		called++
	})
	r.Register("testing % b", func(caller, arg string) {
		if caller != "gopher" || arg != "hello" {
			t.FailNow()
		}
		if called != 1 {
			t.FailNow()
		}
		called++
	})

	checkError(t, r.Execute("gopher", "testing hello a"))
	checkError(t, r.Execute("gopher", "testing hello b"))
	if called != 2 {
		t.FailNow()
	}
}
