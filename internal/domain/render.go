package domain

import "strings"

func tileToChar(tt TileType) byte {
	var c byte
	switch tt {
	case Air:
		c = ' '
	case Ground:
		c = '#'
	case Platform:
		c = '='
	case Spike:
		c = '^'
	case Start:
		c = 'P'
	case Exit:
		c = 'E'
	default:
		c = '?'
	}
	return c
}

func (tm TileMap) String() string {
	var sb strings.Builder

	for y := range tm.Grid{
		for _, tile := range tm.Grid[y]{
			tc := tileToChar(tile) // translate each tile to character
			sb.WriteByte(tc) // write to string builder byte buffer
		}
		sb.WriteByte('\n') // at the end of each row, append newline
	}

	return sb.String()
}