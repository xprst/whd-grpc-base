package server

// OptionFn configures options of server.
type OptionFn func(*Server)

// WithPort 设置端口号
func WithPort(port int32) OptionFn {
	return func(s *Server) {
		s.port = port
	}
}
