package trace

// Tracer is Interface recording something in code
type Tracer interface {
	Trace(...interface{})
}
