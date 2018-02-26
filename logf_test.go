package logf

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLevelOutput(t *testing.T) {
	testCase := []struct {
		l Level
		m string
		s string
	}{
		{l: TRACE, m: "Tracing", s: "[TRACE] Tracing\n"},
		{l: DEBUG, m: "Debugging", s: "[DEBUG] Debugging\n"},
		{l: INFO, m: "Information", s: "[INFO] Information\n"},
		{l: WARN, m: "Warning", s: "[WARN] Warning\n"},
		{l: ERROR, m: "Erroring", s: "[ERROR] Erroring\n"},
		{l: FATAL, m: "Fatal Erroring", s: "[FATAL] Fatal Erroring\n"},
		{l: FATAL + 1, m: "Unknown", s: "[] Unknown\n"},
	}
	for _, tst := range testCase {
		outBuf := new(bytes.Buffer)
		l := New(
			Writer(outBuf),
			Flags(Llevel),
			Prefix(""),
			MinLevel(TRACE),
		)
		l.Output(tst.l, 2, tst.m)
		s := outBuf.String()
		if s != tst.s {
			t.Errorf("Logger.Output(%d, \"%s\")  = \"%v\", want \"%v\".", int(tst.l), tst.m, s, tst.s)
		}
	}
	for _, tst := range testCase {
		outBuf := new(bytes.Buffer)
		std.SetFlags(Llevel)
		std.SetOutput(outBuf)
		Output(tst.l, 3, tst.m)
		s := outBuf.String()
		if s != tst.s {
			t.Errorf("Logger.Output(%d, \"%s\")  = \"%v\", want \"%v\".", int(tst.l), tst.m, s, tst.s)
		}
	}
}

func TestLevelOutput2(t *testing.T) {
	testCase := []struct {
		l Level
		m string
		s string
	}{
		{l: TRACE, m: "Tracing", s: ""},
		{l: DEBUG, m: "Debugging", s: ""},
		{l: INFO, m: "Information", s: ""},
		{l: WARN, m: "Warning", s: "[WARN] Warning\n"},
		{l: ERROR, m: "Erring", s: "[ERROR] Erring\n"},
		{l: FATAL, m: "Fatal Erring", s: "[FATAL] Fatal Erring\n"},
		{l: FATAL + 1, m: "Unknown", s: "[] Unknown\n"},
	}
	for _, tst := range testCase {
		outBuf := new(bytes.Buffer)
		l := New(
			Writer(outBuf),
			Flags(Llevel),
			Prefix(""),
			MinLevel(WARN),
		)
		l.Output(tst.l, 3, tst.m)
		s := outBuf.String()
		if s != tst.s {
			t.Errorf("Logger.Output(%d, \"%s\")  = \"%v\", want \"%v\".", int(tst.l), tst.m, s, tst.s)
		}
	}
}

func TestLevelOutput3(t *testing.T) {
	testCase := []struct {
		l Level
		m string
		s string
	}{
		{l: TRACE, m: "Tracing", s: "Tracing\n"},
		{l: DEBUG, m: "Debugging", s: "Debugging\n"},
		{l: INFO, m: "Information", s: "Information\n"},
		{l: WARN, m: "Warning", s: "Warning\n"},
		{l: ERROR, m: "Erring", s: "Erring\n"},
		{l: FATAL, m: "Fatal Erring", s: "Fatal Erring\n"},
		{l: FATAL + 1, m: "Unknown", s: "Unknown\n"},
	}
	for _, tst := range testCase {
		outBuf := new(bytes.Buffer)
		l := New(
			Writer(outBuf),
			Flags(0),
			Prefix(""),
			MinLevel(TRACE),
		)
		l.Output(tst.l, 3, tst.m)
		s := outBuf.String()
		if s != tst.s {
			t.Errorf("Logger.Output(%d, \"%s\")  = \"%v\", want \"%v\".", int(tst.l), tst.m, s, tst.s)
		}
	}
}

