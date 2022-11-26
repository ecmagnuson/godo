package tests

import (
	"godo/cmd"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestNoPriority(t *testing.T) {
	var task string = cmd.Format("do something @home")
	wantedTask := "do something @home +p3"
	if task != wantedTask {
		t.Fatalf("failed to add priority 3 when none given.")
	}
}
