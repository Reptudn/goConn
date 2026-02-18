package actions

type ActionMove struct {
	Type   string `json:"type"` // "move"
	UnitId uint   `json:"unit_id"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

func NewActionMove(unitId uint, x int, y int) ActionMove {
	return ActionMove{
		Type:   "move",
		UnitId: unitId,
		X:      x,
		Y:      y,
	}
}
