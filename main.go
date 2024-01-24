package main

import (
	"fmt"
	"log"

	"github.com/4lexir4/trading/orderbook"
	"github.com/4lexir4/trading/socket"
	"github.com/gorilla/websocket"
)

func main() {
	ws, _, err := websocket.DefaultDialer.Dial("ws://localhost:4000", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer ws.Close()

	msg := socket.Message{
		Type:    "subscribe",
		Topic:   "spreads",
		Symbols: []string{"BTCUSD", "ETHUSD"},
	}

	if err := ws.WriteJSON(msg); err != nil {
		log.Fatal(err)
	}

	for {
		msgg := []orderbook.CrossSpread{}
		if err := ws.ReadJSON(&msgg); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msgg)
	}
}
