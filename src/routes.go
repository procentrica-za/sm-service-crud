package main

func (s *Server) routes() {
	s.router.HandleFunc("/user", s.handleregisteruser()).Methods("POST")
	s.router.HandleFunc("/user", s.handlegetuser()).Methods("GET")
	s.router.HandleFunc("/user", s.handleupdateuser()).Methods("PUT")
	s.router.HandleFunc("/user", s.handledeleteuser()).Methods("DELETE")
	s.router.HandleFunc("/userlogin", s.handleloginuser()).Methods("GET")

	//Adhandler routes
	s.router.HandleFunc("/advertisement", s.handlepostadvertisement()).Methods("POST")
	s.router.HandleFunc("/advertisement", s.handleupdateadvertisement()).Methods("PUT")
	s.router.HandleFunc("/advertisement", s.handleremoveadvertisement()).Methods("DELETE")
	s.router.HandleFunc("/advertisement", s.handlegetadvertisement()).Methods("GET")
	s.router.HandleFunc("/advertisements", s.handledeleteuseradvertisements()).Methods("DELETE")
	//s.router.HandleFunc("/advertisementtype", s.handlegetadvertisementbytype()).Methods("GET")
	//s.router.HandleFunc("/advertisements", s.handlegetadvertisements()).Methods("GET")

}
