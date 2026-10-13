package domain

import (
	"errors"
	"testing"
)

func TestLevelParamsn(t *testing.T){
	cases := []struct {
		in LevelParams
		want error
	}{
		{
			LevelParams{
				Width: 30,
				Height: 30,
				PlatformDensity: 0.2,
				SpikeDensity: 0.1,
				Seed: 12,
			}, nil,
		},
		{
			LevelParams{
				Width: 401,
				Height: 30,
				PlatformDensity: 0.2,
				SpikeDensity: 0.1,
				Seed: 12,
			}, ErrInvalidDimension,
		},
		{
			LevelParams{
				Width: 400,
				Height: 401,
				PlatformDensity: 0.2,
				SpikeDensity: 0.1,
				Seed: 12,
			}, ErrInvalidDimension,
		},
		{
			LevelParams{
				Width: 40,
				Height: 40,
				PlatformDensity: 1.1,
				SpikeDensity: 0.1,
				Seed: 12,
			}, ErrInvalidDensity,
		},
		{
			LevelParams{
				Width: 40,
				Height: 40,
				PlatformDensity: 1.0,
				SpikeDensity: 1.1,
				Seed: 12,
			}, ErrInvalidDensity,
		},
		{
			LevelParams{
				Width: 40,
				Height: 40,
				PlatformDensity: 1.0,
				SpikeDensity: -0.1,
				Seed: 12,
			}, ErrInvalidDensity,
		},
	}

	for _, c := range cases {
		got := c.in.Validate()
		if !errors.Is(got, c.want) {
			t.Errorf("Validate() == %q, want: %q", got, c.want)
		}
	}
}

func TestSetTitle(t *testing.T){
	cases := []struct {
		x, y int
		tile TileType
		want bool
	}{
		{0, 0, Air, false},
		{1, 2, Air, true},
		{1, 1, Air, false},
		{-1, 1, Air, true},
		{1, -1, Air, true},
		{3, 2, Air, true},
		{1, 1, TileType(99), true},
		{2, 1, TileType(3), false},
	}

	for _, c := range cases{
		tm := TileMap{
			Width: 3,
			Height: 2,
			Grid: [][]TileType {
				{Air, Air, Air},
				{Air, Air, Air},
			},
		}

		got := tm.SetTile(c.x, c.y, c.tile)
		if (got != nil) != c.want {
			t.Errorf("SetTitle(%v, %v, %v) == %v, want: %v", c.x, c.y, c.tile, got, c.want)
		}
	}
}