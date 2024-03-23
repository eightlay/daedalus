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

func (d *MyFancyData) CopyFrom(data daedalus.Data) {
	if conv, ok := data.(*MyFancyData); ok {
		d.message = conv.message
	} else {
		panic("Trying to copy data from different types")
	}
}

type MyFancyStep1 struct {
}

func (s *MyFancyStep1) Run(resolver *daedalus.Resolver) {
	resolver.PushData(&MyFancyData{message: "Hello, I'm MyFancyStep1!"})
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

func (s *MyFancyStep2) Run(resolver *daedalus.Resolver) {
	data := MyFancyData{}
	resolver.GetData(&data)

	fmt.Println("Message from MyFancyStep1: ", data.message)
}

func (s *MyFancyStep2) GetRequiredData() []string {
	// return []string{"MyFancyData"}
	return []string{(&MyFancyData{}).GetName()}
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
