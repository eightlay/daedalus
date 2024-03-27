package daedalus

import (
	"errors"
	"fmt"
	"sync"
	"time"

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

func (s *stage) build(previous_stages_data map[string]bool) (map[string]bool, error) {
	execution_order := sort_map_keys(s.steps)
	stage_data := map[string]bool{}

	for _, id := range execution_order {
		missing_data := []error{}

		for _, data := range s.steps[id].GetRequiredData() {
			if _, ok := previous_stages_data[data]; ok {
				continue
			}
			if _, ok := stage_data[data]; ok {
				continue
			}
			missing_data = append(missing_data, fmt.Errorf("\tstep %d: %s", id, data))
		}

		if s.run_steps_as_goroutines {
			continue
		}

		for _, data := range s.steps[id].GetOutputData() {
			stage_data[data] = true
		}

		if len(missing_data) > 0 {
			return nil, combine_errors(errors.New("missing data: "), missing_data)
		}
	}

	return stage_data, nil
}

func (s *stage) run(resolver *resolver) (map[int]*RunStats, error) {
	if s.run_steps_as_goroutines {
		return s.run_as_goroutines(resolver)
	}

	execution_order := sort_map_keys(s.steps)
	stats := make(map[int]*RunStats, len(execution_order))

	for _, id := range execution_order {
		step_stats, err := s.run_step(id, s.steps[id], resolver)

		if err != nil {
			return nil, err
		}

		stats[id] = step_stats
	}

	return stats, nil
}

func (s *stage) run_as_goroutines(resolver *resolver) (map[int]*RunStats, error) {
	steps_errors := []error{}
	var wg sync.WaitGroup

	type stepStats struct {
		step_id int
		stats   *RunStats
	}

	stats_chan := make(chan stepStats, len(s.steps))

	for i, step := range s.steps {
		wg.Add(1)
		go func(step Step) {
			defer wg.Done()
			stats, err := s.run_step(i, step, resolver)

			if err != nil {
				steps_errors = append(steps_errors, err)
			}

			stats_chan <- stepStats{i, stats}
		}(step)
	}

	wg.Wait()

	if len(steps_errors) > 0 {
		return nil, combine_errors(errors.New("errors occurred during goroutine execution"), steps_errors)
	}

	stats := make(map[int]*RunStats, len(s.steps))

	for range s.steps {
		step_stats := <-stats_chan
		stats[step_stats.step_id] = step_stats.stats
	}

	return stats, nil
}

func (s *stage) run_step(step_num int, step Step, resolver *resolver) (*RunStats, error) {
	step_data := resolver.get_data_for_step(step)

	time_start := time.Now()
	step_result := step.Run(step_data)
	duration := time.Since(time_start)

	if err := resolver.push_data(step, step_result); err != nil {
		return nil, prepend_to_error(fmt.Sprintf("stage %d:\n", step_num), err)
	}
	return new_run_stats(duration), nil
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
