package pbs

import "time"

type Task struct {
	FileName     string    // "PYTHON_SCRIPT":"myscript.py",
	AccountSid   string    // id user'a
	Status       string    // job_state
	CreationTime time.Time // qtime
	UpdateTime   time.Time // mtime
	LaunchTime   time.Time // stime
	ServiceId    int       // queue
	ProcessId    int       // pbs id 1005.qot-calc0
}

// Job status
const (
	Running   = "running"
	Stopped   = "stopped"
	Completed = "completed"
	Queued    = "queued"
	Unknown   = "unknown"
)
const LogFile = "logs/pbs.log"
