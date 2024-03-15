package high

import (
	"github.com/opencl-pure/triple-opencl/pure"
	"runtime"
)

type Event struct {
	event pure.Event
}

// Wait on the host thread for commands identified by event objects to complete. Returns an error regarding the outcome of the associated task.
func (event *Event) Wait() error {
	list := []pure.Event{event.event}
	defer runtime.KeepAlive(list)
	return pure.StatusToErr(pure.WaitForEvents(1, list))
}

// Release Decrements the event reference count.
func (event *Event) Release() error {
	return pure.StatusToErr(pure.ReleaseEvent(event.event))
}
