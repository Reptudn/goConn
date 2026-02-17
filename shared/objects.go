package shared

type ObjectType int

const (
	ObjectCore ObjectType = iota
	ObjectUnit
	ObjectDeposit
	ObjectWall
	ObjectGemPile
	ObjectBomb
)

type Position struct {
	X int
	Y int
}

type ObjectData interface {
	isObjectData()
}

type CoreData struct {
	TeamId        uint64
	Gems          uint64
	SpawnCooldown uint64
}

func (CoreData) isObjectData() {}

type UnitData struct {
	UnitType       UnitType
	TeamId         uint64
	Gems           uint64
	ActionCooldown uint64
}

func (UnitData) isObjectData() {}

type DepositData struct {
	Gems uint64
}

func (DepositData) isObjectData() {}

type BombData struct {
	Countdown uint64
}

func (BombData) isObjectData() {}

type Object struct {
	Type       ObjectType
	Data       *any
	Id         uint64
	Pos        Position
	Hp         int32
	ObjectData ObjectData
}
