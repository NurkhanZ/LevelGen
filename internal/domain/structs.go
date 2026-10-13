package domain

import (
	"fmt"
	"errors"
)

const (
	MAX_WIDTH, MIN_WIDTH   = 400, 30
	MAX_HEIGHT, MIN_HEIGHT = 400, 30
)

type LevelParams struct {
	Width           int
	Height          int
	PlatformDensity float64
	SpikeDensity    float64
	Seed            int
}

type TileMap struct {
	Width  int
	Height int
	Grid   [][]TileType
}

type LevelTask struct {	 // Status      Result   Error 		Progress
	ID 		 string		 // pending	    nil 	 empty 		0
	Params 	 LevelParams // processing 	nil 	 empty 		0-99 
	Status 	 TaskStatus  // completed  	non-nil  empty 		100
	Progress int		 // failed 		nil 	 string  	0-100
	Result   *TileMap
	Error 	 string
}

// shared sentinel errors
var ErrInvalidDimension = errors.New("Invalid dimension")
var ErrInvalidDensity = errors.New("Invalid density")
var ErrInvalidPosition = errors.New("Invalid tile position")
var ErrInvalidTileType = errors.New("Invalid TileType")

func (lp LevelParams) Validate() error {
	if !(lp.Width <= MAX_WIDTH && lp.Width >= MIN_WIDTH) {
		// fmt.Errorf returns error value with a string message. 
		return fmt.Errorf("Error: %w, width must be within max: %v, min: %v, given: %v", ErrInvalidDimension, MAX_WIDTH, MIN_WIDTH, lp.Width)
	}
	if !(lp.Height <= MAX_HEIGHT && lp.Height >= MIN_HEIGHT) {
		return fmt.Errorf("Error: %w, height must be within max: %v, min: %v, given: %v", ErrInvalidDimension, MAX_HEIGHT, MIN_HEIGHT, lp.Height)
	}
	if !(lp.PlatformDensity <= 1.0 && lp.PlatformDensity >= 0.0){
		return fmt.Errorf("Error: %w, platform density must be within max: 1.0, min: 0.0, given: %v", ErrInvalidDensity, lp.PlatformDensity)
	}
	if !(lp.SpikeDensity <= 1.0 && lp.SpikeDensity >= 0.0){
		return fmt.Errorf("Error: %w, spike density must be within max: 1.0, min: 0.0, given: %v", ErrInvalidDensity, lp.SpikeDensity)
	}
	return nil
}

func (tm TileMap) SetTile(x, y int, tt TileType) error {
	if (x > tm.Width - 1 || x < 0){
		return ErrInvalidPosition
	}
	if (y > tm.Height - 1 || y < 0) {
		return ErrInvalidPosition
	}
	if !tt.IsValid() {
		return ErrInvalidTileType
	}
	tm.Grid[y][x] = tt
	return nil
}