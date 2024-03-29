package main

import (
	"fmt"
	"time"

	"github.com/eightlay/daedalus/pkg/daedalus"
)

type MyFancyStep struct {
	iters        int
	wait_message string
	wait_time    int
}

func (s *MyFancyStep) Run(data map[string]daedalus.Data) []daedalus.Data {
	for i := 0; i < s.iters; i++ {
		fmt.Println(s.wait_message)
		time.Sleep(time.Duration(s.wait_time) * time.Second)
	}
	return nil
}

func (s *MyFancyStep) GetRequiredData() []string {
	return []string{}
}

func (s *MyFancyStep) GetOutputData() []string {
	return []string{}
}

func main() {
	d := daedalus.NewDaedalus(daedalus.STEPS)
	stage_num := d.AddStage(true)
	d.AddStep(stage_num, &MyFancyStep{4, "step 1 waits 1 second", 1})
	d.AddStep(stage_num, &MyFancyStep{2, "step 2 waits 2 seconds", 2})

	d.Build()
	d.Run()
}
