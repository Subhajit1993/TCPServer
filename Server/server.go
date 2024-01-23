package Server

type ServerConfig struct {
	Port string
	Host string
}

type TCPServer interface {
	Start() (bool, error)
	Stop() string
}

type SingleThreaded struct {
	config ServerConfig
}

func (s SingleThreaded) Stop() string {
	panic("implement me")
}

type Concurrent struct {
	config ServerConfig
}

func (s Concurrent) Stop() string {
	//TODO implement me
	panic("implement me")
}

func StartServer(s TCPServer) (bool, error) {
	return s.Start()
}

func (s ServerConfig) Start() (bool, error) {
	newServer := SingleThreaded{
		config: s,
	}
	return StartServer(newServer)
}
