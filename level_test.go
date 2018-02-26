package logf

import "testing"

func TestLevelString(t *testing.T) {
	testCase := []struct {
		l Level
		s string
	}{
		{l: TRACE, s: "TRACE"},
		{l: DEBUG, s: "DEBUG"},
		{l: INFO, s: "INFO"},
		{l: WARN, s: "WARN"},
		{l: ERROR, s: "ERROR"},
		{l: FATAL, s: "FATAL"},
		{l: FATAL + 1, s: ""},
	}
	for _, tst := range testCase {
		if tst.l.String() != tst.s {
			t.Errorf("Level(%d)  = \"%v\", want \"%v\".", int(tst.l), tst.l, tst.s)
		}
	}
}
