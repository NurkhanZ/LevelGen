package generator

import (
	"testing"
	"errors"
	"github.com/NurkhanZ/LevelGen/internal/domain"
)

func TestGenerate(t *testing.T){
	cases := []struct {
		in domain.LevelParams
		want error
	}{
		{
			domain.LevelParams{
				Width: 30,
				Height: 30,
				PlatformDensity: 0.2,
				SpikeDensity: 0.1,
				Seed: 12,
			}, nil,
		},
		{
			domain.LevelParams{
				Width: 401,
				Height: 30,
				PlatformDensity: 0.2,
				SpikeDensity: 0.1,
				Seed: 12,
			}, domain.ErrInvalidDimension,
		},
		{
			domain.LevelParams{
				Width: 400,
				Height: 401,
				PlatformDensity: 0.2,
				SpikeDensity: 0.1,
				Seed: 12,
			}, domain.ErrInvalidDimension,
		},
		{
			domain.LevelParams{
				Width: 40,
				Height: 40,
				PlatformDensity: 1.1,
				SpikeDensity: 0.1,
				Seed: 12,
			}, domain.ErrInvalidDensity,
		},
		{
			domain.LevelParams{
				Width: 40,
				Height: 40,
				PlatformDensity: 1.0,
				SpikeDensity: 1.1,
				Seed: 12,
			}, domain.ErrInvalidDensity,
		},
		{
			domain.LevelParams{
				Width: 40,
				Height: 40,
				PlatformDensity: 1.0,
				SpikeDensity: -0.1,
				Seed: 12,
			}, domain.ErrInvalidDensity,
		},
		{
			domain.LevelParams{
				Width: 3,
				Height: 2,
				PlatformDensity: 0.1,
				SpikeDensity: 0.01,
				Seed: 123,
			}, domain.ErrInvalidDimension,
		},
		{
			domain.LevelParams{
				Width: 50,
				Height: 40,
				PlatformDensity: 0.0001,
				SpikeDensity: 0.1,
				Seed: 193874,
			}, nil,
		},
	}

	for _, c := range cases {
		pg := PlatformGenerator{}

		got_map, got_err := pg.Generate(c.in)

		if !errors.Is(got_err, c.want) {
			t.Fatalf("Generate(%v) returned error %q, want: %q", c.in, got_err, c.want)
		} else {
			// nil tile map expectation, if error is present
			if got_map == nil {
				t.Fatalf("Generate(%v) returned tile map %v", c.in, got_map)
			}
			return
		}		

		if got_map.Width != c.in.Width {
			t.Errorf("Generate() returned invalid map width, got %d, want %d", got_map.Width, c.in.Width)
		} 
		if got_map.Height != c.in.Height {
			t.Errorf("Generate() returned invalid map height, got %d, want %d", got_map.Height, c.in.Height)
		}
		
		if len(got_map.Grid) != c.in.Height { // row count basically
			t.Errorf("Generate() returned invalid grid row number, got %d, want %d", len(got_map.Grid), c.in.Height)
		}

		for y, x := range got_map.Grid { // row length
			if len(x) != c.in.Width {
				t.Errorf("Generate() returned invalid grid row %d length, got %d, want %d", y, len(x), c.in.Width)
			}
		}

	}
}