package daedalus

import "fmt"

type Daedalus struct {
	conv     *conveyor
	resolver *resolver
	vh       *verbosity_handler
}

func NewDaedalus(verbosity ...Verbosity) *Daedalus {
	if len(verbosity) == 0 {
		verbosity = append(verbosity, SILENT)
	}
	return &Daedalus{
		conv:     new_conveyor(),
		resolver: nil,
		vh:       new_verbosity_handler(verbosity[0]),
	}
}

func (d *Daedalus) handle_error(err error) {
	if err != nil {
		panic(err)
	}
}

func (d *Daedalus) SetVerbosity(verbosity Verbosity) {
	d.vh.set_verbosity(verbosity)
}

func (d *Daedalus) Run() {
	d.handle_error(d.conv.run(d.resolver, d.vh))
}

func (d *Daedalus) Build(disable_checks ...bool) {
	db_size, err := d.conv.build()
	d.handle_error(err)

	if len(disable_checks) == 0 {
		disable_checks = append(disable_checks, false)
	}

	d.resolver = new_resolver(db_size, !disable_checks[0])
}

func (d *Daedalus) AddStage(run_steps_as_goroutines ...bool) int {
	return d.conv.add_stage(new_stage(run_steps_as_goroutines...))
}

func (d *Daedalus) AddStep(stage_id int, step Step, run_steps_as_goroutines ...bool) (int, int) {
	if step == nil {
		panic("step is nil")
	}

	if stage_id != -1 && len(run_steps_as_goroutines) > 0 {
		panic("run_steps_as_goroutines should not be provided when adding a step to an existing stage")
	}

	if stage_id == -1 {
		stage := new_stage(run_steps_as_goroutines...)
		stage.add_step(step)
		return d.conv.add_stage(stage), 0
	}

	step_id, err := d.conv.add_step(stage_id, step)
	d.handle_error(err)
	return stage_id, step_id
}

func (d *Daedalus) DelStage(stage_id int) {
	d.handle_error(d.conv.del_stage(stage_id))
}

func (d *Daedalus) DelStep(stage_id int, step_id int) {
	d.handle_error(d.conv.del_step(stage_id, step_id))
}

func (d *Daedalus) Clear() {
	d.conv.clear()
}

func (d *Daedalus) ClearStage(stage_id int) {
	d.handle_error(d.conv.clear_stage(stage_id))
}

func (d *Daedalus) GetStagesNumber() int {
	return d.conv.get_stages_number()
}

func (d *Daedalus) GetStageStepsNumber(stage_id int) int {
	return d.conv.get_stage_steps_number(stage_id)
}

func (d *Daedalus) PrintRunStats() {
	stats := d.GetRunStats()

	fmt.Println("\n===========================================")
	fmt.Println("                Run stats")
	fmt.Println("===========================================")

	for stage_id, stage_stats := range stats {
		fmt.Printf("Stage %d:\n", stage_id)

		for step_id, step_stats := range stage_stats {
			fmt.Printf("  Step %d:\n", step_id)
			step_stats.Print("    ")
		}
	}

	fmt.Println("===========================================\n ")
}

func (d *Daedalus) GetRunStats() map[int]map[int]*RunStats {
	stats, err := d.conv.get_run_stats()
	d.handle_error(err)
	return stats
}