func TestTraceOutput(t *testing.T) {
	m1 := 123
	m2 := "string"
	res := []string{
		"logf_test.go:127: [TRACE] 123 string\n",
		"logf_test.go:129: [TRACE] 123string\n",
		"logf_test.go:131: [TRACE] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			Writer(outBuf),
			Flags(Llevel|Lshortfile),
			Prefix(""),
			MinLevel(TRACE),
		)
		switch i {
		case 0:
			l.Tracef("%v %v", m1, m2)
		case 1:
			l.Trace(m1, m2)
		default:
			l.Traceln(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Trace(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
	res2 := []string{
		"logf_test.go:149: [TRACE] 123 string\n",
		"logf_test.go:151: [TRACE] 123string\n",
		"logf_test.go:153: [TRACE] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		std.SetFlags(Llevel | Lshortfile)
		std.SetOutput(outBuf)
		switch i {
		case 0:
			Tracef("%v %v", m1, m2)
		case 1:
			Trace(m1, m2)
		default:
			Traceln(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Trace(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
}

func TestDebugOutput(t *testing.T) {
	m1 := 123
	m2 := "string"
	res := []string{
		"logf_test.go:180: [DEBUG] 123 string\n",
		"logf_test.go:182: [DEBUG] 123string\n",
		"logf_test.go:184: [DEBUG] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			Writer(outBuf),
			Flags(Llevel|Lshortfile),
			Prefix(""),
			MinLevel(TRACE),
		)
		switch i {
		case 0:
			l.Debugf("%v %v", m1, m2)
		case 1:
			l.Debug(m1, m2)
		default:
			l.Debugln(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Debug(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
	res2 := []string{
		"logf_test.go:202: [DEBUG] 123 string\n",
		"logf_test.go:204: [DEBUG] 123string\n",
		"logf_test.go:206: [DEBUG] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		std.SetFlags(Llevel | Lshortfile)
		std.SetOutput(outBuf)
		switch i {
		case 0:
			Debugf("%v %v", m1, m2)
		case 1:
			Debug(m1, m2)
		default:
			Debugln(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Debug(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
}

func TestPrintOutput(t *testing.T) {
	m1 := 123
	m2 := "string"
	res := []string{
		"logf_test.go:233: [INFO] 123 string\n",
		"logf_test.go:235: [INFO] 123string\n",
		"logf_test.go:237: [INFO] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			Writer(outBuf),
			Flags(Llevel|Lshortfile),
			Prefix(""),
			MinLevel(TRACE),
		)
		switch i {
		case 0:
			l.Printf("%v %v", m1, m2)
		case 1:
			l.Print(m1, m2)
		default:
			l.Println(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Print(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
	res2 := []string{
		"logf_test.go:255: [INFO] 123 string\n",
		"logf_test.go:257: [INFO] 123string\n",
		"logf_test.go:259: [INFO] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		std.SetFlags(Llevel | Lshortfile)
		std.SetOutput(outBuf)
		switch i {
		case 0:
			Printf("%v %v", m1, m2)
		case 1:
			Print(m1, m2)
		default:
			Println(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Print(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
}

func TestWarnOutput(t *testing.T) {
	m1 := 123
	m2 := "string"
	res := []string{
		"logf_test.go:286: [WARN] 123 string\n",
		"logf_test.go:288: [WARN] 123string\n",
		"logf_test.go:290: [WARN] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			Writer(outBuf),
			Flags(Llevel|Lshortfile),
			Prefix(""),
			MinLevel(TRACE),
		)
		switch i {
		case 0:
			l.Warnf("%v %v", m1, m2)
		case 1:
			l.Warn(m1, m2)
		default:
			l.Warnln(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Warn(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
	res2 := []string{
		"logf_test.go:308: [WARN] 123 string\n",
		"logf_test.go:310: [WARN] 123string\n",
		"logf_test.go:312: [WARN] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		std.SetFlags(Llevel | Lshortfile)
		std.SetOutput(outBuf)
		switch i {
		case 0:
			Warnf("%v %v", m1, m2)
		case 1:
			Warn(m1, m2)
		default:
			Warnln(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Warn(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
}

func TestErrorOutput(t *testing.T) {
	m1 := 123
	m2 := "string"
	res := []string{
		"logf_test.go:339: [ERROR] 123 string\n",
		"logf_test.go:341: [ERROR] 123string\n",
		"logf_test.go:343: [ERROR] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			Writer(outBuf),
			Flags(Llevel|Lshortfile),
			Prefix(""),
			MinLevel(TRACE),
		)
		switch i {
		case 0:
			l.Errorf("%v %v", m1, m2)
		case 1:
			l.Error(m1, m2)
		default:
			l.Errorln(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Error(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
	res2 := []string{
		"logf_test.go:361: [ERROR] 123 string\n",
		"logf_test.go:363: [ERROR] 123string\n",
		"logf_test.go:365: [ERROR] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		std.SetFlags(Llevel | Lshortfile)
		std.SetOutput(outBuf)
		switch i {
		case 0:
			Errorf("%v %v", m1, m2)
		case 1:
			Error(m1, m2)
		default:
			Errorln(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Error(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
}

func TestFatalOutput(t *testing.T) {
	m1 := 123
	m2 := "string"
	res := []string{
		"logf_test.go:392: [FATAL] 123 string\n",
		"logf_test.go:394: [FATAL] 123string\n",
		"logf_test.go:396: [FATAL] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			Writer(outBuf),
			Flags(Llevel|Lshortfile),
			Prefix(""),
			MinLevel(TRACE),
		)
		switch i {
		case 0:
			l.Fatalf("%v %v", m1, m2)
		case 1:
			l.Fatal(m1, m2)
		default:
			l.Fatalln(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Fatal(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
	res2 := []string{
		"logf_test.go:414: [FATAL] 123 string\n",
		"logf_test.go:416: [FATAL] 123string\n",
		"logf_test.go:418: [FATAL] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		std.SetFlags(Llevel | Lshortfile)
		std.SetOutput(outBuf)
		switch i {
		case 0:
			Fatalf("%v %v", m1, m2)
		case 1:
			Fatal(m1, m2)
		default:
			Fatalln(m1, m2)
		}
		s := outBuf.String()
		if s != r {
			t.Errorf("Logger.Fatal(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
}

func catchPanicf(f func(string, ...interface{}), format string, v ...interface{}) (err error) {
	err = nil
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic: %v", r)
		}
	}()
	f(format, v...)
	return
}
func catchPanic(f func(...interface{}), v ...interface{}) (err error) {
	err = nil
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic: %v", r)
		}
	}()
	f(v...)
	return
}
func catchPanicln(f func(...interface{}), v ...interface{}) (err error) {
	err = nil
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic: %v", r)
		}
	}()
	f(v...)
	return
}

func TestPanicOutput(t *testing.T) {
	m1 := 123
	m2 := "string"
	res := []string{
		"logf_test.go:434: [FATAL] 123 string\n",
		"logf_test.go:444: [FATAL] 123string\n",
		"logf_test.go:454: [FATAL] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			Writer(outBuf),
			Flags(Llevel|Lshortfile),
			Prefix(""),
			MinLevel(TRACE),
		)
		var err error
		switch i {
		case 0:
			err = catchPanicf(l.Panicf, "%v %v", m1, m2)
		case 1:
			err = catchPanic(l.Panic, m1, m2)
		default:
			err = catchPanicln(l.Panicln, m1, m2)
		}
		s := outBuf.String()
		if err == nil {
			t.Errorf("Logger.Panic(%d, \"%s\")  = nil, want \"%v\".", m1, m2, err)
		}
		if s != r {
			t.Errorf("Logger.Panic(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
	res2 := []string{
		"logf_test.go:503: [FATAL] 123 string\n",
		"logf_test.go:505: [FATAL] 123string\n",
		"logf_test.go:507: [FATAL] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		std.SetFlags(Llevel | Lshortfile)
		std.SetOutput(outBuf)
		var err error
		switch i {
		case 0:
			err = catchPanicf(Panicf, "%v %v", m1, m2)
		case 1:
			err = catchPanic(Panic, m1, m2)
		default:
			err = catchPanicln(Panicln, m1, m2)
		}
		s := outBuf.String()
		if err == nil {
			t.Errorf("Logger.Panic(%d, \"%s\")  = nil, want \"%v\".", m1, m2, err)
		}
		if s != r {
			t.Errorf("Logger.Panic(%d, \"%s\")  = \"%v\", want \"%v\".", m1, m2, s, r)
		}
	}
}
