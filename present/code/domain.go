package main

import "golang.org/x/net/websocket"

//START OMIT
type User struct {
	Name       string `json:"name"`
	Controller *Controller
	Conn       *websocket.Conn
	Message    chan Message
}

//END OMIT
//START_CONTROLER OMIT
type Controller struct {
	Users     map[string]*User
	Login     chan *User
	Logout    chan *User
	Broadcast chan Message
}

//END_CONTROLER OMIT

//START_MESSAGE OMIT
type Message struct {
	Username string `json:"name"`
	Message  string `json:"message"`
}

//END_MESSAGE OMIT
