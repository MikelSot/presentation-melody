package main

import (
	"log"
	"time"
)

//START OMIT
// leemos el mensaje y lo enviamos a frontend
func (u user) Write(user *domain.User) {
	defer func() {
		user.Conn.Close() // HL
	}()

	for { // HL
		select {
		case message, status := <-user.Message: // HL
			if !status {
				return
			}
			user.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := user.Conn.WriteJSON(message) // HL
			if err != nil {
				log.Println("no se puedo escribir el mensaje en el ws")
				return
			}
		}
	} // HL
}

//END OMIT
