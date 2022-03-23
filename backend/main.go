package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var Ctx = context.Background()

var RedisClient *redis.Client

func serveWs(pool *Pool,w http.ResponseWriter, r *http.Request) {

    clientName := strings.Join(r.URL.Query()["user"], "")
    fmt.Println("websocket endpoint hit", clientName)
    conn, err := Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &Client {
        Conn: conn,
        Pool: pool,
        Name: clientName,
    }

    pool.Register <- client

    client.Read()
}

func home(rw http.ResponseWriter, r *http.Request) {
    fmt.Println("hit to entry")
    client := RedisClient.Conn(Ctx)
    defer client.Close()
    iter := client.Scan(Ctx, 0, "*", 0).Iterator()

    var messages []Message
    fmt.Println(iter, "iter")
    for iter.Next(Ctx) {
        fmt.Println("keys", iter.Val())
        
        val := client.Get(Ctx,iter.Val()).Val()
        fmt.Println(val, "values")
        var message Message
        
        json.Unmarshal([]byte(val), &message)
        fmt.Println(message, "here is a message")
        messages = append(messages, message)
    }
    if err := iter.Err(); err != nil {
        panic(err)
    }

    sort.Slice(messages, func (i, j int) bool {
        return messages[i].Time < messages[j].Time
    })


    rw.Header().Set("Content-Type", "application/json")
    json.NewEncoder(rw).Encode(messages)
    
    fmt.Println(messages)
}

func newRouter() *mux.Router {

    pool := NewPool()

    go pool.Start()

	router := mux.NewRouter()
	router.HandleFunc("/allMessages", home).Methods("GET")
    router.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
        serveWs(pool, rw, r)
    })
	return router
}

func main() {
    RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

    router := newRouter()

    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))

}