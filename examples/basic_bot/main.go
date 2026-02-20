package main

import (
	"fmt"

	"github.com/Reptudn/goConn"
	"github.com/Reptudn/goConn/shared"
)

const teamName = "GoConnBotExample"

var CoreGameBot *goConn.CoreGameBot

func tick(game *shared.Game) {
	game.Log("Tick %d", game.ElapsedTicks)

	// Create a unit if we have enough gems
	_ = CoreGameBot.CreateUnit(shared.UnitWarrior)

	// Move all our units
	for _, obj := range game.GetTeamUnits() {
		game.Log("Unit %d at position (%d, %d)", obj.Id, obj.Pos.X, obj.Pos.Y)
		_ = CoreGameBot.Move(obj, shared.NewPosition(obj.Pos.X+1, obj.Pos.Y))
	}
}

func main() {
	cgb, err := goConn.NewCoreGameBot(teamName)
	if err != nil {
		fmt.Printf("Error creating bot: %v\n", err)
		return
	}
	CoreGameBot = cgb

	if err := CoreGameBot.Run(tick); err != nil {
		fmt.Printf("Bot error: %v\n", err)
	}
}
