package config

// DataBase defines a database struct
type DataBase struct {
	Host     string
	UserName string
	Password string
	Name     string
	Port     int
	SSL      string
}

// GetDataBase gets new database
func GetDataBase(host string, dbusername string, dbpassword string, dbname string, dbport int, sslmode string) DataBase {
	return DataBase{
		Host:     host,
		UserName: dbusername,
		Password: dbpassword,
		Name:     dbname,
		Port:     dbport,
		SSL:      sslmode,
	}
}
