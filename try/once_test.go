package try_test

import (
	"errors"
	"testing"

	. "github.com/tenntenn/sync/try"
)

func TestOnce(t *testing.T) {
	expectedErr := errors.New("error")
	cases := map[string]struct {
		once  Once
		ok    bool
		err   error
		panic interface{}
		f     func(o *Once) (bool, error)
	}{
		"do once": {ok: true, f: func(o *Once) (ok bool, err error) {
			o.Do(func() { ok = true })
			return
		}},
		"do twice": {ok: false, f: func(o *Once) (ok bool, err error) {
			o.Do(func() {})
			o.Do(func() { ok = true })
			return
		}},
		"try once": {ok: true, f: func(o *Once) (ok bool, err error) {
			err = o.Try(func() error { ok = true; return nil })
			return
		}},
		"try twice": {ok: false, f: func(o *Once) (ok bool, err error) {
			err = o.Try(func() error { return nil })
			o.Try(func() error { ok = true; return nil })
			return
		}},
		"try error": {ok: true, f: func(o *Once) (ok bool, err error) {
			err = o.Try(func() error { return expectedErr })
			o.Try(func() error { ok = true; return nil })
			return
		}, err: expectedErr},
		"do and try": {ok: false, f: func(o *Once) (ok bool, err error) {
			o.Do(func() {})
			err = o.Try(func() error { ok = true; return nil })
			return
		}},
		"try and do": {ok: false, f: func(o *Once) (ok bool, err error) {
			err = o.Try(func() error { return nil })
			o.Do(func() { ok = true })
			return
		}},
		"try error and do": {ok: true, f: func(o *Once) (ok bool, err error) {
			err = o.Try(func() error { return expectedErr })
			o.Do(func() { ok = true })
			return
		}, err: expectedErr},
		"try with panic": {panic: "panic", f: func(o *Once) (ok bool, err error) {
			err = o.Try(func() error { panic("panic") })
			return
		}},
	}

	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			defer func() {
				if r := recover(); r != tc.panic {
					t.Errorf("want %v got %v", tc.panic, r)
				}
			}()
			ok, err := tc.f(&tc.once)
			if ok != tc.ok {
				t.Errorf("want %v got %v", tc.ok, ok)
			}
			if err != tc.err {
				t.Errorf("want %v got %v", tc.err, err)
			}
		})
	}
}
