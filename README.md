# [logf] - Simple logging package by Golang


## Install 

```
$ go get -u github.com/spiegel-im-spiegel/logf
```

## Usage

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

[logf]: https://github.com/spiegel-im-spiegel/logf "spiegel-im-spiegel/logf: Simple logging package by Golang"
