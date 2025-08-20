package error

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// StackFrame represents a single frame in the stack trace
type StackFrame struct {
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

func (sf StackFrame) String() string {
	return fmt.Sprintf("%s:%d %s", sf.File, sf.Line, sf.Function)
}

// StackTrace represents the full stack trace
type StackTrace []*StackFrame

func (st StackTrace) String() string {
	var lines []string
	for _, frame := range st {
		lines = append(lines, frame.String())
	}
	return strings.Join(lines, "\n")
}

// DomainError represents domain-level errors
type DomainError struct {
	Code      string                 `json:"code"`
	Status    int                    `json:"status"`
	Message   string                 `json:"message"`
	Details   map[string]interface{} `json:"details,omitempty"`
	Stack     StackTrace             `json:"-"`
	Timestamp time.Time              `json:"timestamp"`
	Cause     error                  `json:"-"`
}

func (e *DomainError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *DomainError) Unwrap() error {
	return e.Cause
}

func (e *DomainError) StackTrace() StackTrace {
	return e.Stack
}

func (e *DomainError) Is(target error) bool {
	t, ok := target.(*DomainError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

// Format implements fmt.Formatter for detailed error display
func (e *DomainError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%s (%s)\n", e.Message, e.Code)
			if e.Cause != nil {
				fmt.Fprintf(s, "Caused by: %v\n", e.Cause)
			}
			if len(e.Details) > 0 {
				fmt.Fprintf(s, "Details: %+v\n", e.Details)
			}
			fmt.Fprintf(s, "Timestamp: %s\n", e.Timestamp.Format(time.RFC3339))
			fmt.Fprintf(s, "Stack trace:\n%s", e.Stack)
		} else {
			fmt.Fprintf(s, "%s", e.Error())
		}
	case 's':
		fmt.Fprintf(s, "%s", e.Error())
	}
}

// captureStackTrace captures the current stack trace
func captureStackTrace(skip int) StackTrace {
	const maxDepth = 32
	var pcs [maxDepth]uintptr

	n := runtime.Callers(skip+2, pcs[:])
	frames := make(StackTrace, 0, n)
	callersFrames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := callersFrames.Next()
		frames = append(frames, &StackFrame{
			Function: frame.Function,
			File:     frame.File,
			Line:     frame.Line,
		})
		if !more {
			break
		}
	}

	return frames
}
