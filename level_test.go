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

/* Copyright 2018 Spiegel
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
