package config

type UserConfig struct {
	Email    string
	Password string
	Address  string
}

type Config struct {
	// остальные настройки (DB, HTTP-сервер и т.д.)
	User UserConfig
}
