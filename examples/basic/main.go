package main

import (
	"fmt"

	"github.com/eightlay/daedalus/pkg/daedalus"
)

type MyFancyStep struct {
}

func (s *MyFancyStep) Run(data map[string]daedalus.Data) []daedalus.Data {
	fmt.Println("MyFancyStep")
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
	d.Run()
}
