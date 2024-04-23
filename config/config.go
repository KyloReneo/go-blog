package config

//Define some custom types for reading configs from the .yml file
//Config struct includes both app and server types
type Config struct {
	App    App
	Server Server
	DB     DB
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}
