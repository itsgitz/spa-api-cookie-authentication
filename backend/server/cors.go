package server

import "os"

func (s *Server) setCors() {
	s.CORS.AllowCredentials = true
	s.CORS.AllowHeaders = "Content-Type,Content-Length, Authorization, Accept,X-Requested-With, Set-Cookie, If-Modified-Since, Cache-Control, User-Agent, DNT"
	s.CORS.AllowOrigins = os.Getenv("CORS_ALLOW_ORIGIN")
	s.CORS.ExposeHeaders = "Set-Cookie"
}
