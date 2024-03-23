package main

import (
	"fmt"

	"github.com/eightlay/daedalus/pkg/daedalus"
)

type MyFancyStep struct {
}

func (s *MyFancyStep) Run(resolver *daedalus.Resolver) {
	fmt.Println("MyFancyStep")
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
	d.Run()
}
