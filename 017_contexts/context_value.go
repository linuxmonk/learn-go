package main

import (
	"context"
	"fmt"
)

type TraceIDKey int64

func main() {

	const traceIDKey TraceIDKey = 0

	traceID := "228321F5-33D1-4CD1-976C-32D1785BB639"

	// Save a value in the context
	ctx := context.WithValue(context.Background(), traceIDKey, traceID)

	// Retrieve the value from the context
	if uuid, ok := ctx.Value(traceIDKey).(string); ok {
		fmt.Println("TraceID:", uuid)
	}

	// Retrieve value from context using wrong type of value fails.
	if _, ok := ctx.Value(0).(string); !ok {
		fmt.Println("No TraceID found")
	}
}
