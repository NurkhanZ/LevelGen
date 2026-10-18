package generator

import (
	"testing"

	"github.com/NurkhanZ/LevelGen/internal/domain"
)

func TestGeneratedLevelsReachableAndDeterministic(t *testing.T) {
	g := PlatformGenerator{}
	for _, width := range []int{30, 36, 80} {
		for seed := -2; seed < 20; seed++ {
			p := domain.LevelParams{Width: width, Height: 30, Seed: seed}
			m, err := g.Generate(p)
			if err != nil {
				t.Fatalf("width %d seed %d: %v", width, seed, err)
			}
			if !Reachable(m) {
				t.Fatalf("unreachable width %d seed %d", width, seed)
			}
			again, err := g.Generate(p)
			if err != nil || again.String() != m.String() {
				t.Fatalf("nondeterministic width %d seed %d", width, seed)
			}
		}
	}
}

func TestSeveralVisibleStructures(t *testing.T) {
	for _, width := range []int{30, 50} {
		want := 3
		if width >= 36 {
			want = 4
		}
		for seed := 0; seed < 10; seed++ {
			m, err := (PlatformGenerator{}).Generate(domain.LevelParams{Width: width, Height: 30, Seed: seed})
			if err != nil {
				t.Fatal(err)
			}
			// Each stamped room has an air pocket above the connecting tunnel;
			// solid ground between pockets keeps the structures distinct.
			row := m.Grid[m.Height-4]
			rooms := 0
			inside := false
			for _, tile := range row {
				if tile == domain.Air && !inside {
					rooms++
					inside = true
				} else if tile == domain.Ground {
					inside = false
				}
			}
			if rooms != want {
				t.Fatalf("width %d seed %d: got %d visible rooms, want %d", width, seed, rooms, want)
			}
		}
	}
}

func TestSeedVariationAndFullHeight(t *testing.T) {
	g := PlatformGenerator{}
	seen := make(map[string]bool)
	for seed := 1; seed <= 5; seed++ {
		m, err := g.Generate(domain.LevelParams{Width: 40, Height: 30, Seed: seed})
		if err != nil {
			t.Fatal(err)
		}
		seen[m.String()] = true
		for section := 0; section < 3; section++ {
			openTiles := 0
			for y := section * 10; y < (section+1)*10; y++ {
				for _, tile := range m.Grid[y] {
					if tile == domain.Air || tile == domain.Start || tile == domain.Exit {
						openTiles++
					}
				}
			}
			if openTiles == 0 {
				t.Fatalf("seed %d: no carved spaces in section %d", seed, section)
			}
		}
		if m.Grid[6][1] != domain.Start || m.Grid[28][1] != domain.Exit {
			t.Fatalf("seed %d: Start or Exit is missing from its tier", seed)
		}
	}
	if len(seen) < 4 {
		t.Fatalf("five seeds produced only %d distinct maps", len(seen))
	}
}

func TestTallMapsRemainReachable(t *testing.T) {
	for _, height := range []int{31, 32, 40, 80, 400} {
		m, err := (PlatformGenerator{}).Generate(domain.LevelParams{Width: 30, Height: height, Seed: 23})
		if err != nil {
			t.Fatalf("height %d: %v", height, err)
		}
		if !Reachable(m) {
			t.Fatalf("height %d: no valid path", height)
		}
	}
}

func TestReachablePitWidth(t *testing.T) {
	makeMap := func(gap int) *domain.TileMap {
		m := &domain.TileMap{Width: 10, Height: 4, Grid: make([][]domain.TileType, 4)}
		for y := range m.Grid {
			m.Grid[y] = make([]domain.TileType, 10)
		}
		for x := range m.Grid[3] {
			m.Grid[3][x] = domain.Ground
		}
		for x := 3; x < 3+gap; x++ {
			m.Grid[3][x] = domain.Air
		}
		m.Grid[2][1], m.Grid[2][8] = domain.Start, domain.Exit
		return m
	}
	if !Reachable(makeMap(2)) {
		t.Fatal("two-tile pit should be jumpable")
	}
	if Reachable(makeMap(3)) {
		t.Fatal("three-tile pit should block the path")
	}
}
