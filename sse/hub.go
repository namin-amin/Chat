package sse

import (
	"fmt"
	"strconv"
	"sync"
)

// Hub
//
//	it's a central place where all SSE connections are handled
type Hub struct {
	Broadcast    chan *Message      //BroadCast messages to all clients
	SendMsg      chan *Message      //Send E2E message
	Subscribe    chan *Client       //Subscribe client with the Hub
	Unsubscribe  chan *Client       //Unsubscribe client with Hub
	Clients      map[string]*Client //list of clients subscribed to the hub
	ConMutes     sync.RWMutex
	VerifyClient chan *ClientVerifyDto
}

// Run
//
// Listen to different events on the HUB
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Subscribe: //subscribe
			fmt.Println("Register client : " + client.ID())
			h.ConMutes.RLock()
			if _, isRoomExist := h.Clients[client.ID()]; !isRoomExist {
				h.Clients[client.ID()] = client
			}
			fmt.Println("total number of clients now :" + strconv.Itoa(len(h.Clients)))
			h.ConMutes.RUnlock()
			if client.Channel() != nil {
				client.Channel() <- &Message{
					SenderId:   client.LinkedId(),
					ReceiverId: "",
					Event:      Open,
					Retry:      5000,
					Data:       client.ID(),
				}
			}
		case client := <-h.Unsubscribe: //unsubscribe
			fmt.Println("unregistered client : " + client.ID())
			h.ConMutes.RLock()
			delete(h.Clients, client.ID())
			h.ConMutes.RUnlock()
		case msg := <-h.SendMsg: //send direct message
			fmt.Println("dm to :" + msg.SenderId + "   \nreceiver: " + msg.ReceiverId)
			go func(m *Message) {
				for _, client := range h.Clients {
					if client.LinkedId() == m.ReceiverId { //&& client.IsVerified() {
						client.Channel() <- m
						break
					}
				}
			}(msg)
		case verify := <-h.VerifyClient:
			h.ConMutes.RLock()
			fmt.Println(verify.SseId + "  " + verify.UserId)
			if client, isRoomExist := h.Clients[verify.SseId]; !isRoomExist {
				fmt.Println("Error user not found to verify")
			} else {
				client.SetLinkedId(verify.UserId)
				client.Verify(true)
				fmt.Println("verified")
			}
			h.ConMutes.RUnlock()
		}
	}
}

// BroadcastMsg
//
// Broadcast given message to all the users
func (h *Hub) BroadcastMsg() {
	for msg := range h.Broadcast {
		for _, client := range h.Clients {
			h.ConMutes.Lock()
			if client.Channel() != nil && client.IsVerified() {
				client.Channel() <- msg
			}
			h.ConMutes.Unlock()
		}
	}

}

// NewHub
//
// Create a new instance of Hub and return it
func NewHub() *Hub {
	return &Hub{
		Broadcast:    make(chan *Message, 5),
		SendMsg:      make(chan *Message),
		Subscribe:    make(chan *Client),
		Unsubscribe:  make(chan *Client),
		Clients:      make(map[string]*Client),
		VerifyClient: make(chan *ClientVerifyDto),
	}
}
