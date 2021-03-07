package trace

import (
	"fmt"
	"io"
)

// Tracer is Interface recording something in code
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off returns Tracer to ignore Trace method call.
func Off() Tracer {
	return &nilTracer{}
}

// New impliments Trace().
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
