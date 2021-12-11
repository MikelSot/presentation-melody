package main

////START OMIT
//func main() {
//	serveMux := http.NewServeMux() // HL
//	controller := domain.NewController()
//	user := &domain.User{}
//
//	data := data.NewUser(user, controller)
//	rw := infrastructure.NewUser()
//	cont := infrastructure.NewController(data, controller)
//	go cont.Run() // HL
//
//	handler := infrastructure.NewHandler(data, rw)
//	router := infrastructure.NewRouter(serveMux, handler, controller)
//	router.Router() // HL
//	log.Println(http.ListenAndServe(":8000", serveMux))
//}
//
////END OMIT
