package main

//START OMIT
func (c controller) Run() {
	for { // HL
		select {
		case user := <-c.controller.Login: // HL
			c.data.AddUser(user)
		case user := <-c.controller.Logout: // HL
			c.data.DeleteUser(user)
		case message := <-c.controller.Broadcast: // HL
			c.data.AddMessageQueue(message)
		}
	} // HL
}

//END OMIT
