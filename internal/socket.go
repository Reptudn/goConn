package internal

import (
	"bufio"
	"fmt"
	"net"

	"github.com/Reptudn/goConn/actions"
	"github.com/Reptudn/goConn/shared"
)

type Connection struct {
	socket         *net.Conn
	reader         *bufio.Reader
	game           *shared.Game
	onTickCallback func(*shared.Game)
	actionQueue    *actions.ActionQueue
}

func NewConnection(serverAddr string) (*Connection, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to server: %v", err)
	}

	return &Connection{
		socket:         &conn,
		reader:         bufio.NewReader(conn),
		game:           &shared.Game{},
		onTickCallback: nil,
		actionQueue:    actions.NewActionQueue(100),
	}, nil
}

func (connection *Connection) Start() {
	go func() {
		buffer := make([]byte, 4096)
		for {
			n, err := connection.reader.Read(buffer)
			if err != nil {
				panic(err)
			}
			if err := connection.handleRecv(buffer[:n]); err != nil {
				fmt.Println("Failed to handle the latest socket message!")
				return
			}
		}
	}()
}

func (connection *Connection) SendActions(plannedActions []actions.Action) error {
	for _, action := range plannedActions {

		buffer, err := action.Marshal()
		if err != nil {
			return fmt.Errorf("error marshaling action: %v", err)
		}

		if _, err := (*connection.socket).Write(buffer); err != nil {
			return fmt.Errorf("error sending action to server: %v", err)
		}
	}
	return nil
}

func (connection *Connection) handleRecv(buffer []byte) error {

	if err := connection.parseGameState(buffer); err != nil {
		return fmt.Errorf("error parsing game state: %v", err)
	}

	queue := connection.GetActionQueue()
	plannedActions := queue.GetAll()
	if connection.onTickCallback != nil {
		connection.onTickCallback(connection.game)
	}
	// send all the actions to the server
	if err := connection.SendActions(plannedActions); err != nil {
		return fmt.Errorf("error sending actions to server: %v", err)
	}

	return nil
}

func (connection *Connection) Close() error {
	return (*connection.socket).Close()
}

func (connection *Connection) GetGame() *shared.Game {
	return connection.game
}

func (connection *Connection) GetActionQueue() *actions.ActionQueue {
	return connection.actionQueue
}

func (connection *Connection) SetTickCallback(callback func(*shared.Game)) {
	connection.onTickCallback = callback
}

func (connection *Connection) AddAction(action actions.Action) {
	connection.actionQueue.Add(action)
}
