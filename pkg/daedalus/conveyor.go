package daedalus

import (
	"errors"
	"fmt"

	idcounter "github.com/eightlay/daedalus/internal/id_counter"
)

type conveyor struct {
	is_built         bool
	stages           map[int]*stage
	stage_id_counter *idcounter.IdCounter
}

func new_conveyor() *conveyor {
	return &conveyor{
		is_built:         false,
		stages:           map[int]*stage{},
		stage_id_counter: idcounter.NewIdCounter(),
	}
}

func (c *conveyor) run(resolver *resolver) error {
	if !c.is_built {
		return errors.New("conveyor is not built")
	}

	execution_order := sort_map_keys(c.stages)

	for _, id := range execution_order {
		if err := c.stages[id].run(resolver); err != nil {
			return prepend_to_error(fmt.Sprintf("stage %d,", id), err)
		}
	}
	return nil
}

func (c *conveyor) build() (int, error) {
	// TODO
	c.is_built = true
	return 0, nil
}

func (c *conveyor) perform_action(fn interface{}, stage_id int, args ...interface{}) ([]interface{}, error) {
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

func (c *conveyor) check_stage_id(stage_id int) error {
	if _, ok := c.stages[stage_id]; !ok {
		return errors.New("stage not found")
	}
	return nil
}

func (c *conveyor) add_stage(stage *stage) int {
	c.is_built = false
	c.stages[c.stage_id_counter.Next()] = stage
	return c.stage_id_counter.Current()
}

func (c *conveyor) add_step(stage_id int, step Step) (int, error) {
	res, err := c.perform_action(c.stages[stage_id].add_step, stage_id, step)
	return res[0].(int), err
}

func (c *conveyor) del_stage(stage_id int) error {
	_, err := c.perform_action(func() { delete(c.stages, stage_id) }, stage_id)
	return err
}

func (c *conveyor) del_step(stage_id, step_id int) error {
	_, err := c.perform_action(c.stages[stage_id].del_step, stage_id, step_id)
	return err
}

func (c *conveyor) clear() {
	c.is_built = false
	c.stages = map[int]*stage{}
	c.stage_id_counter.Clear()
}

func (c *conveyor) clear_stage(stage_id int) error {
	_, err := c.perform_action(c.stages[stage_id].clear, stage_id)
	return err
}

func (c *conveyor) get_stages_number() int {
	return len(c.stages)
}

func (c *conveyor) get_stage_steps_number(stage_id int) int {
	return c.stages[stage_id].get_steps_number()
}
