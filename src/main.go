package main

import (
	"github.com/Reptudn/goConn/bot"
	"github.com/Reptudn/goConn/shared"
)

const teamName = "GoConnBot"

var CoreGameBot *bot.CoreGameBot

func tick(game *shared.Game) {
	game.Log("Tick %d", game.ElapsedTicks)

	_ = CoreGameBot.CreateUnit(shared.UnitWarrior)

	// do stuff
	for _, obj := range game.GetTeamUnits() {
		game.Log("Unit %d at position (%d, %d)", obj.Id, obj.Pos.X, obj.Pos.Y)
		_ = CoreGameBot.Move(obj, shared.NewPosition(obj.Pos.X+1, obj.Pos.Y))
	}
}

func main() {
	cgb, err := bot.NewCoreGameBot(teamName)
	if err != nil {
		panic(err)
	}
	CoreGameBot = cgb

	CoreGameBot.Run(tick)
}
