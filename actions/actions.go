package actions

import (
	"encoding/json"

	"github.com/Reptudn/goConn/shared"
)

type ActionType int

const (
	ActionCreate ActionType = iota
	ActionMove
	ActionAttack
	ActionTransfer
	ActionBuild
)

type ActionData interface {
	isActionData()
}

type CreateActionData struct {
	UnitType shared.UnitType
}

func (CreateActionData) isActionData() {}

type MoveActionData struct {
	ObjectId uint64
	Position shared.Position
}

func (MoveActionData) isActionData() {}

type AttackActionData struct {
	ObjectId uint64
	TargetId uint64
}

func (AttackActionData) isActionData() {}

type TransferActionData struct {
	SourceId  uint64
	TargetPos shared.Position
	Amount    uint64
}

func (TransferActionData) isActionData() {}

type BuildActionData struct {
	BuilderId uint64
	Position  shared.Position
}

func (BuildActionData) isActionData() {}

type Action struct {
	Type ActionType
	Data ActionData
}

func (Action) Marshal() ([]byte, error) {
	return json.Marshal(map[string]interface{}{})
}

func CreateUnit(unitType shared.UnitType, actionQueue *ActionQueue) {
	actionQueue.Add(Action{Type: ActionCreate, Data: CreateActionData{UnitType: unitType}})
}

func Move(object *shared.Object, pos shared.Position, actionQueue *ActionQueue) {
	actionQueue.Add(Action{Type: ActionMove, Data: MoveActionData{ObjectId: object.Id, Position: pos}})
}

func SimplePathfind(object *shared.Object, pos shared.Position, actionQueue *ActionQueue) {}

func Attack(object *shared.Object, target *shared.Object, actionQueue *ActionQueue) {
	actionQueue.Add(Action{Type: ActionAttack, Data: AttackActionData{ObjectId: object.Id, TargetId: target.Id}})
}

func TransferGems(source *shared.Object, targetPos shared.Position, amount uint64, actionQueue *ActionQueue) {
	actionQueue.Add(Action{Type: ActionTransfer, Data: TransferActionData{SourceId: source.Id, TargetPos: targetPos, Amount: amount}})
}

func Build(builder *shared.Object, pos shared.Position, actionQueue *ActionQueue) {
	actionQueue.Add(Action{Type: ActionBuild, Data: BuildActionData{BuilderId: builder.Id, Position: pos}})
}
