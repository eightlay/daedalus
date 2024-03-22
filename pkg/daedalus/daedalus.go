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

func (d *Daedalus) Run() {
	err := d.conv.run(d.resolver)

	if err != nil {
		panic(err)
	}
}

func (d *Daedalus) Build() error {
	return d.conv.build()
}

func (d *Daedalus) AddStage() int {
	return d.conv.add_stage(new_stage())
}

func (d *Daedalus) AddStep(stage_ind int, step Step) (int, int) {
	if stage_ind == -1 {
		stage := new_stage()
		stage.add_step(step)
		return d.conv.add_stage(stage), 0
	}
	return stage_ind, d.conv.add_step(stage_ind, step)
}

func (d *Daedalus) DelStage(stage_ind int) {
	d.conv.del_stage(stage_ind)
}

func (d *Daedalus) DelStep(stage_ind int, step_ind int) {
	d.conv.del_step(stage_ind, step_ind)
}

func (d *Daedalus) Clear() {
	d.conv.clear()
}

func (d *Daedalus) ClearStage(stage_ind int) {
	d.conv.clear_stage(stage_ind)
}
