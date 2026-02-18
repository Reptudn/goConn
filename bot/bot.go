package bot

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Reptudn/goConn/internal"
	"github.com/Reptudn/goConn/shared"
	"github.com/Reptudn/goConn/shared/schmeas/actions"
)

type CoreGameBot struct {
	conn     *internal.Connection
	teamName string
	teamId   uint
}

func NewCoreGameBot(teamName string) (*CoreGameBot, error) {

	envIp, exists := os.LookupEnv("SERVER_IP")
	if !exists {
		fmt.Println("Environment variable SERVER_IP not set")
		envIp = "127.0.0.1"
	}

	envPort, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		fmt.Println("Environment variable SERVER_PORT not set")
		envPort = "4444"
	}

	id := os.Args[1]
	if id == "" {
		return nil, fmt.Errorf("team id not provided as first argument")
	}
	teamId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("team id not a number")
	}

	serverAddr := fmt.Sprintf("%s:%s", envIp, envPort)

	conn, err := internal.NewConnection(serverAddr, uint(teamId))
	if err != nil {
		return nil, fmt.Errorf("could not create bot: %w", err)
	}

	return &CoreGameBot{conn: conn, teamName: teamName, teamId: uint(teamId)}, nil
}

func (bot *CoreGameBot) GetGame() *shared.Game {
	return bot.conn.GetGame()
}

func (bot *CoreGameBot) CreateUnit(unitType shared.UnitType) error {
	// TODO: check if spawning is possible
	bot.conn.GetActionQueue().Add(actions.NewActionCreate(unitType))
	return nil
}

func (bot *CoreGameBot) Move(object *shared.Object, pos shared.Position) error {
	if !object.IsReadyForAction() {
		return fmt.Errorf("unit %d is on cooldown for %d more ticks", object.Id, object.ObjectData.(shared.UnitData).ActionCooldown)
	}
	bot.conn.GetActionQueue().Add(actions.NewActionMove(object.Id, pos.X, pos.Y))
	return nil
}

func (bot *CoreGameBot) SimplePathfind(object *shared.Object, pos shared.Position) {
	panic("unimplemented")
}

func (bot *CoreGameBot) Attack(object *shared.Object, target *shared.Object) error {
	if !object.IsReadyForAction() {
		return fmt.Errorf("unit %d is on cooldown for %d more ticks", object.Id, object.ObjectData.(shared.UnitData).ActionCooldown)
	}
	bot.conn.GetActionQueue().Add(actions.NewActionAttack(object.Id, target.Id))
	return nil
}

func (bot *CoreGameBot) TransferGems(source *shared.Object, targetPos shared.Position, amount uint) error {
	if !source.IsReadyForAction() {
		return fmt.Errorf("unit %d is on cooldown for %d more ticks", source.Id, source.ObjectData.(shared.UnitData).ActionCooldown)
	}

	bot.conn.GetActionQueue().Add(actions.NewActionTransferGems(source.Id, targetPos.X, targetPos.Y, amount))
	return nil
}

func (bot *CoreGameBot) Build(builder *shared.Object, pos shared.Position) error {
	if !builder.IsReadyForAction() {
		return fmt.Errorf("unit %d is on cooldown for %d more ticks", builder.Id, builder.ObjectData.(shared.UnitData).ActionCooldown)
	}
	bot.conn.GetActionQueue().Add(actions.NewActionBuild(builder.Id, pos.X, pos.Y))
	return nil
}

func (bot *CoreGameBot) Run(callback func(game *shared.Game)) {
	bot.conn.SetTickCallback(callback)
	bot.conn.Start(bot.teamId, bot.teamName)
	defer bot.stop() // XXX: defer probably not needed since Start() is blocking, but just in case
}

func (bot *CoreGameBot) stop() {
	if err := bot.conn.Close(); err != nil {
		fmt.Printf("Error closing connection: %v\n", err)
	}
}
