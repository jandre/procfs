package stat

import (
	"io/ioutil"
	"github.com/jandre/procfs/util"
	"reflect"
	"strings"
)

//
// Abstraction for /proc/<pid>/stat
//
type Stat struct {
	Pid                 int                 // process id
	Comm                string              // filename of the executable
	State               string              // state (R is running, S is sleeping, D is sleeping in an uninterruptible wait, Z is zombie, T is traced or stopped)
	Ppid                int                 // process id of the parent process
	Pgrp                int                 // pgrp of the process
	Session             int                 // session id
	TtyNr               int                 // tty the process uses
	Tpgid               int                 // pgrp of the tty
	Flags               int64               // task flags
	Minflt              int64               // number of minor faults
	Cminflt             int64               // number of minor faults with child's
	Majflt              int64               // number of major faults
	Cmajflt             int64               // number of major faults with child's
	Utime               util.EpochTimestamp // user mode jiffies
	Stime               util.EpochTimestamp // kernel mode jiffies
	Cutime              util.EpochTimestamp // user mode jiffies with child's
	Cstime              util.EpochTimestamp // kernel mode jiffies with child's
	Priority            int64               // priority level
	Nice                int64               // nice level
	NumThreads          int64               // number of threads
	Itrealvalue         int64               // (obsolete, always 0)
	Starttime           util.EpochTimestamp // time the process started after system boot
	Vsize               int64               // virtual memory size
	Rss                 int64               // resident set memory size
	Rlim                uint64              // current limit in bytes on the rss
	Startcode           int64               // address above which program text can run
	Endcode             int64               // address below which program text can run
	Startstack          int64               // address of the start of the main process stack
	Kstkesp             int64               // current value of ESP
	Kstkeip             int64               // current value of EIP
	Signal              int64               // bitmap of pending signals
	Blocked             int64               // bitmap of blocked signals
	Sigignore           int64               // bitmap of ignored signals
	Sigcatch            int64               // bitmap of catched signals
	Wchan               uint64              // address where process went to sleep
	Nswap               int64               // (place holder)
	Cnswap              int64               // (place holder)
	ExitSignal          int                 // signal to send to parent thread on exit
	Processor           int                 // which CPU the task is scheduled on
	RtPriority          int64               // realtime priority
	Policy              int64               // scheduling policy (man sched_setscheduler)
	DelayacctBlkioTicks int64               // time spent waiting for block IO
}

//
// Load /proc/<pid>/stat from path
//
func New(path string) (*Stat, error) {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(buf), " ")
	stat := &Stat{}

	v := reflect.ValueOf(stat).Elem()
	err = util.StructParser(&v, lines)
	return stat, err
}
