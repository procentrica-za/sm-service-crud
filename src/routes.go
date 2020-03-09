package main

func (s *Server) routes() {
	// Service User Manager Routes
	s.router.HandleFunc("/user", s.handleregisteruser()).Methods("POST")
	s.router.HandleFunc("/user", s.handlegetuser()).Methods("GET")
	s.router.HandleFunc("/user", s.handleupdateuser()).Methods("PUT")
	s.router.HandleFunc("/user", s.handledeleteuser()).Methods("DELETE")
	s.router.HandleFunc("/userlogin", s.handleloginuser()).Methods("GET")
	s.router.HandleFunc("/forgotpassword", s.handleforgotpassword()).Methods("GET")
	s.router.HandleFunc("/userpassword", s.handleupdatepassword()).Methods("PUT")

	//Adhandler routes
	s.router.HandleFunc("/advertisement", s.handlepostadvertisement()).Methods("POST")
	s.router.HandleFunc("/advertisement", s.handleupdateadvertisement()).Methods("PUT")
	s.router.HandleFunc("/advertisement", s.handleremoveadvertisement()).Methods("DELETE")
	s.router.HandleFunc("/advertisement", s.handlegetadvertisement()).Methods("GET")
	s.router.HandleFunc("/useradvertisements", s.handlegetuseradvertisements()).Methods("GET")
	s.router.HandleFunc("/useradvertisements", s.handledeleteuseradvertisements()).Methods("DELETE")

	s.router.HandleFunc("/advertisementtype", s.handlegetadvertisementbytype()).Methods("GET")

	s.router.HandleFunc("/advertisementposttype", s.handlegetadvertisementbyposttype()).Methods("GET")
	s.router.HandleFunc("/advertisements", s.handlegetalladvertisements()).Methods("GET")
	s.router.HandleFunc("/textbooks", s.handlegettextbooksbyfilter()).Methods("GET")
	s.router.HandleFunc("/textbook", s.handleaddtextbook()).Methods("POST")
	s.router.HandleFunc("/textbook", s.handleupdatetextbook()).Methods("PUT")
	s.router.HandleFunc("/textbook", s.handleremovetextbook()).Methods("DELETE")

	s.router.HandleFunc("/notes", s.handlegetnotesbyfilter()).Methods("GET")
	s.router.HandleFunc("/note", s.handleaddnote()).Methods("POST")
	s.router.HandleFunc("/note", s.handleupdatenote()).Methods("PUT")
	s.router.HandleFunc("/note", s.handleremovenote()).Methods("DELETE")

	s.router.HandleFunc("/tutors", s.handlegettutorsbyfilter()).Methods("GET")
	s.router.HandleFunc("/tutor", s.handleaddtutor()).Methods("POST")
	s.router.HandleFunc("/tutor", s.handleupdatetutor()).Methods("PUT")
	s.router.HandleFunc("/tutor", s.handleremovetutor()).Methods("DELETE")

	s.router.HandleFunc("/accomodations", s.handlegetaccomodationsbyfilter()).Methods("GET")
	s.router.HandleFunc("/accomodation", s.handleaddaccomodation()).Methods("POST")
	s.router.HandleFunc("/accomodation", s.handleupdateaccomodation()).Methods("PUT")
	s.router.HandleFunc("/accomodation", s.handleremoveaccomodation()).Methods("DELETE")

	/*
		======================================= File Manager =========================================
	*/
	s.router.HandleFunc("/cardimage", s.handlegetcardimagepath()).Methods("GET")
	s.router.HandleFunc("/cardimagebatch", s.handlegetcardimagepathbatch()).Methods("POST")
	s.router.HandleFunc("/advertisementimages", s.handlegetadvertisementimages()).Methods("GET")
	/*
		======================================= Messaging =========================================
	*/

	s.router.HandleFunc("/chat", s.handleaddchat()).Methods("POST")
	s.router.HandleFunc("/chat", s.handledeletechat()).Methods("DELETE")
	s.router.HandleFunc("/chats", s.handlegetactivechats()).Methods("GET")
	s.router.HandleFunc("/message", s.handlegetmessages()).Methods("GET")
	s.router.HandleFunc("/message", s.handleaddmessage()).Methods("POST")
}
