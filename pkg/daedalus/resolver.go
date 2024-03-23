package daedalus

type Resolver struct {
	data map[string]Data
}

func new_resolver(database_size int) *Resolver {
	return &Resolver{
		data: make(map[string]Data, database_size),
	}
}

func (r *Resolver) PushData(data Data) {
	// NOTE: no need to check if the key exists, as the key is always guaranteed to exist
	// thanks to the conveyor's build process
	r.data[data.GetName()] = data
}

func (r *Resolver) GetData(name string) Data {
	// NOTE: no need to check if the key exists, as the key is always guaranteed to exist
	// thanks to the conveyor's build process
	return r.data[name]
}

func (r *Resolver) DelData(name string) {
	// NOTE: no need to check if the key exists, as the key is always guaranteed to exist
	// thanks to the conveyor's build process
	delete(r.data, name)
}
