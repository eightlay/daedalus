package daedalus

type Step interface {
	Run(*Resolver)
	GetRequiredData() []string
	GetOutputData() []string
}
