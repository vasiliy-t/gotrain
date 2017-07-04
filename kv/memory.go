package main

// server is a dumb simple single threaded concurrent in memory kv
type server struct {
	store map[string]*string
	ops   chan (func())
}

// NewServer creates and returns new server instance
func NewServer() *server {
	s := &server{
		store: make(map[string]*string),
		ops:   make(chan func()),
	}
	go s.loop()
	return s
}

// loop is internal server operations loop,
// executes received from command channel functions
func (s *server) loop() {
	for {
		select {
		case op := <-s.ops:
			op()
		}
	}
}

// Get returns stored value
func (s *server) Get(key string) *string {
	res := make(chan *string)
	s.ops <- func() {
		v, ok := s.store[key]
		if !ok {
			res <- nil
			return
		}
		res <- v
	}
	return <-res
}

// Set stores new value
func (s *server) Set(key string, value *string) error {
	res := make(chan error)
	s.ops <- func() {
		s.store[key] = value
		res <- nil
	}
	return <-res
}

// Del removes value from server
func (s *server) Del(key string) {
	res := make(chan bool)
	s.ops <- func() {
		delete(s.store, key)
		res <- true
	}
	<-res
}
