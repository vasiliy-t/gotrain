package main

import (
	"testing"
)

func TestAddOnTermHook_AddsHook(t *testing.T) {
	termHandlers = []TermHook{}

	AddOnTermHook(func() {
		//example hook
	})

	if len(termHandlers) != 1 {
		t.Fatalf("AddOnTermHook does not adds hooks")
	}
}

func TestOnTerm_ExecutesRegisteredTermHandlers(t *testing.T) {
	termHandlers = []TermHook{}

	fired := false
	AddOnTermHook(func() {
		fired = true
	})

	OnTerm()

	if fired != true {
		t.Fatalf("OnTerm does not executes registered handlers")
	}
}
