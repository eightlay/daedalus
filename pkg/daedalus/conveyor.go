package daedalus

import (
	"errors"
	"fmt"

	idcounter "github.com/eightlay/daedalus/internal/id_counter"
)

type conveyor struct {
	is_built         bool
	is_run_done      bool
	stages           map[int]*stage
	stage_id_counter *idcounter.IdCounter
	stage_run_stats  map[int]map[int]*RunStats
}

func new_conveyor() *conveyor {
	return &conveyor{
		is_built:         false,
		is_run_done:      false,
		stages:           map[int]*stage{},
		stage_id_counter: idcounter.NewIdCounter(),
		stage_run_stats:  map[int]map[int]*RunStats{},
	}
}

func (c *conveyor) reset_flags() {
	c.is_built = false
	c.is_run_done = false
}

func (c *conveyor) run(resolver *resolver, vh *verbosity_handler) error {
	if !c.is_built {
		return errors.New("conveyor is not built")
	}

	execution_order := sort_map_keys(c.stages)

	vh.start()

	for _, id := range execution_order {
		stats, err := c.stages[id].run(resolver, vh)

		if err != nil {
			return prepend_to_error(fmt.Sprintf("stage %d,", id), err)
		}

		c.stage_run_stats[id] = stats
	}

	vh.end()

	c.is_run_done = true
	return nil
}

func (c *conveyor) build() (int, error) {
	execution_order := sort_map_keys(c.stages)
	total_data := map[string]bool{}

	for i, id := range execution_order {
		data, err := c.stages[id].build(total_data)

		if err != nil {
			return 0, prepend_to_error(fmt.Sprintf("stage %d,", i), err)
		}

		for k := range data {
			total_data[k] = true
		}
	}

	c.is_built = true
	c.is_run_done = false
	return len(total_data), nil
}

func (c *conveyor) perform_action(fn interface{}, stage_id int, args ...interface{}) ([]interface{}, error) {
	if err := c.check_stage_id(stage_id); err != nil {
		return nil, err
	}
	c.reset_flags()

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
	c.reset_flags()
	stage_id := c.stage_id_counter.Next()
	stage.id = stage_id
	c.stages[stage_id] = stage
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
	c.reset_flags()
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

func (c *conveyor) get_run_stats() (map[int]map[int]*RunStats, error) {
	if !c.is_run_done {
		return nil, errors.New("run has not been performed yet")
	}
	return c.stage_run_stats, nil
}
