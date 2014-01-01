package procfs

import (
	"testing"
)

func TestParseMeminfo(t *testing.T) {
	// set the GLOBAL_SYSTEM_START
	meminfo, err :=  ParseMeminfo("./testfiles/meminfo")

	if err != nil {
		t.Fatal("Got error", err)
	}

	if meminfo == nil {
		t.Fatal("meminfo is missing")
	}

	// log.Println("Meminfo:", meminfo)
	if meminfo.MemTotal != 1011932 {
		t.Fatal("Expected 1011932 from MemToal")
	}
}

