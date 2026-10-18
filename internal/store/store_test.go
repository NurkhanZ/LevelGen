package store

import (
	"errors"
	"testing"

	"github.com/NurkhanZ/LevelGen/internal/domain"
)

func TestMemoryStore(t *testing.T) {
	s := NewMemoryStore()
	a := domain.LevelTask{ID: "1", Status: domain.StatusPending}
	if err := s.Save(a); err != nil {
		t.Fatal(err)
	}
	if !errors.Is(s.Save(a), ErrDuplicate) {
		t.Fatal("duplicate ID accepted")
	}
	if err := s.Save(domain.LevelTask{ID: "2"}); err != nil {
		t.Fatal(err)
	}
	if list := s.List(); len(list) != 2 || list[0].ID != "1" || list[1].ID != "2" {
		t.Fatalf("wrong order: %+v", list)
	}
	m := &domain.TileMap{Width: 1, Height: 1, Grid: [][]domain.TileType{{domain.Start}}}
	if err := a.Complete(m); err != nil {
		t.Fatal(err)
	}
	if err := s.Update(a); err != nil {
		t.Fatal(err)
	}
	m.Grid[0][0] = domain.Exit
	got, err := s.GetByID("1")
	if err != nil || got.Status != domain.StatusCompleted || got.Result.Grid[0][0] != domain.Start {
		t.Fatalf("wrong stored task: %+v %v", got, err)
	}
	got.Result.Grid[0][0] = domain.Air
	again, _ := s.GetByID("1")
	if again.Result.Grid[0][0] != domain.Start {
		t.Fatal("store result was aliased")
	}
	if _, err := s.GetByID("missing"); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
}
