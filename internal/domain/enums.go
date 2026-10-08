package domain

type TileType int 

type TaskStatus string

const (
	Air TileType = iota
	Ground
	Platform
	Spike
	Start
	Exit
)

const (
	StatusPending TaskStatus = "pending"
	StatusProcessing TaskStatus = "processing"
	StatusCompleted TaskStatus = "completed"
	StatusFailed TaskStatus = "failed"
)

// var stateName = map[TaskStatus]string {
// 	StatusPending: 		"pending",
// 	StatusProcessing: 	"processing",
// 	StatusCompleted:	"completed",
// 	StatusFailed:		"failed",
// }

// fmt package checks for implementation of fmt.Stringer interface, String() specifically.
// So, ts := TaskStatus(0) will map StatusPending to "pending" for ts var.
// func (ts TaskStatus) String() string {
// 	return stateName[ts]
// }

func (tt TileType) IsValid() bool {
	switch tt {
	case Air:
		return true
	case Ground:
		return true
	case Platform:
		return true
	case Spike:
		return true
	case Start:
		return true
	case Exit:
		return true
	default:
		return false
	}
}