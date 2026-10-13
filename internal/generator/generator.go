package generator

import (
	"fmt"
	"github.com/NurkhanZ/LevelGen/internal/domain"
)

type PlatformGenerator struct {}

type Generator interface {
	Generate(params domain.LevelParams) (*domain.TileMap, error)
}

func (pg PlatformGenerator) Generate(params domain.LevelParams) (*domain.TileMap, error) {
	e := params.Validate();
	if e != nil {
		return nil, fmt.Errorf("Generation failed. %w", e)
	}
	// outer slice allocation
	grid := make([][]domain.TileType, params.Height)
	for y := range grid {
		// inner slice allocation, these are non shared arrays
		grid[y] = make([]domain.TileType, params.Width)
	}

	tm := domain.TileMap{
		Width: params.Width,
		Height: params.Height,
		Grid: grid,
	}
	// fmt.Print(tm)
	return &tm, e
}