package sse

import (
	"encoding/json"
	"strconv"
)

// FormatSSEMsg
//
// Format given message model to sse event data format
func FormatSSEMsg(message *Message) string {
	data := ""

	if message.SenderId != "" {
		data = "id:" + message.SenderId + "\n"
	}
	if message.Event != "" {
		data = data + "event:" + string(message.Event) + "\n"
	}
	if message.Retry != 0 {
		data = data + "retry:" + strconv.Itoa(message.Retry) + "\n"
	}

	jsonData, _ := json.Marshal(message)
	data = data + "data:" + string(jsonData) + "\n\n"
	return data
}
