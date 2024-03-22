package idcounter_test

import (
	"testing"

	idcounter "github.com/eightlay/daedalus/internal/id_counter"
)

func TestNewIdCounter(t *testing.T) {
	counter := idcounter.NewIdCounter()

	if counter == nil {
		t.Error("NewIdCounter() returned nil")
	}

	if counter.Current() != -1 {
		t.Error("NewIdCounter() did not initialize counter to -1")
	}
}

func TestNext(t *testing.T) {
	counter := idcounter.NewIdCounter()

	if counter.Next() != 0 {
		t.Error("First call to Next() should return 0")
	}

	if counter.Next() != 1 {
		t.Error("Second call to Next() should return 1")
	}
}

func TestCurrent(t *testing.T) {
	counter := idcounter.NewIdCounter()

	if counter.Current() != -1 {
		t.Error("Current() should return -1 initially")
	}

	counter.Next()

	if counter.Current() != 0 {
		t.Error("Current() should return 0 after one call to Next()")
	}
}

func TestClear(t *testing.T) {
	counter := idcounter.NewIdCounter()

	counter.Next()
	counter.Clear()

	if counter.Current() != -1 {
		t.Error("Clear() did not reset counter to -1")
	}
}
