package config

// Redis define a new redis config struct
type Redis struct {
	Host     string
	Port     int
	DBNumber int
}

// GetRedis gets new redis
func GetRedis(host string, port int, dbnumber int) Redis {
	return Redis{
		Host:     host,
		Port:     port,
		DBNumber: dbnumber,
	}
}
