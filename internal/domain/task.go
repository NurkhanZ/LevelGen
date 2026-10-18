package domain

import "fmt"

// Complete moves a task to its terminal success state.
func (t *LevelTask) Complete(result *TileMap) error {
	if t == nil || result == nil {
		return fmt.Errorf("cannot complete task without a result")
	}
	if t.Status != StatusPending && t.Status != StatusProcessing {
		return fmt.Errorf("cannot complete task in %q state", t.Status)
	}
	t.Status, t.Progress, t.Result, t.Error = StatusCompleted, 100, result, ""
	return nil
}

// Fail records a terminal generation error.
func (t *LevelTask) Fail(err error) error {
	if t == nil || err == nil {
		return fmt.Errorf("cannot fail task without an error")
	}
	if t.Status != StatusPending && t.Status != StatusProcessing {
		return fmt.Errorf("cannot fail task in %q state", t.Status)
	}
	t.Status, t.Result, t.Error = StatusFailed, nil, err.Error()
	return nil
}
