package opts

type SomeServer struct {
	MaxClients int
}

type Option func(server *SomeServer)

func WithMaxClients(maxClients int) Option {
	return func(server *SomeServer) {
		server.MaxClients = maxClients
	}
}

func NewServer(opts ...func(server *SomeServer)) *SomeServer {
	srv := &SomeServer{}
	for _, opt := range opts {
		opt(srv)
	}

	return srv
}
