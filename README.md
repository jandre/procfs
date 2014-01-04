# Procfs

Procfs is a parser for the /proc virtual filesystem on Linux.

But don't use it for production yet, I'm still refining it.

## Installation

go get github.com/jandre/procfs 

## Examples

See the `*_test` files. 

```go
import (
  "github.com/jandre/procfs"
)

// fetch all processes from /proc
// returns a map of pid -> *Process 
processes, err := procfs.Processes();

```

E# Documentation

Documentation can be found at: http://godoc.org/github.com/jandre/procfs
