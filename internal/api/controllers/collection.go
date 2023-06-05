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
	"time"

	"github.com/ably/ably-go/ably"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsServer struct {
	clients    map[*WsClient]bool
	register   chan *WsClient
	unregister chan *WsClient
	broadcast  chan []byte
}

// Client represents the websocket client at the server
type WsClient struct {
	// The actual websocket connection.
	conn     *websocket.Conn
	wsServer *WsServer
	send     chan []byte
}

type Collection struct{}

// GetTrends godoc
//
// @Summary         Get collection trends
// @Description     get trending collections
// @Tags            collections
// @Accept          json
// @Produce         json
// @Param           period   query         string  true         "Period (1h|1d|7d)"
// @Param           sort_by  query         string  true         "Sort by (volume)"
// @Param           order    query         string  true         "Order (asc|desc)"
// @Param           limit    query         int     true         "Limit"
// @Param           offset   query         int     true         "Offset"
// @Success		    200	     {object}	   types.TrendRes
// @Failure		    400
// @Failure         500
// @Router          /collections/trend     [get]
func (ctrl Collection) GetTrends(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=60, stale-while-revalidate")
	params, err := utils.GetTrendParams(c)
	if err != nil {
		log.Printf("Collection GetTrends >> Util GetTrendParams; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	dataProvider := utils.GetProvider()
	trends, err := dataProvider.GetCollectionTrends(&params)
	if err != nil {
		log.Printf("Collection GetTrends >> DataProvder GetCollectionTrends; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, trends)
}

// GetNFTs godoc
//
// @Summary         Get collection NFTs
// @Description     get the list of NFTs of the collection
// @Tags            collections
// @Accept          json
// @Produce         json
// @Param           address         query         string  true         "Collection slug"
// @Param           attributes      query         string  false        "NFT attributes to filter ([{'name': 'Tattoos', 'type': 'CATEGORY', 'values': ['Barbwire']}])"
// @Param           listing_only    query         string  false        "Only listed NFTs? (true|false)"
// @Param           program         query         string  false        "Marketplace program address"
// @Param           auction_house   query         string  false        "Auction house address"
// @Param           name            query         string  false        "NFT name"
// @Param           min             query         number  false        "Minimum listing price"
// @Param           max             query         number  false        "Maximum listing price"
// @Param           sort_by         query         string  true         "Sort By (lowest_listing_block_timestamp)"
// @Param           order           query         string  true         "Order (asc|desc)"
// @Param           limit           query         int     true         "Limit"
// @Param           offset          query         int     true         "Offset"
// @Success		    200	            {object}	  types.NFTRes
// @Failure		    400
// @Failure         500
// @Router          /collections/nfts     [get]
func (ctrl Collection) GetNFTs(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=10, stale-while-revalidate")
	params, err := utils.GetNFTParams(c)
	if err != nil {
		log.Printf("Collection GetNFTs >> Util GetNFTParams; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	dataProvider := utils.GetProvider()
	nfts, err := dataProvider.GetCollectionNFTs(&params)

	if err != nil {
		log.Printf("Collection GetNFTs >> DataProvder GetCollectionNFTs; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nfts)
}

// GetTimeSeries godoc
//
// @Summary         Get collection historical data
// @Description     get the historical stats for the collection
// @Tags            collections
// @Accept          json
// @Produce         json
// @Param           address       query         string  true         "Collection slug"
// @Param           from_time     query         int     true         "Start timestamp"
// @Param           to_time       query         int     true         "End timestamp"
// @Param           granularity   query         string  true         "Granularity (per_hour|per_day)"
// @Param           limit         query         int     true         "Limit"
// @Param           offset        query         int     true         "Offset"
// @Success		    200	          {object}	    types.TimeSeriesRes
// @Failure		    400
// @Failure         500
// @Router          /collections/series     [get]
func (ctrl Collection) GetTimeSeries(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=300, stale-while-revalidate")
	params, err := utils.GetTimeSeriesParams(c)
	if err != nil {
		log.Printf("Collection GetTimeSeries >> Util GetTimeSeriesParams; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	dataProvider := utils.GetProvider()
	series, err := dataProvider.GetCollectionTimeSeries(&params)

	if err != nil {
		log.Printf("Collection GetTimeSeries >> DataProvder GetCollectionTimeSeries; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, series)
}

// GetDetail godoc
//
// @Summary         Get collection detail
// @Description     get collection detail information with the address
// @Tags            collections
// @Accept          json
// @Produce         json
// @Param           address  path          string true                     "Collection slug"
// @Success		    200	     {object}	   types.Collection
// @Failure         500
// @Router          /collections/{address} [get]
func (ctrl Collection) GetDetail(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=300, stale-while-revalidate")
	address := c.Param("address")

	dataProvider := utils.GetProvider()
	collection, err := dataProvider.GetCollectionDetail(address)
	if err != nil {
		log.Printf("Collection GetDetail >> DataProvder GetDetail; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, collection)
}

// GetActivities godoc
//
// @Summary         Get collection activities
// @Description     get the activities with related to the collection
// @Tags            collections
// @Accept          json
// @Produce         json
// @Param           address          query         string  true         "Collection slug"
// @Param           limit            query         int     true         "Limit"
// @Param           offset           query         int     true         "Offset"
// @Param           activity_types   query         string  true         "Activity types (['listing'])"
// @Success		    200	             {object}	   types.ActivityRes
// @Failure		    400
// @Failure         500
// @Router          /collections/activities     [get]
func (ctrl Collection) GetActivities(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=10, stale-while-revalidate")
	params, err := utils.GetActivityParams(c, false)
	if err != nil {
		log.Printf("Collection GetActivities >> Util GetActivityParams; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	dataProvider := utils.GetProvider()
	activityRes, err := dataProvider.GetCollectionActivities(&params)

	if err != nil {
		log.Printf("Collection GetTimeSeries >> DataProvder GetCollectionTimeSeries; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, activityRes)
}

// NewWebsocketServer creates a new WsServer type
func (ctrl Collection) NewWebsocketServer() *WsServer {
	return &WsServer{
		clients:    make(map[*WsClient]bool),
		register:   make(chan *WsClient),
		unregister: make(chan *WsClient), broadcast: make(chan []byte),
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

func (server *WsServer) broadcastToClients(message []byte) {
	for client := range server.clients {
		client.send <- message
	}
}

func (server *WsServer) registerClient(client *WsClient) {
	server.clients[client] = true
}

func (server *WsServer) unregisterClient(client *WsClient) {
	if _, ok := server.clients[client]; ok {
		delete(server.clients, client)
	}
}

func newClient(conn *websocket.Conn, wsServer *WsServer) *WsClient {
	return &WsClient{
		conn:     conn,
		wsServer: wsServer,
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
	pongWait = 10 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 5) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

func (client *WsClient) disconnect() {
	if client.send != nil {
		close(client.send)
	}

	client.wsServer.unregister <- client
	client.conn.Close()
	// fmt.Println("----------> WEBSOCKET CLOSED")
}

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

func (ctrl Collection) GetWs(wsServer *WsServer, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := newClient(conn, wsServer)
	params, err := utils.GetWebsocketParams(c)
	// fmt.Println(params)

	if err != nil {
		log.Printf("Collection Websocket >> Util GetWebsocketParams; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	ablyKey := os.Getenv("ABLY_KEY")
	var transportParams url.Values = url.Values{}
	//Heartbeats enable Ably to identify clients that abruptly disconnect from the service, such as where an internet connection drops out or a client changes networks.
	transportParams.Add("heartbeatInterval", "10000") // 10 sec ( default is 15 sec)

	ablyClient, err := ably.NewRealtime(ably.WithKey(ablyKey), ably.WithTransportParams(transportParams))
	if err != nil {
		panic(err)
	}

	ablyClient.Connect()
	//* That 'firehose' channel may need to be put in env
	channel := ablyClient.Channels.Get("firehose")

	unsubscribe, err := channel.SubscribeAll(context.Background(), func(msg *ably.Message) {

		for {
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			// fmt.Println("Pinging!")

			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				// fmt.Println("----------> UNSUBSCRIBED!")
				channel.OffAll()
				ablyClient.Close()
				return
			}

			var res *types.AblyResponseType

			err = json.Unmarshal([]byte(msg.Data.(string)), &res)
			if err != nil {
				fmt.Println("Ably stream data transformation error!")
				fmt.Println(err)
			}

			// fmt.Println(res.Item.ProjectSlug)

			if res.Item.ProjectID != params.CollectionID && res.Item.ProjectSlug != params.CollectionID {
				return
			}

			if res.Item.ProjectSlug == params.CollectionID || res.Item.ProjectID == params.CollectionID {
				// fmt.Println("=====NEW UPDATE FOUND======")
				// fmt.Println(res.ActionType)
				// fmt.Println(res.Item.TokenAddress)

				// Write message back to browser
				if err = client.conn.WriteJSON(res); err != nil {
					return
				}
				return
			}
		}
	})

	go client.readPump()

	wsServer.register <- client

	// fmt.Println(len(wsServer.clients))

	if err != nil {
		unsubscribe()
		// panic(err)
		fmt.Println(err)
	}
}
