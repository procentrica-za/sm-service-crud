package main

func (s *Server) routes() {
	// Service User Manager Routes
	s.router.HandleFunc("/user", s.handleregisteruser()).Methods("POST")
	s.router.HandleFunc("/user", s.handlegetuser()).Methods("GET")
	s.router.HandleFunc("/user", s.handleupdateuser()).Methods("PUT")
	s.router.HandleFunc("/user", s.handledeleteuser()).Methods("DELETE")
	s.router.HandleFunc("/userlogin", s.handleloginuser()).Methods("GET")
	s.router.HandleFunc("/forgotpassword", s.handleforgotpassword()).Methods("GET")

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

}
