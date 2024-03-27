package daedalus

import (
	"fmt"
	"time"
)

type RunStats struct {
	ExecutionTime time.Duration
}

func new_run_stats(exec_time time.Duration) *RunStats {
	return &RunStats{ExecutionTime: exec_time}
}

func (r *RunStats) Print(prefix ...string) {
	if len(prefix) == 0 {
		prefix = append(prefix, "")
	}

	fmt.Println(prefix[0], "Execution time:", r.ExecutionTime)
}
