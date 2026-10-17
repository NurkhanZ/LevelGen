package domain

import (
	"fmt"
	"testing"
)

func TestTileMapString(t *testing.T){
	cases := []struct{
		in TileMap
		want string
	}{
		{
			TileMap{
				Width: 3,
				Height: 2,
				Grid: [][]TileType {
					{Air, Ground, Platform},
					{Spike, Start, Exit},
				},
			}, " #=\n^PE\n",
		},
		{
			TileMap{
				Width: 3,
				Height: 2,
				Grid: [][]TileType {
					{TileType(12), Ground, Platform},
					{Spike, Start, Exit},
				},
			}, "?#=\n^PE\n",
		},
	}

	for _, c := range cases{
		got := fmt.Sprint(c.in)
		if got != c.want{
			t.Errorf("String() == %q, want: %q", got, c.want)
		}
	}	
}