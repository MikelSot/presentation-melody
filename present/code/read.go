package main

import "log"

//START OMIT
// leemos el mensaje del frontend
func (u user) Read(user *domain.User) {
	defer func() {
		user.Controller.Logout <- user // HL
		user.Conn.Close()              // HL
	}()

	user.Conn.SetReadLimit(maxMessageSize)
	for { // HL
		message := domain.Message{}
		err := user.Conn.ReadJSON(&message) // HL
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err, websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Println("No se puede leer el mensaje", err)
			}
			return
		}
		user.Controller.Broadcast <- message // HL
	} // HL
}

//END OMIT
