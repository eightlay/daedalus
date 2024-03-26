package main

import (
	"fmt"

	"github.com/eightlay/daedalus/pkg/daedalus"
)

type MyFancyStep struct {
	stage_id int
}

func (s *MyFancyStep) Run(data map[string]daedalus.Data) []daedalus.Data {
	fmt.Printf("MyFancyStep, stage_id == %d\n", s.stage_id)
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
	d.AddStep(-1, &MyFancyStep{stage_id: 1})
	d.AddStep(-1, &MyFancyStep{stage_id: 2})

	d.Build()
	d.Run()
}
