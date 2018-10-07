# [logf] - Simple leveled logging package by Golang

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/logf.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/logf)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/logf/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/logf.svg)](https://github.com/spiegel-im-spiegel/logf/releases/latest)

## Declare [logf] module

See [go.mod](https://github.com/spiegel-im-spiegel/logf/blob/master/go.mod) file. 

## Usage of [logf] package

```go
package main

import (
	"os"

	"github.com/spiegel-im-spiegel/logf"
)

func main() {
	logf.SetOutput(os.Stdout)
	for i := 0; i < 6; i++ {
		logf.SetMinLevel(logf.TRACE + logf.Level(i))
		logf.Tracef("Traceing: No. %d\n", i+1)
		logf.Debugf("Debugging: No. %d\n", i+1)
		logf.Printf("Information: No. %d\n", i+1)
		logf.Warnf("Warning: No. %d\n", i+1)
		logf.Errorf("Erroring: No. %d\n", i+1)
		logf.Fatalf("Fatal Erroring: No. %d\n", i+1)
	}
}
```

Output:

```
2009/11/10 23:00:00 [TRACE] Traceing: No. 1
2009/11/10 23:00:00 [DEBUG] Debugging: No. 1
2009/11/10 23:00:00 [INFO] Information: No. 1
2009/11/10 23:00:00 [WARN] Warning: No. 1
2009/11/10 23:00:00 [ERROR] Erroring: No. 1
2009/11/10 23:00:00 [FATAL] Fatal Erroring: No. 1
2009/11/10 23:00:00 [DEBUG] Debugging: No. 2
2009/11/10 23:00:00 [INFO] Information: No. 2
2009/11/10 23:00:00 [WARN] Warning: No. 2
2009/11/10 23:00:00 [ERROR] Erroring: No. 2
2009/11/10 23:00:00 [FATAL] Fatal Erroring: No. 2
2009/11/10 23:00:00 [INFO] Information: No. 3
2009/11/10 23:00:00 [WARN] Warning: No. 3
2009/11/10 23:00:00 [ERROR] Erroring: No. 3
2009/11/10 23:00:00 [FATAL] Fatal Erroring: No. 3
2009/11/10 23:00:00 [WARN] Warning: No. 4
2009/11/10 23:00:00 [ERROR] Erroring: No. 4
2009/11/10 23:00:00 [FATAL] Fatal Erroring: No. 4
2009/11/10 23:00:00 [ERROR] Erroring: No. 5
2009/11/10 23:00:00 [FATAL] Fatal Erroring: No. 5
2009/11/10 23:00:00 [FATAL] Fatal Erroring: No. 6
```

### Create logger instance

```go
package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spiegel-im-spiegel/logf"
)

func main() {
	rl, err := rotatelogs.New("./log.%Y%m%d%H%M.txt")
	if err != nil {
		logf.Fatal(err)
		return
	}
	logger := logf.New(
		logf.WithFlags(logf.LstdFlags|logf.Lshortfile),
		logf.WithPrefix("[Sample] "),
		logf.WithWriter(rl),
		logf.WithMinLevel(logf.INFO),
	)
	logger.Print("Information")
	//Output:
	//[Sample] 2009/11/10 23:00:00 sample.go:20: [INFO] Information
}
```

## Reference

- [lestrrat-go/file-rotatelogs: Port of perl5 File::RotateLogs to Go](https://github.com/lestrrat-go/file-rotatelogs)
- [rs/zerolog: Zero Allocation JSON Logger](https://github.com/rs/zerolog) : my favorite logger!

[logf]: https://github.com/spiegel-im-spiegel/logf "spiegel-im-spiegel/logf: Simple logging package by Golang"
