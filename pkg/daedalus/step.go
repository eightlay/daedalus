package daedalus

type Step interface {
	Run(map[string]Data) []Data
	GetRequiredData() []string
	GetOutputData() []string
}
