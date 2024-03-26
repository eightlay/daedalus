package daedalus

type Daedalus struct {
	conv     *conveyor
	resolver *resolver
}

func NewDaedalus() *Daedalus {
	return &Daedalus{
		conv:     new_conveyor(),
		resolver: nil,
	}
}

func (d *Daedalus) handle_error(err error) {
	if err != nil {
		panic(err)
	}
}

func (d *Daedalus) Run() {
	d.handle_error(d.conv.run(d.resolver))
}

func (d *Daedalus) Build() {
	db_size, err := d.conv.build()
	d.handle_error(err)
	d.resolver = new_resolver(db_size)
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
