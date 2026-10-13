package domain

import (
	"testing"
)
func TestTileType(t *testing.T){
	cases := []struct {
		in TileType
		want bool
	}{
		{Air, true},
		{Ground, true},
		{Platform, true},
		{Spike, true},
		{Start, true},
		{Exit, true},
		{TileType(42), false},
		{TileType(-1), false},
		{TileType(0), true},
	}
	for _, c := range cases {
		got := c.in.IsValid()
		if got != c.want {
			t.Errorf("IsValid(%v) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestTaskStatus(t *testing.T) {
	cases := []struct {
		in TaskStatus
		want string
	}{
		{StatusPending, "pending"},
		{StatusProcessing, "processing"},
		{StatusCompleted, "completed"},
		{StatusFailed, "failed"},
	}
	for _, c := range cases {
		if string(c.in) != c.want {
			t.Errorf("What status entered: %q, want %q", c.in, c.want)
		}
	}
}