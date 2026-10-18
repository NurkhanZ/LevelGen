package generator

import (
	"fmt"
	"math/rand"

	"github.com/NurkhanZ/LevelGen/internal/domain"
)

// Generator is the generation boundary used by the CLI.
type Generator interface {
	Generate(domain.LevelParams) (*domain.TileMap, error)
}

type PlatformGenerator struct{}

// Generate carves seeded rooms and tunnels across the map. Each tier has a
// walkable route; alternating shafts connect tiers by a controlled fall.
func (PlatformGenerator) Generate(params domain.LevelParams) (*domain.TileMap, error) {
	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("generation failed: %w", err)
	}
	m := &domain.TileMap{Width: params.Width, Height: params.Height, Grid: make([][]domain.TileType, params.Height)}
	for y := range m.Grid {
		m.Grid[y] = make([]domain.TileType, params.Width)
		for x := range m.Grid[y] {
			m.Grid[y][x] = domain.Ground
		}
	}
	rng := rand.New(rand.NewSource(int64(params.Seed)))
	floors := tierFloors(params.Height, rng)
	count := 3
	if params.Width >= 36 {
		count = 4
	}
	pits := make([]int, len(floors))
	for i, floor := range floors {
		pits[i] = stampTier(m, floor, count, rng)
	}
	// Carve after every template is stamped, so adjacent tiers cannot close a
	// previously made passage. Leave Ground below each standing position.
	for i, floor := range floors {
		for x := 1; x < params.Width-1; x++ {
			m.Grid[floor-2][x] = domain.Air
			m.Grid[floor-1][x] = domain.Air
			m.Grid[floor][x] = domain.Ground
		}
		m.Grid[floor][pits[i]] = domain.Air
		m.Grid[floor][pits[i]+1] = domain.Air
	}
	for i := 0; i < len(floors)-1; i++ {
		shaftX := 3
		if i%2 == 0 {
			shaftX = params.Width - 5
		}
		for y := floors[i]; y < floors[i+1]; y++ {
			m.Grid[y][shaftX] = domain.Air
			m.Grid[y][shaftX+1] = domain.Air
		}
	}
	m.Grid[floors[0]-1][1] = domain.Start
	exitX := params.Width - 2
	if (len(floors)-1)%2 == 1 {
		exitX = 1
	}
	m.Grid[floors[len(floors)-1]-1][exitX] = domain.Exit
	if !Reachable(m) {
		return nil, fmt.Errorf("generation failed: no path from Start to Exit")
	}
	return m, nil
}

func tierFloors(height int, rng *rand.Rand) []int {
	const first = 7 // Room templates are at most seven tiles high.
	distance := height - 1 - first
	intervals := (distance + 8) / 9
	steps := make([]int, intervals)
	for i := range steps {
		steps[i] = distance / intervals
	}
	for i := 0; i < distance%intervals; i++ {
		steps[i]++
	}
	rng.Shuffle(len(steps), func(i, j int) { steps[i], steps[j] = steps[j], steps[i] })
	floors := []int{first}
	for _, step := range steps {
		floors = append(floors, floors[len(floors)-1]+step)
	}
	return floors
}

// stampTier uses several predefined structures and distributes the spare
// width randomly among the connectors. The pit remains away from both shafts.
func stampTier(m *domain.TileMap, floor, count int, rng *rand.Rand) int {
	chosen := []Structure{structures[0], structures[6]}
	variants := []Structure{structures[3], structures[4], structures[7]}
	order := rng.Perm(len(variants))
	for i := 0; len(chosen) < count; i++ {
		chosen = append(chosen, variants[order[i]])
	}
	total := 0
	for _, s := range chosen {
		total += len(s.Grid[0])
	}
	gaps := make([]int, count-1)
	for i := range gaps {
		gaps[i] = 2
	}
	spare := m.Width - 4 - total - 2*len(gaps)
	for i := 0; i < spare; i++ {
		gaps[rng.Intn(len(gaps))]++
	}
	x, pitX := 2, -1
	for i, s := range chosen {
		top := floor - len(s.Grid)
		for sy, row := range s.Grid {
			copy(m.Grid[top+sy][x:], row)
		}
		if s.Name == "Pit" {
			pitX = x + 1
		}
		if i < len(gaps) {
			x += len(s.Grid[0]) + gaps[i]
		}
	}
	return pitX
}
