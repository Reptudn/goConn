package bot

import (
	"fmt"
	"os"

	"github.com/Reptudn/goConn/internal"
	"github.com/Reptudn/goConn/shared"
	"github.com/Reptudn/goConn/shared/schmeas/actions"
)

type CoreGameBot struct {
	conn     *internal.Connection
	teamName string
	teamId   uint
}

func NewCoreGameBot(teamName string, teamId uint) (*CoreGameBot, error) {

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

	serverAddr := fmt.Sprintf("%s:%s", envIp, envPort)

	conn, err := internal.NewConnection(serverAddr)
	if err != nil {
		return nil, fmt.Errorf("could not create bot: %w", err)
	}

	return &CoreGameBot{conn: conn, teamName: teamName, teamId: teamId}, nil
}

func (bot *CoreGameBot) GetGame() *shared.Game {
	return bot.conn.GetGame()
}

func (bot *CoreGameBot) CreateUnit(unitType shared.UnitType) {
	bot.conn.GetActionQueue().Add(actions.NewActionCreate(unitType))
}

func (bot *CoreGameBot) Move(object *shared.Object, pos shared.Position) {
	bot.conn.GetActionQueue().Add(actions.NewActionMove(object.Id, pos.X, pos.Y))
}

func (bot *CoreGameBot) SimplePathfind(object *shared.Object, pos shared.Position) {
	panic("unimplemented")
}

func (bot *CoreGameBot) Attack(object *shared.Object, target *shared.Object) {
	bot.conn.GetActionQueue().Add(actions.NewActionAttack(object.Id, target.Id))
}

func (bot *CoreGameBot) TransferGems(source *shared.Object, targetPos shared.Position, amount uint) {
	bot.conn.GetActionQueue().Add(actions.NewActionTransferGems(source.Id, targetPos.X, targetPos.Y, amount))
}

func (bot *CoreGameBot) Build(builder *shared.Object, pos shared.Position) {
	bot.conn.GetActionQueue().Add(actions.NewActionBuild(builder.Id, pos.X, pos.Y))
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
