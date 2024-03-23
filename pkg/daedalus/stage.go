package daedalus

import (
	"errors"
	"sync"

	idcounter "github.com/eightlay/daedalus/internal/id_counter"
)

type Stage struct {
	steps                   map[int]Step
	steps_id_counter        *idcounter.IdCounter
	run_steps_as_goroutines bool
}

func new_stage(run_steps_as_goroutines ...bool) *Stage {
	if len(run_steps_as_goroutines) > 1 {
		panic("run_steps_as_goroutines should be a single value")
	}

	if len(run_steps_as_goroutines) == 0 {
		run_steps_as_goroutines = append(run_steps_as_goroutines, false)
	}

	return &Stage{
		steps:                   map[int]Step{},
		steps_id_counter:        idcounter.NewIdCounter(),
		run_steps_as_goroutines: run_steps_as_goroutines[0],
	}
}

func (s *Stage) run(resolver *Resolver) {
	if s.run_steps_as_goroutines {
		s.run_as_goroutines(resolver)
		return
	}

	execution_order := sort_map_keys(s.steps)

	for id := range execution_order {
		s.steps[id].Run(resolver)
	}
}

func (s *Stage) run_as_goroutines(resolver *Resolver) {
	var wg sync.WaitGroup

	for _, step := range s.steps {
		wg.Add(1)
		go func(step Step) {
			defer wg.Done()
			step.Run(resolver)
		}(step)
	}

	wg.Wait()
}

func (s *Stage) add_step(step Step) int {
	s.steps[s.steps_id_counter.Next()] = step
	return s.steps_id_counter.Current()
}

func (s *Stage) del_step(step_id int) error {
	if _, ok := s.steps[step_id]; !ok {
		return errors.New("Step not found")
	}
	delete(s.steps, step_id)
	return nil
}

func (s *Stage) clear() {
	s.steps = map[int]Step{}
	s.steps_id_counter.Clear()
}

func (s *Stage) get_steps_number() int {
	return len(s.steps)
}
