package controllers

import (
	"andromeda/internal/api/utils"
	"andromeda/pkg/service/entrance/types"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/ably/ably-go/ably"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type CollectionClient struct {
	Client       *WsClient
	CollectionID string
	Disconnect   func()
}

type WsServer struct {
	clients    map[*WsClient]bool
	register   chan *WsClient
	unregister chan *WsClient
	broadcast  chan []byte
	message    chan *ably.Message
}

// Client represents the websocket client at the server
type WsClient struct {
	// The actual websocket connection.
	conn       *websocket.Conn
	wsServer   *WsServer
	send       chan []byte
	disconnect func()
}

var clientsMap sync.Map

type WS struct{}

// NewWebsocketServer creates a new WsServer type
func (ctrl WS) NewWebsocketServer() *WsServer {
	return &WsServer{
		clients:    make(map[*WsClient]bool),
		register:   make(chan *WsClient),
		unregister: make(chan *WsClient),
		broadcast:  make(chan []byte),
		message:    make(chan *ably.Message),
	}
}

// Run our websocket server, accepting various requests
func (server *WsServer) Run() {
	for {
		select {

		case client := <-server.register:
			server.registerClient(client)

		case client := <-server.unregister:
			server.unregisterClient(client)

		case message := <-server.broadcast:
			server.broadcastToClients(message)

		}
	}
}

// Initialize our websocket server
func (ctrl *WS) InitWS() *WsServer {

	ablyKey := os.Getenv("ABLY_KEY")
	transportParams := url.Values{}
	wsServer := ctrl.NewWebsocketServer()
	// fmt.Println("CREATED NEW WS SERVER!")
	go wsServer.Run()

	//Heartbeats enable Ably to identify clients that abruptly disconnect from the service, such as where an internet connection drops out or a client changes networks.
	transportParams.Add("heartbeatInterval", "10000") // 10 sec ( default is 15 sec)

	ablyClient, err := ably.NewRealtime(ably.WithKey(ablyKey), ably.WithTransportParams(transportParams))
	if err != nil {
		panic(err)
	}

	ablyClient.Connect()

	channel := ablyClient.Channels.Get("firehose")
	unsubscribeAll, err := channel.SubscribeAll(context.Background(), func(msg *ably.Message) {
		// fmt.Println(msg.Name)
		wsServer.message <- msg
	})

	if err != nil {
		err := fmt.Errorf("error subscribing to channel: %w", err)
		fmt.Println(err)
		unsubscribeAll()
	}

	return wsServer
}

func (server *WsServer) broadcastToClients(message []byte) {
	for client := range server.clients {
		client.send <- message
	}
}

func (server *WsServer) registerClient(client *WsClient) {
	server.clients[client] = true
}

func (server *WsServer) unregisterClient(client *WsClient) {
	if ok := server.clients[client]; ok {
		delete(server.clients, client)
	}
}

func newClient(conn *websocket.Conn, wsServer *WsServer, disconnect func()) *WsClient {
	client := &WsClient{
		conn:     conn,
		wsServer: wsServer,
		send:     make(chan []byte),
	}

	// Set the disconnect function
	client.disconnect = func() {
		disconnect()

		if client.send != nil {
			close(client.send)
		}

		client.wsServer.unregister <- client
		client.conn.Close()
	}

	return client
}

func removeClient(collectionID string, client *WsClient) {
	value, ok := clientsMap.Load(collectionID)
	if !ok {
		return
	}

	clients := value.([]*CollectionClient)
	for i, c := range clients {
		if c.Client == client {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}

	if len(clients) > 0 {
		clientsMap.Store(collectionID, clients)
	} else {
		clientsMap.Delete(collectionID)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, //! Here should be changed on prod.
}

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

func (client *WsClient) readPump() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		client.wsServer.broadcast <- jsonMessage
	}
}

func (client *WsClient) writePump(wsServer *WsServer, params types.WebsocketParams) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()

	}()

	for {

		var err error
		var res *types.AblyResponseType

		msg := <-wsServer.message

		client.conn.SetWriteDeadline(time.Now().Add(writeWait))
		if err := client.conn.WriteMessage(websocket.PingMessage, []byte("pinging to : "+params.CollectionID)); err != nil {
			fmt.Println("----------> CLIENT DISCONNECTED!")
			return
		}

		err = json.Unmarshal([]byte(msg.Data.(string)), &res)
		if err != nil {
			fmt.Println("Ably stream data transformation error!")
			fmt.Println(err)
		}

		// fmt.Println(res.Item.ProjectSlug)

		if res.Item.ProjectSlug == params.CollectionID || res.Item.ProjectID == params.CollectionID {
			fmt.Println("=====NEW UPDATE FOUND======")
			fmt.Println(res.Item.ProjectSlug)

			value, ok := clientsMap.Load(params.CollectionID)
			if !ok {
				return
			}

			clients := value.([]*CollectionClient)
			for _, c := range clients {

				if err = c.Client.conn.WriteJSON(res); err != nil {
					fmt.Println(err)
					return
				}
			}

		}

	}
}

func (ctrl WS) GetWS(wsServer *WsServer, c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	params, err := utils.GetWebsocketParams(c)
	if err != nil {
		log.Printf("Collection Websocket >> Util GetWebsocketParams; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	client := newClient(conn, wsServer, func() {})
	client.disconnect = func() {
		removeClient(params.CollectionID, client)
	}

	collectionClient := &CollectionClient{
		Client:       client,
		CollectionID: params.CollectionID,
		Disconnect:   client.disconnect,
	}

	value, ok := clientsMap.LoadOrStore(collectionClient.CollectionID, []*CollectionClient{collectionClient})
	if ok {
		clientsMap.Store(collectionClient.CollectionID, append(value.([]*CollectionClient), collectionClient))
		// return
	}

	go client.writePump(wsServer, params)
	go client.readPump()

	wsServer.register <- client

	if err != nil {
		fmt.Println(err)
	}
}
