package locales

import (
	"sync"
)

type Trigger struct {
	ch      chan bool
	methods []func()
	wg      sync.WaitGroup
	mu      sync.Mutex
	started bool
}

var trigger *Trigger

// GetTrigger returns a pointer to the Trigger singleton. If the trigger has not
// been initialized, it is initialized with a new channel and returned. The
// returned trigger is safe to use concurrently.
func GetTrigger() *Trigger {
	if trigger == nil {
		trigger = &Trigger{
			ch: make(chan bool),
		}
	}
	return trigger
}

// AddMethod adds the given method to the trigger. When the trigger is activated,
// all registered methods are executed in the order they were registered. The
// methods are executed in a goroutine, so the order of execution is not
// guaranteed. Use this method to register methods that should be executed when
// the locale changes.
func (t *Trigger) AddMethod(method func()) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.methods = append(t.methods, method)
}

// Start starts the trigger goroutine. This goroutine waits for the trigger
// channel and when a message is received, executes all registered methods in a
// goroutine. If the trigger has already been started, this function does
// nothing.
func (t *Trigger) Start() {
	if !t.started {
		go func() {
			for {
				<-t.ch // Wait for the trigger
				t.executeMethods()
			}
		}()
		t.started = true
	}
}

// executeMethods executes all registered methods in a goroutine. It creates a
// copy of the methods slice to avoid the race condition when calling AddMethod
// while the methods are being executed. The method execution is done in a
// goroutine to avoid blocking the main goroutine. The WaitGroup is used to
// wait for all methods to finish before returning.
func (t *Trigger) executeMethods() {
	t.mu.Lock()
	methodsCopy := make([]func(), len(t.methods))
	copy(methodsCopy, t.methods)
	t.mu.Unlock()

	t.wg.Add(len(methodsCopy))
	for _, method := range methodsCopy {
		go func(m func()) {
			defer t.wg.Done()
			m()
		}(method)
	}
	t.wg.Wait()
}

// Activate sends a message to the trigger channel to execute all registered
// methods. This should be called when the locale changes, for example when
// calling Load.
func (t *Trigger) Activate() {
	t.ch <- true
}
