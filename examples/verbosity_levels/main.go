package main

import (
	"fmt"

	"github.com/eightlay/daedalus/pkg/daedalus"
)

type MyFancyStep struct {
}

func (s *MyFancyStep) Run(data map[string]daedalus.Data) []daedalus.Data {
	return nil
}

func (s *MyFancyStep) GetRequiredData() []string {
	return []string{}
}

func (s *MyFancyStep) GetOutputData() []string {
	return []string{}
}

func main() {
	d := daedalus.NewDaedalus()
	d.AddStep(-1, &MyFancyStep{})

	d.Build()

	// SILENT is the default verbosity level
	fmt.Println("SILENT is the default verbosity level")
	d.Run()

	// Set verbosity level to START_END
	fmt.Println("Set verbosity level to START_END")
	d.SetVerbosity(daedalus.START_END)
	d.Run()

	// Set verbosity level to STAGES
	fmt.Println("Set verbosity level to STAGES")
	d.SetVerbosity(daedalus.STAGES)
	d.Run()

	// Set verbosity level to STEPS
	fmt.Println("Set verbosity level to STEPS")
	d.SetVerbosity(daedalus.STEPS)
	d.Run()
}
