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
	// validate generator inputs
	e := params.Validate();
	if e != nil {
		return nil, fmt.Errorf("Generation failed. %w", e)
	}

	// create TileMap of size Width*Height, filled with Air
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

	// generate a floor, for every column, at the last row.
	bottom_row := tm.Grid[tm.Height-1]
	for x := range bottom_row{ 
		bottom_row[x] = domain.Ground
	}

	return &tm, e
}