package main

func (s *Server) routes() {
	s.router.HandleFunc("/respond", s.handlerespond()).Methods("GET")
	s.router.HandleFunc("/user", s.handleregisteruser()).Methods("POST")
	s.router.HandleFunc("/user", s.handlegetuser()).Methods("GET")
	s.router.HandleFunc("/user", s.handleupdateuser()).Methods("PUT")
	s.router.HandleFunc("/user", s.handledeleteuser()).Methods("DELETE")
}
