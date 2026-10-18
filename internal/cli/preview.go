package cli

import (
	"fmt"
	"strings"

	"github.com/NurkhanZ/LevelGen/internal/domain"
)

// Preview renders a bordered interactive view. TileMap.String remains plain ASCII.
func Preview(m *domain.TileMap, p domain.LevelParams) (string, error) {
	if m == nil || m.Width < 1 || m.Height < 1 || len(m.Grid) != m.Height {
		return "", fmt.Errorf("invalid map dimensions")
	}
	for _, row := range m.Grid {
		if len(row) != m.Width {
			return "", fmt.Errorf("irregular map grid")
		}
	}
	rows := strings.Split(strings.TrimSuffix(m.String(), "\n"), "\n")
	var b strings.Builder
	fmt.Fprintf(&b, "Map %dx%d  seed %d\n", m.Width, m.Height, p.Seed)
	border := "+" + strings.Repeat("-", m.Width) + "+\n"
	b.WriteString(border)
	for _, row := range rows {
		fmt.Fprintf(&b, "|%s|\n", row)
	}
	b.WriteString(border)
	b.WriteString("P Start  E Exit  # Ground  = Platform  ^ Spike\n")
	if m.Width > 100 {
		b.WriteString("Wide map: terminal lines may wrap.\n")
	}
	return b.String(), nil
}
