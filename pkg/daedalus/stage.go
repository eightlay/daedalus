package daedalus

type Stage struct {
	steps []Step
}

func new_stage() *Stage {
	return &Stage{
		steps: []Step{},
	}
}

func (s *Stage) run(resolver *Resolver) {
	for _, step := range s.steps {
		step.Run(resolver)
	}
}

func (s *Stage) add_step(step Step) int {
	step_ind := len(s.steps)
	s.steps = append(s.steps, step)
	return step_ind
}

func (s *Stage) del_step(step_ind int) {
	if step_ind < 0 || step_ind >= len(s.steps) {
		panic("step_ind is out of range")
	}
	s.steps = append(s.steps[:step_ind], s.steps[step_ind+1:]...)
}

func (s *Stage) clear() {
	s.steps = []Step{}
}
