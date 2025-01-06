# Go Array Index Out of Bounds

This example demonstrates a common error in Go: accessing an array index that is out of bounds. Go arrays have a fixed size, and attempting to access an element beyond that size will result in a runtime panic.  The solution shows how to use length checking or error handling to prevent this.

## How to Run

1. Save the code in `bug.go` and `bugSolution.go`.
2. Run `go run bug.go` to see the panic.
3. Run `go run bugSolution.go` to see the safe version.
