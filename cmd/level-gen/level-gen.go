package main

import (
	"flag"
	"fmt"

	"github.com/NurkhanZ/LevelGen/internal/domain"
	"github.com/NurkhanZ/LevelGen/internal/generator"
)

func main(){
	
	// command line flags
	width_flag := flag.Int("width", 30, "sets level's width")
	height_flag := flag.Int("height", 30, "sets level's height")
	platform_density_flag := flag.Float64("platform-density", 0.0, "sets generation density of platforms")
	spike_density_flag := flag.Float64("spike-density", 0.0, "sets generation density of spikes")
	seed_flag := flag.Int("seed", 1, "sets seed for random level generation")
	flag.Parse()
	
	// set level parameters
	level_params := domain.LevelParams{
		Width: *width_flag,
		Height: *height_flag,
		PlatformDensity: *platform_density_flag,
		SpikeDensity: *spike_density_flag,
		Seed: *seed_flag,
	}

	// generate level
	platform_generator := generator.PlatformGenerator{}
	tile_map, err := platform_generator.Generate(level_params)
	if err != nil {
		panic(err)
	}
	if tile_map != nil {
		fmt.Print(tile_map)
	}
}