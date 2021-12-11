package main

//START OMIT
func (u User) AddUser(user *domain.User) {
	u.controller.Users[user.Name] = user // HL
}

func (u User) DeleteUser(user *domain.User) {
	if _, state := u.controller.Users[user.Name]; state {
		delete(u.controller.Users, user.Name)
		close(user.Message) // HL
	}
}

func (u User) AddMessageQueue(message domain.Message) {
	for name, user := range u.controller.Users {
		if message.Username != name {
			select { // HL
			case user.Message <- message: // HL
			default:
				delete(u.controller.Users, user.Name)
				close(user.Message) // HL
			} // HL
		}
	}
}

//END OMIT
