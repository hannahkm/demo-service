package main

import (
	"strconv"
	"time"

	"github.com/DataDog/dd-trace-go/v2/ddtrace/tracer"
)

func main() {
	tracer.Start(
		tracer.WithService("subservice1"),
		tracer.WithEnv("hannahskim"),
		tracer.WithRuntimeMetrics(),
		// tracer.WithSendRetries(10),
	)
	defer tracer.Stop()

	for t := 0; ; t++ {
		root := tracer.StartSpan("trace1." + strconv.Itoa(t))

		for s := 0; s < 5; s++ {
			child := root.StartChild("span1." + strconv.Itoa(t) + "." + strconv.Itoa(s))
			child.SetTag("span.type", "dummy1")
			child.SetTag("example-tag", "tag-value")
			child.SetTag("tag-int", 33)
			child.SetTag("tag-float", 0.01)
			child.SetTag("tag-bool", true)
			child.AddEvent("test-event")
			child.Finish()
		}
		root.Finish()
		time.Sleep(100 * time.Millisecond)
	}
}
