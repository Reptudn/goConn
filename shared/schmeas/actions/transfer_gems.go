package actions

type ActionTransferGems struct {
	Type     string `json:"type"` // "transfer_gems"‚
	SourceId uint   `json:"source_id"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Amount   uint   `json:"amount"` // min 1
}

func NewActionTransferGems(sourceId uint, x int, y int, amount uint) ActionTransferGems {
	return ActionTransferGems{
		Type:     "transfer_gems",
		SourceId: sourceId,
		X:        x,
		Y:        y,
		Amount:   amount,
	}
}
