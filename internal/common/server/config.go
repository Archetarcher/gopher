package server

type Config struct {
	RunAddr string
	Session *Session
}

// Session is a struct for obtained session
type Session struct {
	Key string
}
