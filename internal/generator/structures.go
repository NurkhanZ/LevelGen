package generator

import "github.com/NurkhanZ/LevelGen/internal/domain"

type Structure struct {
	Name string
	Grid [][]domain.TileType
	Ports []Port
}

type Port struct {
	X int
	Y int
	Side PortSide
}

type PortSide uint8

const (
	PortNorth PortSide = iota
	PortEast
	PortSouth
	PortWest
)

var structures = []Structure {
	{
		Name: "Safe room",
		Grid: [][]domain.TileType{									      // G = domain.Ground, A = domain.Air, P = generator.Port
			{domain.Ground, domain.Ground, domain.Ground, domain.Ground}, // G G G G
			{domain.Ground, domain.Air,    domain.Air, 	  domain.Ground}, // G A A G 
			{domain.Ground, domain.Air,    domain.Air,    domain.Air},    // G A A AP
			{domain.Ground, domain.Ground, domain.Ground, domain.Ground}, // G G G G
		},																  
		Ports: []Port{
			{X: 3, Y: 2, Side: PortEast},
		},
	},
	{
		Name: "Tunnel",
		Grid: [][]domain.TileType{									      
			{domain.Ground, domain.Ground, domain.Ground, domain.Ground}, // G G G G
			{domain.Air, 	domain.Air,    domain.Air, 	  domain.Air},    //PA A A AP 
			{domain.Air, 	domain.Air,    domain.Air,    domain.Air},    //PA A A AP
			{domain.Ground, domain.Ground, domain.Ground, domain.Ground}, // G G G G
		},
		Ports: []Port{
			{X: 0, Y: 1, Side: PortWest},
			{X: 3, Y: 1, Side: PortEast},
			{X: 0, Y: 2, Side: PortWest},
			{X: 3, Y: 2, Side: PortEast},
		},
	},
	{
		Name: "Small tunnel",
		Grid: [][]domain.TileType{									      
			{domain.Ground, domain.Ground, domain.Ground, domain.Ground}, // G G G G
			{domain.Air, 	domain.Air,    domain.Air, 	  domain.Air},    //PA A A AP
			{domain.Ground, domain.Ground, domain.Ground, domain.Ground}, // G G G G
		},
		Ports: []Port{
			{X: 0, Y: 1, Side: PortWest},
			{X: 3, Y: 1, Side: PortEast},
		},
	},
	{
		Name: "Stairs up",
		Grid: [][]domain.TileType{										    
			{domain.Air, 	domain.Air, 	domain.Air, 	domain.Air}, 	// A A A AP
			{domain.Air, 	domain.Air, 	domain.Air, 	domain.Ground}, // A A A G
			{domain.Air, 	domain.Air, 	domain.Ground, 	domain.Ground}, // A A G G 
			{domain.Air, 	domain.Ground, 	domain.Ground, 	domain.Ground}, //PA G G G
			{domain.Ground, domain.Ground, 	domain.Ground, 	domain.Ground}, // G G G G
		},
		Ports: []Port{
			{X: 3, Y: 0, Side: PortEast},
			{X: 0, Y: 3, Side: PortWest},
		},
	},
	{
		Name: "Stairs down",
		Grid: [][]domain.TileType{
			{domain.Air, 	domain.Air, 	domain.Air, 	domain.Air},    //PA A A A
			{domain.Ground, domain.Air, 	domain.Air, 	domain.Air}, 	// G A A A
			{domain.Ground, domain.Ground, 	domain.Air, 	domain.Air}, 	// G G A A 
			{domain.Ground, domain.Ground, 	domain.Ground, 	domain.Air}, 	// G G G AP
			{domain.Ground, domain.Ground, 	domain.Ground, 	domain.Ground}, // G G G G
		},
		Ports: []Port{
			{X: 0, Y: 0, Side: PortWest},
			{X: 3, Y: 3, Side: PortEast},
		},
	},
	{
		Name: "Drop down",
		Grid: [][]domain.TileType{							//   P  Port here is not outside of the bounds, its on the North side of Ground
			{domain.Ground, domain.Air, 	domain.Ground}, // G A G
			{domain.Ground, domain.Air, 	domain.Ground}, // G A G 
			{domain.Ground, domain.Air, 	domain.Ground}, // G A G 
			{domain.Ground, domain.Air, 	domain.Ground}, // G A G
		},													//   P
		Ports: []Port{
			{X: 1, Y: 0, Side: PortNorth},
			{X: 1, Y: 3, Side: PortSouth},
		},
	},
	{
		Name: "Pit",
		Grid: [][]domain.TileType{
			{domain.Air, 	domain.Air, 	domain.Air, 	domain.Air}, 	// A A A A			
			{domain.Air, 	domain.Air, 	domain.Air, 	domain.Air}, 	//PA A A AP
			{domain.Ground, domain.Air, 	domain.Air, 	domain.Ground}, // G A A G 
			{domain.Ground, domain.Air, 	domain.Air, 	domain.Ground}, // G A A G 
			{domain.Ground, domain.Ground, 	domain.Ground, 	domain.Ground}, // G G G G
		},
		Ports: []Port{
			{X: 0, Y: 1, Side: PortWest},
			{X: 3, Y: 1, Side: PortEast},
		},
	},
	{
		Name: "Chamber",
		Grid: [][]domain.TileType{
			{domain.Ground,	domain.Ground, 	domain.Ground, 	domain.Ground, 	domain.Ground, 	domain.Ground}, // G G G G G G 
			{domain.Ground, domain.Air, 	domain.Air, 	domain.Air, 	domain.Air, 	domain.Ground},	// G A A A A G
			{domain.Ground, domain.Air, 	domain.Air, 	domain.Air, 	domain.Air, 	domain.Ground}, // G A A A A G
			{domain.Air, 	domain.Air,		domain.Air, 	domain.Air, 	domain.Air, 	domain.Ground}, //PA A A A A G
			{domain.Ground, domain.Air,		domain.Air, 	domain.Air, 	domain.Air, 	domain.Air}, 	// G A A A A AP
			{domain.Ground, domain.Air,		domain.Air, 	domain.Air, 	domain.Air, 	domain.Ground}, // G A A A A G
			{domain.Ground, domain.Ground, 	domain.Ground, 	domain.Ground, 	domain.Ground, 	domain.Ground}, // G G G G G G
		},
		Ports: []Port{
			{X: 0, Y: 3, Side: PortWest},
			{X: 5, Y: 4, Side: PortEast},
		},
	},
	{
		Name: "Branch",
		Grid: [][]domain.TileType{																			//   P
			{domain.Ground, domain.Air, 	domain.Ground, 	domain.Ground, 	domain.Ground,	domain.Ground}, // G A G G G G
			{domain.Ground,	domain.Air, 	domain.Ground, 	domain.Ground, 	domain.Air, 	domain.Air},	// G A G G A AP
			{domain.Ground, domain.Air,		domain.Air, 	domain.Air, 	domain.Air, 	domain.Ground}, // G A A A A G
			{domain.Ground, domain.Air, 	domain.Air, 	domain.Air,		domain.Ground, 	domain.Ground}, // G A A A G G
			{domain.Ground, domain.Ground, 	domain.Ground,  domain.Ground, 	domain.Ground, 	domain.Ground}, // G G G G G G
		},
		Ports: []Port{
			{X: 1, Y: 0, Side: PortNorth},
			{X: 5, Y: 1, Side: PortEast},
		},
	},
	{
		Name: "Fork 1-2",
		Grid: [][]domain.TileType{
			{domain.Ground, domain.Ground, 	domain.Ground}, // G G G
			{domain.Ground, domain.Air, 	domain.Air},	// G A AP
			{domain.Air, 	domain.Air,		domain.Ground},	//PA A G
			{domain.Ground,	domain.Air, 	domain.Ground},	// G A G
		},													//   P
		Ports: []Port{
			{X: 2, Y: 1, Side: PortEast},
			{X: 0, Y: 2, Side: PortWest},
			{X: 1, Y: 3, Side: PortSouth},
		},
	},
	{
		Name: "Join 2-1",
		Grid: [][]domain.TileType{
			{domain.Ground, domain.Ground, 	domain.Ground}, // G G G
			{domain.Air, 	domain.Air, 	domain.Ground},	//PA A G
			{domain.Ground, domain.Air,		domain.Air},	// G A AP
			{domain.Air,	domain.Air, 	domain.Ground},	//PA A G
			{domain.Ground,	domain.Ground, 	domain.Ground},	// G G G
		},
		Ports: []Port{
			{X: 0, Y: 1, Side: PortWest},
			{X: 2, Y: 2, Side: PortEast},
			{X: 0, Y: 3, Side: PortWest},
		},
	},
}