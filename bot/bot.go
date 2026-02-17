package bot

import (
	"fmt"
	"os"

	"github.com/Reptudn/goConn/actions"
	"github.com/Reptudn/goConn/internal"
	"github.com/Reptudn/goConn/shared"
)

type CoreGameBot struct {
	conn     *internal.Connection
	teamName string
	teamId   int
}

func NewCoreGameBot(teamName string, teamId int) (*CoreGameBot, error) {

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
	actions.CreateUnit(unitType, bot.conn.GetActionQueue())
}

func (bot *CoreGameBot) Move(object *shared.Object, pos shared.Position) {
	actions.Move(object, pos, bot.conn.GetActionQueue())
}

func (bot *CoreGameBot) SimplePathfind(object *shared.Object, pos shared.Position) {
	actions.SimplePathfind(object, pos, bot.conn.GetActionQueue())
}

func (bot *CoreGameBot) Attack(object *shared.Object, target *shared.Object) {
	actions.Attack(object, target, bot.conn.GetActionQueue())
}

func (bot *CoreGameBot) TransferGems(source *shared.Object, targetPos shared.Position, amount uint64) {
	actions.TransferGems(source, targetPos, amount, bot.conn.GetActionQueue())
}

func (bot *CoreGameBot) Build(builder *shared.Object, pos shared.Position) {
	actions.Build(builder, pos, bot.conn.GetActionQueue())
}

func (bot *CoreGameBot) Run(callback func(game *shared.Game)) {
	bot.conn.SetTickCallback(callback)
	bot.conn.Start()
	defer bot.stop()
}

func (bot *CoreGameBot) stop() {
	if err := bot.conn.Close(); err != nil {
		fmt.Printf("Error closing connection: %v\n", err)
	}
}
