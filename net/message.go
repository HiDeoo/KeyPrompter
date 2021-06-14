package net

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	*ErrorMessage
}

type ErrorMessage struct {
	typedMessage
	Message string `json:"message"`
}

type typedMessage struct {
	Type string `json:"type"`
}

func (message *Message) UnmarshalJSON(data []byte) error {
	var getType typedMessage

	if err := json.Unmarshal(data, &getType); err != nil {
		return err
	}

	switch getType.Type {
	case "error":
		message.ErrorMessage = &ErrorMessage{}
		return json.Unmarshal(data, message.ErrorMessage)
	default:
		if getType.Type == "" {
			return fmt.Errorf("Missing message type.")
		}

		return fmt.Errorf("Unrecognized message type: %q.", getType.Type)
	}
}

func (message Message) MarshalJSON() ([]byte, error) {
	switch {
	case message.ErrorMessage != nil:
		return json.Marshal(message.ErrorMessage)
	default:
		return json.Marshal(nil)
	}
}
