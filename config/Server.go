package config

// Server defines a new configuration server.
type Server struct {
	Host string
	Port int
}

// GetServer gets a new server config
func GetServer(host string, port int) Server {
	return Server{
		Host: host,
		Port: port,
	}
}
