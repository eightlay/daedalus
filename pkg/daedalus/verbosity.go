package daedalus

import (
	"fmt"
	"time"
)

type Verbosity int

const (
	SILENT Verbosity = iota
	START_END
	STAGES
	STEPS
)

type verbosity_handler struct {
	start       func()
	end         func()
	stage_start func(id int)
	stage_end   func(id int)
	step_start  func(id int)
	step_end    func(id int)
	verb        Verbosity
}

func new_verbosity_handler(verbosity Verbosity) *verbosity_handler {
	vh := &verbosity_handler{}
	vh.set_verbosity(verbosity)
	return vh
}

func (v *verbosity_handler) set_verbosity(verbosity Verbosity) {
	v.verb = verbosity

	v.start = v.nothing_no_args
	v.end = v.nothing_no_args
	v.stage_start = v.nothing_with_id
	v.stage_end = v.nothing_with_id
	v.step_start = v.nothing_with_id
	v.step_end = v.nothing_with_id

	if verbosity >= START_END {
		v.start = v._start
		v.end = v._end
	}
	if verbosity >= STAGES {
		v.stage_start = v._stage_start
		v.stage_end = v._stage_end
	}
	if verbosity >= STEPS {
		v.step_start = v._step_start
		v.step_end = v._step_end
	}
}

func (v *verbosity_handler) nothing_no_args() {
}

func (v *verbosity_handler) nothing_with_id(id int) {
}

func (v *verbosity_handler) print(msg ...any) {
	fmt.Printf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintln(msg...))
}

func (v *verbosity_handler) _start() {
	v.print("Run started")
}

func (v *verbosity_handler) _end() {
	v.print("Run ended")
}

func (v *verbosity_handler) _stage_start(id int) {
	v.print("Stage ", id, "started")
}

func (v *verbosity_handler) _stage_end(id int) {
	v.print("Stage", id, "ended")
}

func (v *verbosity_handler) _step_start(id int) {
	v.print("Step", id, "started")
}

func (v *verbosity_handler) _step_end(id int) {
	v.print("Step", id, "ended")
}
