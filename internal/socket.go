package internal

import (
	"bufio"
	"fmt"
	"net"

	"github.com/Reptudn/goConn/actions"
	"github.com/Reptudn/goConn/shared"
	"github.com/Reptudn/goConn/shared/schmeas"
	schema_action "github.com/Reptudn/goConn/shared/schmeas/actions"
)

type Connection struct {
	socket         *net.Conn
	reader         *bufio.Reader
	Game           *shared.Game
	onTickCallback func(*shared.Game)
	actionQueue    *actions.ActionQueue
}

func NewConnection(serverAddr string, selfTeamId uint) (*Connection, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to server: %v", err)
	}

	return &Connection{
		socket:         &conn,
		reader:         bufio.NewReader(conn),
		Game:           &shared.Game{MyTeamId: selfTeamId},
		onTickCallback: nil,
		actionQueue:    actions.NewActionQueue(100),
	}, nil
}

func (connection *Connection) Start(teamId uint, teamName string) {

	if err := connection.sendLoginPacket(teamId, teamName); err != nil {
		fmt.Printf("Failed to send login packet: %v\n", err)
		return
	}

	buffer := make([]byte, 4096)
	for {
		n, err := connection.reader.Read(buffer)
		if err != nil {
			panic(err)
		}
		if err := connection.handleReceive(buffer[:n]); err != nil {
			fmt.Println("Failed to handle the latest socket message!")
			return
		}
	}
}

func (connection *Connection) sendLoginPacket(teamId uint, teamName string) error {
	loginPacket, err := schmeas.NewLoginRequest(teamId, teamName).Marshal()
	if err != nil {
		return fmt.Errorf("Error marshaling login request: %v\n", err)
	}

	if err := connection.Send(loginPacket); err != nil {
		return fmt.Errorf("Error sending login request: %v\n", err)
	}
	return nil
}

func (connection *Connection) SendString(message string) error {
	if _, err := (*connection.socket).Write([]byte(message)); err != nil {
		return fmt.Errorf("error sending data to server: %v", err)
	}
	return nil
}

func (connection *Connection) Send(buffer []byte) error {
	if _, err := (*connection.socket).Write(buffer); err != nil {
		return fmt.Errorf("error sending data to server: %v", err)
	}
	return nil
}

func (connection *Connection) SendActions(plannedActions []schema_action.Action) error {
	clientPacket, err := schmeas.NewClientPacket(plannedActions).Marshal()
	if err != nil {
		return fmt.Errorf("error marshaling client packet: %v", err)
	}

	if err := connection.Send(clientPacket); err != nil {
		return fmt.Errorf("error sending client packet: %v", err)
	}
	return nil
}

func (connection *Connection) handleReceive(buffer []byte) error {

	tick, err := NewGameTick(buffer)
	if err != nil {
		return fmt.Errorf("error parsing game tick: %v", err)
	}
	tick.UpdateGame(connection.Game)

	queue := connection.GetActionQueue()
	plannedActions := queue.GetAll()
	// send all the actions to the server
	if err := connection.SendActions(plannedActions); err != nil {
		return fmt.Errorf("error sending actions to server: %v", err)
	}

	if connection.onTickCallback != nil {
		connection.onTickCallback(connection.Game)
	}

	return nil
}

func (connection *Connection) Close() error {
	return (*connection.socket).Close()
}

func (connection *Connection) GetGame() *shared.Game {
	return connection.Game
}

func (connection *Connection) GetActionQueue() *actions.ActionQueue {
	return connection.actionQueue
}

func (connection *Connection) SetTickCallback(callback func(*shared.Game)) {
	connection.onTickCallback = callback
}
