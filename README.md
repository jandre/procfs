# Procfs

Procfs is a parser for the /proc virtual filesystem on Linux.

But don't use it for production yet, I'm still refining it.

# Install

go get github.com/jandre/procfs 

# Example

See the `*_test` files. 

```go

// fetch all processes from /proc
processes, err := procfs.AllProcesses();

```

# Documentation
