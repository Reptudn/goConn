package shared

import "fmt"

type BuildType int

const (
	BuildTypeNone BuildType = iota
	BuildTypeWall
	BuildTypeBomb
)

type UnitConfig struct {
	Name                   string
	UnitType               UnitType
	Cost                   uint64
	Hp                     uint64
	BaseActionCooldown     uint64
	MaxActionCooldown      uint64
	BalancePerCooldownStep uint64
	DamageCore             uint64
	DamageUnit             uint64
	DamageDeposit          uint64
	DamageWall             uint64
	DamageBomb             uint64
	BuildType              BuildType
}

type Config struct {
	GridSize          uint64
	IdleIncome        uint64
	IdleIncomeTimeout uint64
	DepositHp         uint64
	DepositIncome     uint64
	GemPileIncome     uint64
	CoreHp            uint64
	CoreSpawnCooldown uint64
	InitialBalance    uint64
	WallHp            uint64
	WallBuildCost     uint64
	BombCountdown     uint64
	BombThrowCost     uint64
	BombReach         uint64
	BombDamageCore    uint64
	BombDamageUnit    uint64
	BombDamageDeposit uint64
	Units             []UnitConfig
}

type Game struct {
	ElapsedTicks uint64
	Config       Config
	MyTeamId     uint64
	Objects      []Object
}

func (game *Game) GetObjectById(id uint64) *Object {
	return &game.Objects[0]
}

func (game *Game) GetObjectFromPosition(pos Position) *Object {
	return &game.Objects[0]
}

func (game *Game) GetObjectsFromFilter(filter func(object *Object) bool) []*Object {
	return []*Object{&game.Objects[0]}
}

func (game *Game) GetObjectFromFilterNearest(pos Position, filter func(object *Object) bool) *Object {
	return &game.Objects[0]
}

func (game *Game) GetUnitConfigByType(unitType UnitType) *UnitConfig {
	return &game.Config.Units[0]
}

func (game *Game) Log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
