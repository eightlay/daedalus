package daedalus

import "errors"

type Conveyor struct {
	is_built bool
	stages   []*Stage
}

func new_conveyor() *Conveyor {
	return &Conveyor{
		is_built: false,
		stages:   []*Stage{},
	}
}

func (c *Conveyor) run(resolver *Resolver) error {
	if !c.is_built {
		return errors.New("Conveyor is not built")
	}

	for _, stage := range c.stages {
		stage.run(resolver)
	}
	return nil
}

func (c *Conveyor) build() error {
	// TODO
	c.is_built = true
	return nil
}

func (c *Conveyor) check_stage_index(stage_ind int) {
	if stage_ind < 0 || stage_ind >= len(c.stages) {
		panic("stage_ind is out of range")
	}
}

func (c *Conveyor) add_stage(stage *Stage) int {
	c.is_built = false
	stage_ind := len(c.stages)
	c.stages = append(c.stages, stage)
	return stage_ind
}

func (c *Conveyor) add_step(stage_ind int, step Step) int {
	c.is_built = false
	c.check_stage_index(stage_ind)
	return c.stages[stage_ind].add_step(step)
}

func (c *Conveyor) del_stage(stage_ind int) {
	c.is_built = false
	c.check_stage_index(stage_ind)
	c.stages = append(c.stages[:stage_ind], c.stages[stage_ind+1:]...)
}

func (c *Conveyor) del_step(stage_ind, step_ind int) {
	c.is_built = false
	c.check_stage_index(stage_ind)
	c.stages[stage_ind].del_step(step_ind)
}

func (c *Conveyor) clear() {
	c.is_built = false
	c.stages = []*Stage{}
}

func (c *Conveyor) clear_stage(stage_ind int) {
	c.is_built = false
	c.check_stage_index(stage_ind)
	c.stages[stage_ind].clear()
}
