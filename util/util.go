package util

//
// Copyright Jen Andre (jandre@gmail.com)
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

/*
#include <unistd.h>
*/
import "C"

//
// fetch the system start time from /proc/stat
// look for the btime line, and parse it out
//
//
func systemStart() int64 {
	str, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		log.Fatal("Unable to read btime from /proc/stat - is this Linux?", err)
	}
	lines := strings.Split(string(str), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "btime") {
			parts := strings.Split(line, " ")
			// format is btime 1388417200
			val, err := strconv.ParseInt(parts[1], 10, 32)
			if err != nil {
				log.Fatal("Unable to convert btime value in /proc/stat to int64", parts[1], err)
			}
			return int64(val)
		}
	}

	log.Fatal(`No btime found in /proc/stat.
	This value is needed to calculate timestamps`)

	return 0
}

var GLOBAL_SYSTEM_START int64 = systemStart()

//
// converts jiffies to seconds since epoch
//
func JiffiesToEpoch(jiffies int64) int64 {
	ticks := C.sysconf(C._SC_CLK_TCK)
	return GLOBAL_SYSTEM_START + jiffies/int64(ticks)
}

type EpochTimestamp struct {
	jiffies int64
}

//
// Create a new epoch timestamp from a `jiffies since
// system start time` value
//
//
func NewEpochTimestamp(jiffies int64) EpochTimestamp {
	return EpochTimestamp{jiffies: jiffies}
}

//
// Get the timestamp in seconds since Epoch
//
func (e *EpochTimestamp) EpochSeconds() int64 {
	return JiffiesToEpoch(e.jiffies)
}

//
// Get the timestamp in milliseconds since Epoch
//
func (e *EpochTimestamp) EpochMilliseconds() int64 {
	return e.EpochSeconds() * 1000
}

//
// Get the timestamp as a Time object
//
func (e *EpochTimestamp) Time() time.Time {
	return time.Unix(e.EpochSeconds(), 0)
}

//
// Get the jiffies value
//
func (e *EpochTimestamp) Jiffies() int64 {
	return e.jiffies
}
