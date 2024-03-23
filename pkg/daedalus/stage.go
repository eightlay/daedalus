package daedalus

import (
	"errors"

	idcounter "github.com/eightlay/daedalus/internal/id_counter"
)

type Stage struct {
	steps            map[int]Step
	steps_id_counter *idcounter.IdCounter
}

func new_stage() *Stage {
	return &Stage{
		steps:            map[int]Step{},
		steps_id_counter: idcounter.NewIdCounter(),
	}
}

func (s *Stage) run(resolver *Resolver) {
	execution_order := sort_map_keys(s.steps)

	for id := range execution_order {
		s.steps[id].Run(resolver)
	}
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
