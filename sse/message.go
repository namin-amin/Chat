package sse

// Message
//
//	structure of the message used to communicate
type Message struct {
	SenderId   string       `json:"senderId"`   //Id of the client which sent the message
	ReceiverId string       `json:"receiverId"` //Receiver Id,required for E2E message
	Event      MessageTypes `json:"event"`      //MessageType
	Retry      int          `json:"retry"`      //retry timeout
	Data       string       `json:"data"`       //Data to be sent
}

// MessageTypes
//
// Type of messages to be sent
type MessageTypes string

const (
	Open      MessageTypes = "openConnection" //MessageType connect
	Close     MessageTypes = "close"          //MessageType Disconnect
	DMessage  MessageTypes = "message"        //MessageType Direct E2E messages
	Broadcast MessageTypes = "broadcast"      //BroadCase message
	NewConn   MessageTypes = "newConnection"  //Connection added
)
