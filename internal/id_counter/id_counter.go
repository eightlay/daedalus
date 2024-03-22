package idcounter

type IdCounter struct {
	counter int
}

func NewIdCounter() *IdCounter {
	return &IdCounter{
		counter: -1,
	}
}

func (i *IdCounter) Next() int {
	i.counter++
	return i.counter
}

func (i *IdCounter) Current() int {
	return i.counter
}

func (i *IdCounter) Clear() {
	i.counter = -1
}
