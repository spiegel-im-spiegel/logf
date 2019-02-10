package logf

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLevelOutput(t *testing.T) {
	testCase := []struct {
		l  Level
		m  string
		s  string
		s2 string
	}{
		{l: TRACE, m: "Tracing", s: "[TRACE] Tracing\nTracing\n"},
		{l: DEBUG, m: "Debugging", s: "[DEBUG] Debugging\nDebugging\n"},
		{l: INFO, m: "Information", s: "[INFO] Information\nInformation\n"},
		{l: WARN, m: "Warning", s: "[WARN] Warning\nWarning\n"},
		{l: ERROR, m: "Erroring", s: "[ERROR] Erroring\nErroring\n"},
		{l: FATAL, m: "Fatal Erroring", s: "[FATAL] Fatal Erroring\nFatal Erroring\n"},
		{l: FATAL + 1, m: "Unknown", s: "[] Unknown\nUnknown\n"},
	}
	for _, tst := range testCase {
		outBuf := new(bytes.Buffer)
		l := New(
			WithWriter(outBuf),
			WithFlags(Llevel),
			WithPrefix(""),
			WithMinLevel(TRACE),
		)
		if l.MinLevel() != TRACE {
			t.Errorf("Logger.MinLevel()  = %v, want %v.", l.MinLevel(), TRACE)
		}
		if err := l.Output(tst.l, 2, tst.m); err != nil {
			t.Errorf("Result of Logger.Output()  = %v, want nil.", err)
		}
		if err := l.GetLogger().Output(1, tst.m); err != nil {
			t.Errorf("Result of Logger.lg.Output()  = %v, want nil.", err)
		}
		s := outBuf.String()
		if s != tst.s {
			t.Errorf("Logger.Output(%d, \"%s\")  = \"%v\", want \"%v\".", int(tst.l), tst.m, s, tst.s)
		}
	}
	for _, tst := range testCase {
		outBuf := new(bytes.Buffer)
		SetFlags(Llevel)
		SetOutput(outBuf)
		if MinLevel() != TRACE {
			t.Errorf("Logger.MinLevel()  = %v, want %v.", MinLevel(), TRACE)
		}
		if err := Output(tst.l, 3, tst.m); err != nil {
			t.Errorf("Result of Logger.Output()  = %v, want nil.", err)
		}
		if err := GetLogger().Output(2, tst.m); err != nil {
			t.Errorf("Result of Logger.lg.Output()  = %v, want nil.", err)
		}
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
			WithWriter(outBuf),
			WithFlags(Llevel),
			WithPrefix(""),
			WithMinLevel(WARN),
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
			WithWriter(outBuf),
			WithFlags(0),
			WithPrefix(""),
			WithMinLevel(TRACE),
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
		"logf_test.go:144: [TRACE] 123 string\n",
		"logf_test.go:146: [TRACE] 123string\n",
		"logf_test.go:148: [TRACE] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			WithWriter(outBuf),
			WithFlags(Llevel|Lshortfile),
			WithPrefix(""),
			WithMinLevel(TRACE),
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
		"logf_test.go:166: [TRACE] 123 string\n",
		"logf_test.go:168: [TRACE] 123string\n",
		"logf_test.go:170: [TRACE] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		SetFlags(Llevel | Lshortfile)
		SetOutput(outBuf)
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
		"logf_test.go:197: [DEBUG] 123 string\n",
		"logf_test.go:199: [DEBUG] 123string\n",
		"logf_test.go:201: [DEBUG] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			WithWriter(outBuf),
			WithFlags(Llevel|Lshortfile),
			WithPrefix(""),
			WithMinLevel(TRACE),
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
		"logf_test.go:219: [DEBUG] 123 string\n",
		"logf_test.go:221: [DEBUG] 123string\n",
		"logf_test.go:223: [DEBUG] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		SetFlags(Llevel | Lshortfile)
		SetOutput(outBuf)
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
		"logf_test.go:250: [INFO] 123 string\n",
		"logf_test.go:252: [INFO] 123string\n",
		"logf_test.go:254: [INFO] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			WithWriter(outBuf),
			WithFlags(Llevel|Lshortfile),
			WithPrefix(""),
			WithMinLevel(TRACE),
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
		"logf_test.go:272: [INFO] 123 string\n",
		"logf_test.go:274: [INFO] 123string\n",
		"logf_test.go:276: [INFO] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		SetFlags(Llevel | Lshortfile)
		SetOutput(outBuf)
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
		"logf_test.go:303: [WARN] 123 string\n",
		"logf_test.go:305: [WARN] 123string\n",
		"logf_test.go:307: [WARN] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			WithWriter(outBuf),
			WithFlags(Llevel|Lshortfile),
			WithPrefix(""),
			WithMinLevel(TRACE),
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
		"logf_test.go:325: [WARN] 123 string\n",
		"logf_test.go:327: [WARN] 123string\n",
		"logf_test.go:329: [WARN] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		SetFlags(Llevel | Lshortfile)
		SetOutput(outBuf)
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
		"logf_test.go:356: [ERROR] 123 string\n",
		"logf_test.go:358: [ERROR] 123string\n",
		"logf_test.go:360: [ERROR] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			WithWriter(outBuf),
			WithFlags(Llevel|Lshortfile),
			WithPrefix(""),
			WithMinLevel(TRACE),
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
		"logf_test.go:378: [ERROR] 123 string\n",
		"logf_test.go:380: [ERROR] 123string\n",
		"logf_test.go:382: [ERROR] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		SetFlags(Llevel | Lshortfile)
		SetOutput(outBuf)
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
		"logf_test.go:409: [FATAL] 123 string\n",
		"logf_test.go:411: [FATAL] 123string\n",
		"logf_test.go:413: [FATAL] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			WithWriter(outBuf),
			WithFlags(Llevel|Lshortfile),
			WithPrefix(""),
			WithMinLevel(TRACE),
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
		"logf_test.go:431: [FATAL] 123 string\n",
		"logf_test.go:433: [FATAL] 123string\n",
		"logf_test.go:435: [FATAL] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		SetFlags(Llevel | Lshortfile)
		SetOutput(outBuf)
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
		"logf_test.go:451: [FATAL] 123 string\n",
		"logf_test.go:461: [FATAL] 123string\n",
		"logf_test.go:471: [FATAL] 123 string\n",
	}
	for i, r := range res {
		outBuf := new(bytes.Buffer)
		l := New(
			WithWriter(outBuf),
			WithFlags(Llevel|Lshortfile),
			WithPrefix(""),
			WithMinLevel(TRACE),
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
		"[TEST] logf_test.go:522: [FATAL] 123 string\n",
		"[TEST] logf_test.go:524: [FATAL] 123string\n",
		"[TEST] logf_test.go:526: [FATAL] 123 string\n",
	}
	for i, r := range res2 {
		outBuf := new(bytes.Buffer)
		SetFlags(Llevel | Lshortfile)
		SetOutput(outBuf)
		SetPrefix("[TEST] ")
		SetMinLevel(FATAL)
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

/* Copyright 2018,2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
