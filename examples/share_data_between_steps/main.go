package main

import (
	"fmt"

	"github.com/eightlay/daedalus/pkg/daedalus"
)

type MyFancyData struct {
	message string
}

func (d *MyFancyData) GetName() string {
	return "MyFancyData"
}

type MyFancyStep1 struct {
}

func (s *MyFancyStep1) Run(data map[string]daedalus.Data) []daedalus.Data {
	return []daedalus.Data{
		&MyFancyData{message: "Hello, I'm MyFancyStep1!"},
	}
}

func (s *MyFancyStep1) GetRequiredData() []string {
	return []string{}
}

func (s *MyFancyStep1) GetOutputData() []string {
	// return []string{"MyFancyData"}
	return []string{(&MyFancyData{}).GetName()}
}

type MyFancyStep2 struct {
}

func (s *MyFancyStep2) Run(data map[string]daedalus.Data) []daedalus.Data {
	my_fancy_data := data["MyFancyData"].(*MyFancyData)
	fmt.Println("Message from MyFancyStep1: ", my_fancy_data.message)
	return nil
}

func (s *MyFancyStep2) GetRequiredData() []string {
	// return []string{"MyFancyData"}
	return []string{(&MyFancyData{}).GetName()}
}

func (s *MyFancyStep2) GetOutputData() []string {
	return []string{}
}

func main() {
	d := daedalus.NewDaedalus(daedalus.STEPS)
	stage_num, _ := d.AddStep(-1, &MyFancyStep1{})
	d.AddStep(stage_num, &MyFancyStep2{})

	d.Build()
	d.Run()

	d.PrintRunStats()
}
