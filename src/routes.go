package main

func (s *Server) routes() {
	s.router.HandleFunc("/Respond", s.handlerespond()).Methods("GET")
	s.router.HandleFunc("/User", s.handleregisteruser()).Methods("POST")
}
