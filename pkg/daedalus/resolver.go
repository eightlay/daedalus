package daedalus

type resolver struct {
	data map[string]Data
}

func new_resolver(database_size int) *resolver {
	return &resolver{
		data: make(map[string]Data, database_size),
	}
}

func (r *resolver) push_data(data []Data) {
	// NOTE: no need to check if the key exists, as the key is always guaranteed to exist
	// thanks to the conveyor's build process
	if r.data != nil {
		for _, value := range data {
			r.data[value.GetName()] = value
		}
	}
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
