package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]string)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
var mu sync.Mutex

type Message struct {
	Type     string `json:"type"`
	Username string `json:"username,omitempty"`
	Move     string `json:"move,omitempty"`
	Result   string `json:"result,omitempty"`
	Info     string `json:"info,omitempty"`
}

type GameRoom struct {
	Players   map[*websocket.Conn]string
	Moves     map[string]string
	Broadcast chan Message
	Mutex     sync.Mutex
}

var rooms = make(map[string]*GameRoom)
var roomsMutex sync.Mutex

func InitRoom() {
	fmt.Println("Initializing NetDuel Rock Paper Scissors room...")
	http.HandleFunc("/ws", handleConnections)
	fmt.Println("Room server ready at ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	var joinMsg Message
	err = ws.ReadJSON(&joinMsg)
	if err != nil || joinMsg.Type != "join" || joinMsg.Username == "" || joinMsg.Info == "" {
		ws.WriteJSON(Message{Type: "info", Info: "Invalid join message. Must include room pin."})
		return
	}
	pin := joinMsg.Info
	roomsMutex.Lock()
	room, exists := rooms[pin]
	if !exists {
		room = &GameRoom{
			Players:   make(map[*websocket.Conn]string),
			Moves:     make(map[string]string),
			Broadcast: make(chan Message),
		}
		rooms[pin] = room
		go room.handleMessages()
	}
	roomsMutex.Unlock()
	room.Mutex.Lock()
	room.Players[ws] = joinMsg.Username
	room.Mutex.Unlock()
	room.Broadcast <- Message{Type: "info", Info: joinMsg.Username + " joined room " + pin + "."}

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			room.Mutex.Lock()
			delete(room.Players, ws)
			room.Mutex.Unlock()
			break
		}
		if msg.Type == "move" && (msg.Move == "rock" || msg.Move == "paper" || msg.Move == "scissors") {
			room.Mutex.Lock()
			room.Moves[joinMsg.Username] = msg.Move
			room.Mutex.Unlock()
			room.Broadcast <- Message{Type: "info", Info: joinMsg.Username + " has made a move in room " + pin + "."}
			room.checkAndResolve()
		}
	}
}

func (gr *GameRoom) handleMessages() {
	for {
		msg := <-gr.Broadcast
		gr.Mutex.Lock()
		for client := range gr.Players {
			client.WriteJSON(msg)
		}
		gr.Mutex.Unlock()
	}
}

func (gr *GameRoom) checkAndResolve() {
	gr.Mutex.Lock()
	if len(gr.Moves) == 2 {
		var p1, p2, m1, m2 string
		players := make([]string, 0, 2)
		for _, name := range gr.Players {
			players = append(players, name)
		}
		if len(players) < 2 {
			gr.Mutex.Unlock()
			return
		}
		p1, p2 = players[0], players[1]
		m1, m2 = gr.Moves[p1], gr.Moves[p2]
		result := determineWinner(p1, m1, p2, m2)
		for client, name := range gr.Players {
			if name == p1 || name == p2 {
				client.WriteJSON(Message{Type: "result", Result: result})
			}
		}
		gr.Moves = make(map[string]string)
	}
	gr.Mutex.Unlock()
}

func determineWinner(p1, m1, p2, m2 string) string {
	if m1 == m2 {
		return "It's a tie! Both chose " + m1 + "."
	}
	if (m1 == "rock" && m2 == "scissors") || (m1 == "paper" && m2 == "rock") || (m1 == "scissors" && m2 == "paper") {
		return p1 + " wins! " + m1 + " beats " + m2 + "."
	}
	return p2 + " wins! " + m2 + " beats " + m1 + "."
}
