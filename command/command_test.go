package command

import (
	"testing"
)

func TestBasic(t *testing.T) {
	r := Registry{}

	called := false
	r.Register("test", func(caller string) {
		called = true
	})
	checkError(t, r.Execute("gopher", "test"))
	if !called {
		t.FailNow()
	}
}

func TestSubCommands(t *testing.T) {
	r := Registry{}

	called := 0
	r.Register("test a", func(caller string) {
		called++
	})
	r.Register("test b", func(caller string) {
		called++
	})
	r.Register("test c a", func(caller string) {
		called++
	})
	r.Register("test c b", func(caller string) {
		called++
	})
	r.Register("test d e f g", func(caller string) {
		called++
	})

	checkError(t, r.Execute("gopher", "test a"))
	checkError(t, r.Execute("gopher", "test b"))
	checkError(t, r.Execute("gopher", "test c a"))
	checkError(t, r.Execute("gopher", "test c b"))
	checkError(t, r.Execute("gopher", "test d e f g"))

	if called != 5 {
		t.FailNow()
	}
}

func TestNonFunction(t *testing.T) {
	shouldPanic(t, func() {
		r := Registry{}
		r.Register("test a", "")
	})
}

func TestInvalidDesc(t *testing.T) {
	shouldPanic(t, func() {
		r := Registry{}
		r.Register("", func(caller string) {

		})
	})
}

func TestDoubleAdd(t *testing.T) {
	shouldPanic(t, func() {
		r := Registry{}
		r.Register("test", func(caller string) {

		})
		r.Register("test", func(caller string) {

		})
	})
}

func TestExtraParams(t *testing.T) {
	r := Registry{
		ExtraParameters: 2,
	}
	r.Register("extra", func(caller, a, b string) {
		if caller != "gopher" ||
			a != "a" || b != "b" {
			t.FailNow()
		}
	})
	checkError(t, r.Execute("gopher", "extra", "a", "b"))
}

func TestExtraParamsFail(t *testing.T) {
	shouldPanic(t, func() {
		r := Registry{
			ExtraParameters: 2,
		}
		r.Register("extra", func(caller, a, b string) {
			t.FailNow()
		})
		r.Execute("gopher", "extra", "a", "b", "c")
	})
}

func TestEmpty(t *testing.T) {
	r := Registry{}
	err := r.Execute("gopher", "test")
	if err != ErrCommandNotFound {
		t.FailNow()
	}
}

func TestMissing(t *testing.T) {
	r := Registry{}
	r.Register("hello world", func(caller string) {
		t.FailNow()
	})
	err := r.Execute("gopher", "test")
	if err != ErrCommandNotFound {
		t.FailNow()
	}
}

func TestMissing2(t *testing.T) {
	r := Registry{}
	r.Register("hello world", func(caller string) {
		t.FailNow()
	})
	err := r.Execute("gopher", "hello")
	if err != ErrCommandNotFound {
		t.FailNow()
	}
}

func TestQuoted(t *testing.T) {
	r := Registry{}
	called := false
	r.Register("hello world", func(caller string) {
		called = true
	})
	checkError(t, r.Execute("gopher", "hello \"world\""))
	if !called {
		t.FailNow()
	}
}

func TestCommandPanic(t *testing.T) {
	r := Registry{}
	r.Register("hello world", func(caller string) {
		panic("test panic")
	})
	err := r.Execute("gopher", "hello \"world\"")
	if err == nil || err.Error() != "test panic" {
		t.FailNow()
	}
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func shouldPanic(t *testing.T, f func()) {
	defer func() {
		if err := recover(); err == nil {
			t.FailNow()
		}
	}()
	f()
}
