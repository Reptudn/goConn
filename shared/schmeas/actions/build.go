package actions

type ActionBuild struct {
	Type   string `json:"type"` // "build"‚
	UnitId uint   `json:"unit_id"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

func NewActionBuild(unitId uint, x int, y int) ActionBuild {
	return ActionBuild{
		Type:   "build",
		UnitId: unitId,
		X:      x,
		Y:      y,
	}
}
