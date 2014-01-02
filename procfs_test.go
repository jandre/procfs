package procfs

import (
	"testing"
	"log"
)
// func BenchmarkAllProc() {

// }
//
func TestAllProc(t * testing.T) {
	procs, err := Processes(false)
	if err != nil {
		t.Fatal(err)
	}
	if len(procs) <= 0 {
		t.Fatal("procs length must be > 0")
	}

	log.Println("Pid 1", procs[1])
	// for ;; {
		// // noop
	// }
}
