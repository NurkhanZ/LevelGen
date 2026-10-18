package generator

import "github.com/NurkhanZ/LevelGen/internal/domain"

type position struct{ x, y int }

// Reachable checks player movement: up to one tile up, up to three tiles
// sideways, and an unbounded fall. A landing needs Ground or Platform below.
func Reachable(m *domain.TileMap) bool {
	if m == nil || m.Width == 0 || m.Height < 2 || len(m.Grid) != m.Height {
		return false
	}
	for _, row := range m.Grid {
		if len(row) != m.Width {
			return false
		}
	}
	var start position
	starts, exits := 0, 0
	for y, row := range m.Grid {
		for x, tile := range row {
			if tile == domain.Start {
				start, starts = position{x, y}, starts+1
			}
			if tile == domain.Exit {
				exits++
			}
		}
	}
	if starts != 1 || exits != 1 || !standing(m, start.x, start.y) {
		return false
	}
	queue := []position{start}
	seen := map[position]bool{start: true}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if m.Grid[p.y][p.x] == domain.Exit {
			return true
		}
		for dx := -3; dx <= 3; dx++ {
			for rise := 0; rise <= 1; rise++ {
				if dx == 0 && rise == 0 {
					continue
				}
				nx, ny := p.x+dx, p.y-rise
				if nx < 0 || nx >= m.Width || ny < 0 || !open(m.Grid[ny][nx]) {
					continue
				}
				// A jump needs head room across its horizontal span.
				clear := true
				step := 1
				if dx < 0 {
					step = -1
				}
				for cx := p.x; cx != nx; cx += step {
					if p.y > 0 && !open(m.Grid[p.y-1][cx]) {
						clear = false
						break
					}
				}
				if !clear {
					continue
				}
				for ny < m.Height-1 && open(m.Grid[ny][nx]) {
					if standing(m, nx, ny) {
						n := position{nx, ny}
						if !seen[n] {
							seen[n] = true
							queue = append(queue, n)
						}
						break
					}
					ny++
				}
			}
		}
	}
	return false
}

func open(t domain.TileType) bool { return t == domain.Air || t == domain.Start || t == domain.Exit }

func standing(m *domain.TileMap, x, y int) bool {
	if y < 0 || y+1 >= m.Height || !open(m.Grid[y][x]) {
		return false
	}
	below := m.Grid[y+1][x]
	return below == domain.Ground || below == domain.Platform
}
