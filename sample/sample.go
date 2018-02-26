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
