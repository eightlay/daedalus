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
	stage_num, _ := d.AddStep(-1, &MyFancyStep1{})
	d.AddStep(stage_num, &MyFancyStep2{})

	d.Build()
	d.Run()
}
