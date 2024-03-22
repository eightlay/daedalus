package main

import (
	"fmt"

	"github.com/eightlay/daedalus/pkg/daedalus"
)

type MyFancyStep1 struct {
}

func (s *MyFancyStep1) Run(resolver *daedalus.Resolver) {
	fmt.Println("MyFancyStep1")
}

func (s *MyFancyStep1) GetRequiredData() []string {
	return []string{}
}

func (s *MyFancyStep1) GetOutputData() []string {
	return []string{}
}

type MyFancyStep2 struct {
}

func (s *MyFancyStep2) Run(resolver *daedalus.Resolver) {
	fmt.Println("MyFancyStep2")
}

func (s *MyFancyStep2) GetRequiredData() []string {
	return []string{}
}

func (s *MyFancyStep2) GetOutputData() []string {
	return []string{}
}

func main() {
	d := daedalus.NewDaedalus()
	stage_ind, step_ind := d.AddStep(-1, &MyFancyStep1{})

	d.AddStep(stage_ind, &MyFancyStep2{})
	d.DelStep(stage_ind, step_ind)

	d.AddStep(-1, &MyFancyStep1{})

	if err := d.Build(); err != nil {
		panic(err)
	}

	d.Run()
}
