package internal

import "fmt"

func (connection *Connection) parseGameState(buffer []byte) error {
	fmt.Printf("Game state received: %s\n", string(buffer))
	return nil
}
