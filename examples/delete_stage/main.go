package main

import (
	"fmt"

	"github.com/eightlay/daedalus/pkg/daedalus"
)

type MyFancyStep struct {
	stage_id int
}

func (s *MyFancyStep) Run(resolver *daedalus.Resolver) {
	fmt.Printf("MyFancyStep, stage_id == %d\n", s.stage_id)
}

func (s *MyFancyStep) GetRequiredData() []string {
	return []string{}
}

func (s *MyFancyStep) GetOutputData() []string {
	return []string{}
}

func main() {
	d := daedalus.NewDaedalus()
	stage_id, _ := d.AddStep(-1, &MyFancyStep{stage_id: 1})
	d.AddStep(-1, &MyFancyStep{stage_id: 2})
	d.DelStage(stage_id)

	d.Build()
	d.Run()
}
