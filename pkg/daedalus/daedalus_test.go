package daedalus_test

import (
	"testing"

	"github.com/eightlay/daedalus/pkg/daedalus"
)

type TestStep struct{}

func (t *TestStep) Run(res *daedalus.Resolver) {}
func (t *TestStep) GetRequiredData() []string {
	return []string{}
}
func (t *TestStep) GetOutputData() []string {
	return []string{}
}

func TestNewDeadalus(t *testing.T) {
	d := daedalus.NewDaedalus()

	if d == nil {
		t.Error("NewDaedalus() returned nil")
	}
}

func TestAddStage(t *testing.T) {
	d := daedalus.NewDaedalus()
	d.AddStage()
}

func TestAddStep(t *testing.T) {
	d := daedalus.NewDaedalus()
	stage_id := d.AddStage()
	d.AddStep(stage_id, &TestStep{})
}

func TestAddSugarTest(t *testing.T) {
	d := daedalus.NewDaedalus()
	d.AddStep(-1, &TestStep{})
}

func TestCantAddNilStep(t *testing.T) {
	d := daedalus.NewDaedalus()

	defer func() {
		if r := recover(); r == nil {
			t.Error("AddStep() did not panic")
		}
	}()

	d.AddStep(0, nil)
}

func TestDelStage(t *testing.T) {
	d := daedalus.NewDaedalus()

	stage_id := d.AddStage()
	d.DelStage(stage_id)

	if d.GetStagesNumber() != 0 {
		t.Error("DelStage() did not delete stage")
	}
}

func TestDelUnexistingStage(t *testing.T) {
	d := daedalus.NewDaedalus()

	defer func() {
		if r := recover(); r == nil {
			t.Error("DelStage() did not panic")
		}
	}()

	d.DelStage(0)
}

func TestDelExactlyOne(t *testing.T) {
	d := daedalus.NewDaedalus()

	stage_id := d.AddStage()
	d.AddStage()
	d.DelStage(stage_id)

	if d.GetStagesNumber() == 0 {
		t.Error("DelStage() deleted all stages")
	}

	if d.GetStagesNumber() != 1 {
		t.Error("DelStage() did not delete stage")
	}
}

func TestDelStep(t *testing.T) {
	d := daedalus.NewDaedalus()

	stage_id, step_id := d.AddStep(-1, &TestStep{})
	d.DelStep(stage_id, step_id)

	if d.GetStageStepsNumber(stage_id) != 0 {
		t.Error("DelStep() did not delete step")
	}
}

func TestDelUnexistingStep(t *testing.T) {
	d := daedalus.NewDaedalus()
	stage_id := d.AddStage()

	defer func() {
		if r := recover(); r == nil {
			t.Error("DelStep() did not panic")
		}
	}()

	d.DelStep(stage_id, 0)
}

func TestDelExactlyOneStep(t *testing.T) {
	d := daedalus.NewDaedalus()

	stage_id, step_id := d.AddStep(-1, &TestStep{})
	d.AddStep(stage_id, &TestStep{})
	d.DelStep(stage_id, step_id)

	if d.GetStageStepsNumber(stage_id) == 0 {
		t.Error("DelStep() deleted all steps")
	}

	if d.GetStageStepsNumber(stage_id) != 1 {
		t.Error("DelStep() did not delete step")
	}
}

func TestClear(t *testing.T) {
	d := daedalus.NewDaedalus()

	d.AddStage()
	d.Clear()

	if d.GetStagesNumber() != 0 {
		t.Error("Clear() did not delete all stages")
	}
}

func TestClearStage(t *testing.T) {
	d := daedalus.NewDaedalus()

	stage_id, _ := d.AddStep(-1, &TestStep{})
	d.ClearStage(stage_id)

	if d.GetStageStepsNumber(stage_id) != 0 {
		t.Error("ClearStage() did not delete all steps")
	}
}

func TestClearUnexistingStage(t *testing.T) {
	d := daedalus.NewDaedalus()

	defer func() {
		if r := recover(); r == nil {
			t.Error("ClearStage() did not panic")
		}
	}()

	d.ClearStage(0)
}

func TestGetStagesNumber(t *testing.T) {
	d := daedalus.NewDaedalus()

	d.AddStage()
	d.AddStage()

	if d.GetStagesNumber() != 2 {
		t.Error("GetStagesNumber() returned wrong number of stages")
	}
}

func TestGetStageStepsNumber(t *testing.T) {
	d := daedalus.NewDaedalus()

	stage_id, _ := d.AddStep(-1, &TestStep{})
	d.AddStep(stage_id, &TestStep{})

	if d.GetStageStepsNumber(stage_id) != 2 {
		t.Error("GetStageStepsNumber() returned wrong number of steps")
	}
}

// TODO: Add tests for Run() and Build()
