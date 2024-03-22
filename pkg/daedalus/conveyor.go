package daedalus

import (
	"errors"

	idcounter "github.com/eightlay/daedalus/internal/id_counter"
)

type Conveyor struct {
	is_built         bool
	stages           map[int]*Stage
	stage_id_counter *idcounter.IdCounter
}

func new_conveyor() *Conveyor {
	return &Conveyor{
		is_built:         false,
		stages:           map[int]*Stage{},
		stage_id_counter: idcounter.NewIdCounter(),
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

func (c *Conveyor) perform_action(fn interface{}, stage_id int, args ...interface{}) ([]interface{}, error) {
	if err := c.check_stage_id(stage_id); err != nil {
		return nil, err
	}
	c.is_built = false

	argSlice := make([]interface{}, len(args))
	copy(argSlice, args)

	if f, ok := fn.(func()); ok {
		f()
	} else if f, ok := fn.(func(Step) int); ok {
		return []interface{}{f(argSlice[0].(Step))}, nil
	} else if f, ok := fn.(func(int, int) error); ok {
		return nil, f(argSlice[0].(int), argSlice[1].(int))
	} else if f, ok := fn.(func(int) error); ok {
		return nil, f(argSlice[0].(int))
	} else {
		panic("Invalid function type")
	}

	return nil, nil
}

func (c *Conveyor) check_stage_id(stage_id int) error {
	if _, ok := c.stages[stage_id]; !ok {
		return errors.New("Stage not found")
	}
	return nil
}

func (c *Conveyor) add_stage(stage *Stage) int {
	c.is_built = false
	c.stages[c.stage_id_counter.Next()] = stage
	return c.stage_id_counter.Current()
}

func (c *Conveyor) add_step(stage_id int, step Step) (int, error) {
	res, err := c.perform_action(c.stages[stage_id].add_step, stage_id, step)
	return res[0].(int), err
}

func (c *Conveyor) del_stage(stage_id int) error {
	_, err := c.perform_action(func() { delete(c.stages, stage_id) }, stage_id)
	return err
}

func (c *Conveyor) del_step(stage_id, step_id int) error {
	_, err := c.perform_action(c.stages[stage_id].del_step, stage_id, step_id)
	return err
}

func (c *Conveyor) clear() {
	c.is_built = false
	c.stages = map[int]*Stage{}
	c.stage_id_counter.Clear()
}

func (c *Conveyor) clear_stage(stage_id int) error {
	_, err := c.perform_action(c.stages[stage_id].clear, stage_id)
	return err
}
