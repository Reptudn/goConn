# goConn - Game Server Bot Library

A clean and simple Go library for building game server bots with minimal API surface.

## Installation

```bash
go get github.com/Reptudn/goConn
```

## Quick Start

```go
package main

import (
	"github.com/Reptudn/goConn/bot"
	"github.com/Reptudn/goConn/shared"
)

func tick(game *shared.Game) {
	// Your bot logic here
	for _, unit := range game.GetTeamUnits() {
		bot.CreateUnit(shared.UnitWarrior)
		bot.Move(unit, shared.NewPosition(unit.Pos.X+1, unit.Pos.Y))
	}
}

func main() {
	bot, err := bot.NewCoreGameBot("MyBotName")
	if err != nil {
		panic(err)
	}
	bot.Run(tick)
}
```

## Features

- Simple, clean API for bot development
- Type-safe game object handling
- Action queueing system
- Real-time game updates

## Documentation

- See `/bot` package for the main Bot API
- See `/shared` package for game objects and types

## License

MIT

