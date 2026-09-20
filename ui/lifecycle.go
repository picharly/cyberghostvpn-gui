package ui

import (
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

// shuttingDown is set when the application is terminating (Quit, window
// close, SIGINT/SIGTERM). Background goroutines must stop scheduling work on
// the Fyne main thread once it is set: after the driver queue is drained,
// fyne.Do/DoAndWait callbacks run inline on the calling goroutine, which
// would trigger "Error in Fyne call thread" warnings.
var shuttingDown atomic.Bool

// isShuttingDown reports whether the application is terminating.
func isShuttingDown() bool {
	return shuttingDown.Load()
}

func init() {
	// Observe termination signals as early as possible: the Fyne driver
	// performs a graceful shutdown on SIGINT/SIGTERM, so flag our loops
	// before the event queue is drained.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		shuttingDown.Store(true)
	}()
}
