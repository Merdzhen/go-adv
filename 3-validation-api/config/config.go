package config

type UserConfig struct {
	Email    string
	Password string
	Address  string
	Host string
	Port int
}

type Config struct {
	// остальные настройки (DB, HTTP-сервер и т.д.)
	User UserConfig
}
