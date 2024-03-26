package daedalus

import (
	"errors"
	"sync"

	idcounter "github.com/eightlay/daedalus/internal/id_counter"
)

type stage struct {
	steps                   map[int]Step
	steps_id_counter        *idcounter.IdCounter
	run_steps_as_goroutines bool
}

func new_stage(run_steps_as_goroutines ...bool) *stage {
	if len(run_steps_as_goroutines) > 1 {
		panic("run_steps_as_goroutines should be a single value")
	}

	if len(run_steps_as_goroutines) == 0 {
		run_steps_as_goroutines = append(run_steps_as_goroutines, false)
	}

	return &stage{
		steps:                   map[int]Step{},
		steps_id_counter:        idcounter.NewIdCounter(),
		run_steps_as_goroutines: run_steps_as_goroutines[0],
	}
}

func (s *stage) run(resolver *resolver) {
	if s.run_steps_as_goroutines {
		s.run_as_goroutines(resolver)
		return
	}

	execution_order := sort_map_keys(s.steps)

	for id := range execution_order {
		s.run_step(s.steps[id], resolver)
	}
}

func (s *stage) run_as_goroutines(resolver *resolver) {
	var wg sync.WaitGroup

	for _, step := range s.steps {
		wg.Add(1)
		go func(step Step) {
			defer wg.Done()
			s.run_step(step, resolver)
		}(step)
	}

	wg.Wait()
}

func (s *stage) run_step(step Step, resolver *resolver) {
	step_data := resolver.get_data_for_step(step)
	step_result := step.Run(step_data)
	resolver.push_data(step_result)
}

func (s *stage) add_step(step Step) int {
	s.steps[s.steps_id_counter.Next()] = step
	return s.steps_id_counter.Current()
}

func (s *stage) del_step(step_id int) error {
	if _, ok := s.steps[step_id]; !ok {
		return errors.New("Step not found")
	}
	delete(s.steps, step_id)
	return nil
}

func (s *stage) clear() {
	s.steps = map[int]Step{}
	s.steps_id_counter.Clear()
}

func (s *stage) get_steps_number() int {
	return len(s.steps)
}
