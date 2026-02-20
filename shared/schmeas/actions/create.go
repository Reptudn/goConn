package actions

import "github.com/Reptudn/goConn/shared"

type ActionCreate struct {
	Type     string          `json:"type"`      // "create"
	UnitType shared.UnitType `json:"unit_type"` // needs to be a number
}

func NewActionCreate(unitType shared.UnitType) ActionCreate {
	return ActionCreate{
		Type:     "create",
		UnitType: unitType,
	}
}
