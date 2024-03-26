package daedalus

import "errors"

type resolver struct {
	data map[string]Data
}

func new_resolver(database_size int) *resolver {
	return &resolver{
		data: make(map[string]Data, database_size),
	}
}

func (r *resolver) push_data(step Step, data []Data) error {
	// NOTE: no need to check if the key exists, as the key is always guaranteed to exist
	// thanks to the conveyor's build process
	declared_output_data := make(map[string]bool, len(data))

	for _, val := range step.GetOutputData() {
		declared_output_data[val] = true
	}

	if r.data != nil {
		for _, value := range data {
			name := value.GetName()

			if _, ok := declared_output_data[name]; !ok {
				return errors.New("output data not declared in Step.GetOutputData: " + name)
			}
			r.data[name] = value
		}
	}

	return nil
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
