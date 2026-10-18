package generator

import (
	"errors"
	"fmt"
	"github.com/NurkhanZ/LevelGen/internal/domain"
	"testing"
)

// All tests have same non-nil error -> nil TileMap check

func TestGenerateParams(t *testing.T) {
	cases := []struct {
		in   domain.LevelParams
		want error
	}{
		{
			domain.LevelParams{
				Width:           30,
				Height:          30,
				PlatformDensity: 0.2,
				SpikeDensity:    0.1,
				Seed:            12,
			}, nil,
		},
		{
			domain.LevelParams{
				Width:           401,
				Height:          30,
				PlatformDensity: 0.2,
				SpikeDensity:    0.1,
				Seed:            12,
			}, domain.ErrInvalidDimension,
		},
		{
			domain.LevelParams{
				Width:           400,
				Height:          401,
				PlatformDensity: 0.2,
				SpikeDensity:    0.1,
				Seed:            12,
			}, domain.ErrInvalidDimension,
		},
		{
			domain.LevelParams{
				Width:           40,
				Height:          40,
				PlatformDensity: 1.1,
				SpikeDensity:    0.1,
				Seed:            12,
			}, domain.ErrInvalidDensity,
		},
		{
			domain.LevelParams{
				Width:           40,
				Height:          40,
				PlatformDensity: 1.0,
				SpikeDensity:    1.1,
				Seed:            12,
			}, domain.ErrInvalidDensity,
		},
		{
			domain.LevelParams{
				Width:           40,
				Height:          40,
				PlatformDensity: 1.0,
				SpikeDensity:    -0.1,
				Seed:            12,
			}, domain.ErrInvalidDensity,
		},
		{
			domain.LevelParams{
				Width:           3,
				Height:          2,
				PlatformDensity: 0.1,
				SpikeDensity:    0.01,
				Seed:            123,
			}, domain.ErrInvalidDimension,
		},
		{
			domain.LevelParams{
				Width:           50,
				Height:          40,
				PlatformDensity: 0.0001,
				SpikeDensity:    0.1,
				Seed:            193874,
			}, nil,
		},
	}

	for _, c := range cases {
		err_prefix := fmt.Sprintf("Generate(%+v) failed", c.in)

		pg := PlatformGenerator{}

		got_map, got_err := pg.Generate(c.in)

		// invalid case expectation
		if c.want != nil {
			if got_err == nil {
				t.Errorf("%s: invalid error, got nil, want %q", err_prefix, c.want)
			}
			if !errors.Is(got_err, c.want) {
				t.Errorf("%s: unexpected error %q, want %q", err_prefix, got_err, c.want)
			}
			if got_map != nil {
				t.Errorf("%s: invalid TileMap init, got %+v, want nil", err_prefix, got_map.Grid[:5]) // safe, because we firstly validate in Generate()
			}
			// skip other checks
			return
		}

		// valid case expectation
		if got_map.Width != c.in.Width {
			t.Errorf("%s: invalid map width, got %d, want %d", err_prefix, got_map.Width, c.in.Width)
		}
		if got_map.Height != c.in.Height {
			t.Errorf("%s: invalid map height, got %d, want %d", err_prefix, got_map.Height, c.in.Height)
		}
		if len(got_map.Grid) != c.in.Height { // row count check
			t.Errorf("%s: invalid grid row number, got %d, want %d", err_prefix, len(got_map.Grid), c.in.Height)
		}
		for y, x := range got_map.Grid { // row length check
			if len(x) != c.in.Width {
				t.Errorf("%s: invalid length of grid row %d, got %d, want %d", err_prefix, y, len(x), c.in.Width)
			}
		}
	}
}

// Test for correct Ground generation
func TestGenerateGround(t *testing.T) {
	cases := []struct {
		in   domain.LevelParams
		want error
	}{
		{
			domain.LevelParams{
				Width:           30,
				Height:          30,
				PlatformDensity: 0.2,
				SpikeDensity:    0.1,
				Seed:            12,
			}, nil,
		},
		{
			domain.LevelParams{
				Width:           401,
				Height:          30,
				PlatformDensity: 0.2,
				SpikeDensity:    0.1,
				Seed:            12,
			}, domain.ErrInvalidDimension,
		},
	}

	for _, c := range cases {
		err_prefix := fmt.Sprintf("Generate(%+v) failed", c.in)

		pg := PlatformGenerator{}

		got_map, got_err := pg.Generate(c.in)

		if c.want != nil {
			if got_err == nil {
				t.Errorf("%s: invalid error, got nil, want %q", err_prefix, c.want)
			}
			if !errors.Is(got_err, c.want) {
				t.Errorf("%s: unexpected error %q, want %q", err_prefix, got_err, c.want)
			}
			if got_map != nil {
				t.Errorf("%s: invalid TileMap init, got %+v, want nil", err_prefix, got_map.Grid[:5])
			}
			return
		}

		bottom_row := got_map.Grid[c.in.Height-1]
		pitTiles := 0
		for _, tile := range bottom_row {
			if tile == domain.Air {
				pitTiles++
			} else if tile != domain.Ground {
				t.Fatalf("%s: unexpected floor tile %v", err_prefix, tile)
			}
		}
		if pitTiles != 2 {
			t.Fatalf("%s: got %d pit tiles, want 2", err_prefix, pitTiles)
		}
	}
}
