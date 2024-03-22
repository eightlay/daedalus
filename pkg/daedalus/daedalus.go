package daedalus

type Daedalus struct {
	conv     *Conveyor
	resolver *Resolver
}

func NewDaedalus() *Daedalus {
	return &Daedalus{
		conv:     new_conveyor(),
		resolver: new_resolver(),
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
	d.handle_error(d.conv.build())
}

func (d *Daedalus) AddStage() int {
	return d.conv.add_stage(new_stage())
}

func (d *Daedalus) AddStep(stage_id int, step Step) (int, int) {
	if step == nil {
		panic("step is nil")
	}

	if stage_id == -1 {
		stage := new_stage()
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
