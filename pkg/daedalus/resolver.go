package daedalus

import (
	"errors"
	"fmt"
)

type resolver struct {
	data      map[string]Data
	push_data func(Step, []Data) error
}

func new_resolver(database_size int, with_checks bool) *resolver {
	resolver := &resolver{
		data: make(map[string]Data, database_size),
	}

	if with_checks {
		resolver.push_data = resolver.push_data_with_checks
	} else {
		resolver.push_data = resolver.push_data_without_checks
	}

	return resolver
}

func (r *resolver) push_data_without_checks(_ Step, data []Data) error {
	// NOTE: no need to check if the key exists, as the key is always guaranteed to exist
	// thanks to the conveyor's build process
	if r.data != nil {
		for _, value := range data {
			r.data[value.GetName()] = value
		}
	}

	return nil
}

func (r *resolver) push_data_with_checks(step Step, data []Data) error {
	// NOTE: no need to check if the key exists, as the key is always guaranteed to exist
	// thanks to the conveyor's build process
	declared_output_data := make(map[string]bool, len(data))

	for _, val := range step.GetOutputData() {
		declared_output_data[val] = true
	}

	output_not_decalred := make([]string, 0, len(data))

	if r.data != nil {
		for _, value := range data {
			name := value.GetName()

			if _, ok := declared_output_data[name]; !ok {
				output_not_decalred = append(output_not_decalred, name)
				continue
			}

			declared_output_data[name] = false
			r.data[name] = value
		}
	}

	var err error = nil

	if len(output_not_decalred) > 0 {
		err = errors.New("\toutput data not declared: " + fmt.Sprintf("%v", output_not_decalred))
	}

	not_provided := make([]string, 0, len(declared_output_data))

	for key, val := range declared_output_data {
		if val {
			not_provided = append(not_provided, key)
		}
	}

	if len(not_provided) > 0 {
		err = errors.Join(err, errors.New("\toutput data not provided: "+fmt.Sprintf("%v", not_provided)))
	}

	return err
}

func (r *resolver) get_data_for_step(step Step) map[string]Data {
	// NOTE: no need to check if the key exists, as the key is always guaranteed to exist
	// thanks to the conveyor's build process
	required_data := step.GetRequiredData()
	data := make(map[string]Data, len(required_data))

	for _, key := range required_data {
		data[key] = r.data[key]
	}

	return data
}
