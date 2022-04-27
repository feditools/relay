package http

// Module represents a module that can be added to a http server
type Module interface {
	Name() string
	Route(s *Server) error
}
