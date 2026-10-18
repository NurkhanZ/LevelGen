package store

import (
	"errors"
	"fmt"

	"github.com/NurkhanZ/LevelGen/internal/domain"
)

var ErrNotFound = errors.New("task not found")
var ErrDuplicate = errors.New("task ID already exists")

type TaskStore interface {
	Save(domain.LevelTask) error
	Update(domain.LevelTask) error
	GetByID(string) (domain.LevelTask, error)
	List() []domain.LevelTask
}

// MemoryStore keeps tasks in insertion order for stable menu numbering.
type MemoryStore struct {
	tasks map[string]domain.LevelTask
	order []string
}

func NewMemoryStore() *MemoryStore { 
	return &MemoryStore{
		tasks: make(map[string]domain.LevelTask),
	} 
}

func (s *MemoryStore) Save(task domain.LevelTask) error {
	if s == nil || task.ID == "" {
		return fmt.Errorf("task ID is required")
	}
	if _, ok := s.tasks[task.ID]; ok {
		return ErrDuplicate
	}
	s.tasks[task.ID] = clone(task)
	s.order = append(s.order, task.ID)
	return nil
}

func (s *MemoryStore) Update(task domain.LevelTask) error {
	if s == nil {
		return ErrNotFound
	}
	if _, ok := s.tasks[task.ID]; !ok {
		return ErrNotFound
	}
	s.tasks[task.ID] = clone(task)
	return nil
}

func (s *MemoryStore) GetByID(id string) (domain.LevelTask, error) {
	if s == nil {
		return domain.LevelTask{}, ErrNotFound
	}
	task, ok := s.tasks[id]
	if !ok {
		return domain.LevelTask{}, ErrNotFound
	}
	return clone(task), nil
}

func (s *MemoryStore) List() []domain.LevelTask {
	if s == nil {
		return nil
	}
	out := make([]domain.LevelTask, 0, len(s.order))
	for _, id := range s.order {
		out = append(out, clone(s.tasks[id]))
	}
	return out
}

func clone(task domain.LevelTask) domain.LevelTask {
	if task.Result != nil {
		m := *task.Result
		m.Grid = make([][]domain.TileType, len(task.Result.Grid))
		for y, row := range task.Result.Grid {
			m.Grid[y] = append([]domain.TileType(nil), row...)
		}
		task.Result = &m
	}
	return task
}
