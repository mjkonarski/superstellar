package communication

import (
	"errors"
	"log"
	"superstellar/backend/pb"
	"superstellar/backend/events"
	"time"

	"github.com/golang/protobuf/proto"

	"golang.org/x/net/websocket"
	"superstellar/backend/monitor"
)

const channelBufSize = 100

// Client struct holds client-specific variables.
type Client struct {
	id       uint32
	username string
	ws       *websocket.Conn
	server   *Server
	ch       chan *[]byte
	doneCh   chan bool
	monitor  *monitor.Monitor
	eventDispatcher *events.EventDispatcher
}

// NewClient initializes a new Client struct with given websocket and Server.
func NewClient(ws *websocket.Conn, server *Server, clientID uint32) *Client {
	if ws == nil {
		panic("ws cannot be nil")
	}

	if server == nil {
		panic("server cannot be nil")
	}

	ch := make(chan *[]byte, channelBufSize)
	doneCh := make(chan bool)
	monitor := server.monitor

	return &Client{clientID, "", ws, server, ch, doneCh, monitor, server.eventsDispatcher}
}

// Conn returns client's websocket.Conn struct.
func (c *Client) Conn() *websocket.Conn {
	return c.ws
}

// SendMessage sends game state to the client.
func (c *Client) SendMessage(bytes *[]byte) {
	select {
	case c.ch <- bytes:
	default:
		c.monitor.AddDroppedMessage()
	}
}

// Done sends done message to the Client which closes the conection.
func (c *Client) Done() {
	c.doneCh <- true
}

// Listen Write and Read request via chanel
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {

		case bytes := <-c.ch:
			before := time.Now()
			err := websocket.Message.Send(c.ws, *bytes)
			after := time.Now()

			if err != nil {
				log.Println(err)
			} else {
				elapsed := after.Sub(before)
				c.monitor.AddSendTime(elapsed)
			}

		case <-c.doneCh:
			c.doneCh <- true
			return
		}
	}
}

func (c *Client) listenRead() {
	log.Println("Listening read from client")
	for {
		select {

		case <-c.doneCh:
			c.doneCh <- true
			return

		default:
			c.readFromWebSocket()
		}
	}
}

func (c *Client) readFromWebSocket() {
	var data []byte
	err := websocket.Message.Receive(c.ws, &data)
	if err != nil {
		log.Println(err)

		c.doneCh <- true
		c.server.deleteClient(c)
		c.eventDispatcher.FireUserLeft(&events.UserLeft{ClientID: c.id})
	} else {
		c.unmarshalUserInput(data)
	}
}

func (c *Client) unmarshalUserInput(data []byte) {
	protoUserMessage := &pb.UserMessage{}
	if err := proto.Unmarshal(data, protoUserMessage); err != nil {
		log.Fatalln("Failed to unmarshal UserInput:", err)
		return
	}

	switch x := protoUserMessage.Content.(type) {
	case *pb.UserMessage_UserAction:
		userInputEvent := events.UserInputFromProto(protoUserMessage.GetUserAction().UserInput, c.id)
		c.eventDispatcher.FireUserInput(userInputEvent)
	case *pb.UserMessage_JoinGame:
		c.tryToJoinGame(protoUserMessage.GetJoinGame())
	default:
		log.Fatalln("Unknown message type %T", x)
	}
}

func (c *Client) tryToJoinGame(joinGameMsg *pb.JoinGame) {
	username := joinGameMsg.Username
	ok, err := validateUsername(username)

	if !ok {
		c.sendJoinGameAckMessage(
			&pb.JoinGameAck{Success: ok, Error: err.Error()},
		)
		return
	}

	c.username = username
	c.eventDispatcher.FireUserJoined(&events.UserJoined{ClientID: c.id, UserName: username})
	c.sendJoinGameAckMessage(&pb.JoinGameAck{Success: true})
	c.sendHelloMessage()
}

func validateUsername(username string) (bool, error) {
	length := len(username)

	if length < 3 {
		return false, errors.New("I doubt your name is shorter than 3 characters, Captain.")
	}

	if length > 25 {
		return false, errors.New("Space fleet doesn't allow names longer than 25 characters!")
	}

	return true, nil
}

func (c *Client) sendJoinGameAckMessage(joinGameAck *pb.JoinGameAck) {
	message := &pb.Message{
		Content: &pb.Message_JoinGameAck{
			JoinGameAck: joinGameAck,
		},
	}

	c.server.SendToClient(c.id, message)
}

func (c *Client) sendHelloMessage() {
	idToUsername := make(map[uint32]string)

	for id, client := range c.server.clients {
		idToUsername[id] = client.username
	}

	message := &pb.Message{
		Content: &pb.Message_Hello{
			Hello: &pb.Hello{
				MyId:         c.id,
				IdToUsername: idToUsername,
			},
		},
	}

	c.server.SendToClient(c.id, message)
}