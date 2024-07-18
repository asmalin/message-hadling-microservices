package model

import "time"

type MessageInput struct {
	Text string `json:"text" `
}

type Message struct {
	Id        int       `json:"id" `
	Text      string    `json:"text" `
	Processed bool      `json:"processed" `
	CreatedAt time.Time `json:"created_at"`
}
