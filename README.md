# Procfs

Procfs is a parser for the /proc virtual filesystem on Linux.

But don't use it for production yet, I'm still refining it.

## Installation

go get github.com/jandre/procfs 

## Examples

See the `*_test` files. 

```go
package main

import "github.com/jandre/procfs"
import "fmt"
import "log"
import "strings"

func main() {
  processes, err := procfs.Processes(false, false)

  if err != nil || len(processes) <= 0 {
    log.Fatal("ERROR")
  }

  for _, p := range processes {
    fmt.Printf("\nPID: %-10d\nCWD: %-30s\nEXE: %-30s", p.Pid, p.Cwd, p.Exe)
    fmt.Printf("\nCMDLINE: %-70s\n", strings.Join(p.Cmdline, " "))
  }
}
```

E# Documentation

Documentation can be found at: http://godoc.org/github.com/jandre/procfs
