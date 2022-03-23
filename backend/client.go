package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID string
	Conn *websocket.Conn
	Pool *Pool
	Name string
}

type Message struct {
	Body string `json:"body"`
	User string `json:"user"`
	Id string `json:"id"`
	Time int64 `json:"time"`
}



func (c* Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
	}()
	for {
		_, body, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		messageId := uuid.NewString()
		time := time.Now().Unix()
		message := Message{Body: string(body), User: c.Name, Id: messageId, Time: time}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}