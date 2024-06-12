package main

import (
    "sync"
    "github.com/gorilla/websocket"
)

var (
    timeLayout = "2006-01-02 15:04:05" // time parse format
    clients = make(map[*websocket.Conn]bool)
    broadcast = make(chan *Note)
    upgrader = websocket.Upgrader{}
    mu sync.Mutex
)

