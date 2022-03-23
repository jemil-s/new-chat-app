package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
        select {
        case client := <-pool.Register:
            pool.Clients[client] = true

            body := client.Name + " joined the chat"
            messageId := uuid.NewString()
            time := time.Now().Unix()
            message := Message{ Body: body, Id: messageId, Time: time}
            stringified, _ := json.Marshal(message)
            redisClient := RedisClient.Conn(Ctx)
            defer redisClient.Close()
            RedisClient.Set(Ctx,messageId,stringified,0)
            for client := range pool.Clients {
                client.Conn.WriteJSON(message)
            }
        case client := <-pool.Unregister:
            delete(pool.Clients, client)

            body := client.Name + " disconnected from the chat"
            messageId := uuid.NewString()
            time := time.Now().Unix()
            message := Message{ Body: body, Id: messageId, Time: time}
            stringified, _ := json.Marshal(message)

            redisClient := RedisClient.Conn(Ctx)
            defer redisClient.Close()

            redisClient.Set(Ctx,messageId,stringified,0)
            for client := range pool.Clients {
                client.Conn.WriteJSON(message)
            }
        case message := <-pool.Broadcast:
            stringified, _ := json.Marshal(message)
            redisClient := RedisClient.Conn(Ctx)
            defer redisClient.Close()
            redisClient.Set(Ctx,message.Id,stringified,0)
            for client := range pool.Clients {
                if err := client.Conn.WriteJSON(message); err != nil {
                    fmt.Println(err)
                    return
                }
            }
        }
    }
}
