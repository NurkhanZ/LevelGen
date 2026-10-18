# LevelGen | Tool for dynamic 2D platformer level generation

This project is a CLI program that allows to automatically generate unique and random 2D platformer levels.

## Setup and Run

- You can run this project by cloning this repository: <br>
`git clone: https://github.com/NurkhanZ/LevelGen.git`<br>
And then just run:<br> 
`go run .\cmd\level-gen\level-gen.go`.
    - Generated level map will be printed out in the console as ASCII typed Tile Map of default size:
        - width = 30 tiles, height = 30 tiles.

- Run `go run .\cmd\level-gen\level-gen.go -interactive` for an interactive shell.
    - In it, you can generate new levels, lookup previously generated ones and change the settings
    - For help, run `go run .\cmd\level-gen\level-gen.go -help` 

## Future work
Generator is <b>not</b> 100% implemented yet, as it lacks platforms, spikes, as well as robust generation and variety in structures.

These features are going to be added in next versions:
- Implement server for hosting games (multiplayer)
- Add support for moving objects, entities, traps
- More than one pathways for the player
- Account for robust platformer moveset (wall jumps, etc)
- Standalone application with clean level rendering and UI (no ascii)

Levels are meant to be compatible in the future, with the pygame 2D platformer, inspired by this incredible tutorial: https://www.youtube.com/watch?v=WViyCAa6yLI