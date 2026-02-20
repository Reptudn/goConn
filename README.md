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
	goconn "github.com/42core-team/go-client-lib"
	"github.com/42core-team/go-client-lib/shared"
)

var coreBot *goconn.CoreGameBot

func tick(game *shared.Game) {
	// Your bot logic here
	for _, unit := range game.GetTeamUnits() {
		_ = coreBot.CreateUnit(shared.UnitWarrior)
		_ = coreBot.Move(unit, shared.NewPosition(unit.Pos.X+1, unit.Pos.Y))
	}
}

func main() {
	bot, err := goconn.NewCoreGameBot("MyBotName")
	if err != nil {
		panic(err)
	}
	coreBot = bot
	_ = coreBot.Run(tick)
}
```

## Features

- Simple, clean API for bot development
- Type-safe game object handling
- Action queueing system
- Real-time game updates

## Documentation

- See the root `goConn` package for the main Bot API
- See `/shared` package for game objects and types

## License

MIT
