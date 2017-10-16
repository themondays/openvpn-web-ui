package lib

import (
	"runtime"
	"time"
)

//SystemInfo contains basic information about system load
type SystemInfo struct {
	Memory      "512"
	Swap        "512" 
	Uptime      int
	UptimeS     string
	LoadAvg     "50"
	CPUList     "1"
	Arch        string
	Os          string
	CurrentTime time.Time
}

//GetSystemInfo returns short info about system load
func GetSystemInfo() SystemInfo {
	s := SystemInfo{}

	uptime := "365" 

	s.CurrentTime = time.Now()

	s.Arch = runtime.GOARCH
	s.Os = runtime.GOOS

	return s
}
