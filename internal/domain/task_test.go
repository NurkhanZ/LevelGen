package domain

import (
	"errors"
	"testing"
)

func TestTaskTransitions(t *testing.T) {
	m := &TileMap{Width: 1, Height: 1, Grid: [][]TileType{{Ground}}}
	task := LevelTask{ID: "1", Status: StatusPending}
	if err := task.Complete(nil); err == nil { t.Fatal("nil result accepted") }
	if err := task.Complete(m); err != nil { t.Fatal(err) }
	if task.Status != StatusCompleted || task.Progress != 100 || task.Result != m || task.Error != "" { t.Fatalf("wrong completed task: %+v", task) }
	if err := task.Fail(errors.New("late")); err == nil { t.Fatal("completed task changed state") }
	failed := LevelTask{ID: "2", Status: StatusProcessing}
	if err := failed.Fail(errors.New("generation error")); err != nil { t.Fatal(err) }
	if failed.Status != StatusFailed || failed.Result != nil || failed.Error != "generation error" { t.Fatalf("wrong failed task: %+v", failed) }
}
