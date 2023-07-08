package sse

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

type IClient interface {
	ID() string
	LinkedId() string
	SetLinkedId(id string)
	Channel() chan *Message
	IsVerified() bool
	Verify(verificationStatus bool)
	Close()
}

type ClientVerifyDto struct {
	UserId string `json:"userId"`
	SseId  string `json:"sseId"`
}

// Client
//
// Represents the client object used for communication
type Client struct {
	sseId    string        //unique identification for sse client
	userId   string        //unique identification for the user
	Name     string        //name of the client
	msgChan  chan *Message //channel to send message
	verified bool
}

func (c *Client) ID() string {
	return c.sseId
}

func (c *Client) LinkedId() string {
	return c.userId
}

func (c *Client) Channel() chan *Message {
	return c.msgChan
}

func (c *Client) IsVerified() bool {
	return c.verified
}

func (c *Client) Verify(verificationStatus bool) {
	c.verified = verificationStatus
}

func (c *Client) SetLinkedId(id string) {
	c.userId = id
}

// RegClients
//
// Register client in a hub for SSE connections
func RegClients(h *Hub) echo.HandlerFunc {
	fmt.Println("test")
	return func(c echo.Context) error {
		client := NewClient()
		// close the channel after exit the function
		defer func() {
			client.Close()
			log.Printf("client connection is closed")
		}()
		initSSEHeaders(c)
		h.Subscribe <- &client
		for {
			select {
			case msg := <-client.Channel(): //Normal message
				w := c.Response().Writer
				_, err := fmt.Fprintf(w, "%s", FormatSSEMsg(msg))

				if err != nil {
					fmt.Println("could not send the message")
					fmt.Println(err)
				}
				c.Response().Flush()
			case <-c.Request().Context().Done(): //Close client
				h.Unsubscribe <- &client
				return nil
			}
		}
	}
}

// ClientCheck
//
// Periodic check for checking if the client is still active
func ClientCheck(w *bufio.Writer, c IClient) { //Implement This check
	for {
		time.Sleep(1 * time.Minute)
		_, err := fmt.Fprintf(w, "data: Message: %s\n\n", "health check")
		if err != nil {
			fmt.Println("could not send the message")
			fmt.Println(err)
		}
		err = w.Flush()
		if err != nil {
			fmt.Printf("no response from client: %s. Closing http connection.\n", c.LinkedId())
			c.Channel() <- nil
			break
		}

	}
}

// initSSEHeaders
//
// initiate SSE headers for the connection
func initSSEHeaders(c echo.Context) {
	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("Transfer-Encoding", "chunked")
}

// NewClient
//
// Create a new client and return its instance
func NewClient() Client {
	return Client{
		sseId:    uuid.NewString(),
		userId:   "",
		Name:     "",
		msgChan:  make(chan *Message),
		verified: false,
	}
}

// Close
//
//	empty client and close all channels also
func (c *Client) Close() {
	close(c.Channel())
	c.msgChan = nil
}
