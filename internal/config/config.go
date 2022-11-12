package config

type Config struct {
	Debug     bool   `json:"debug"`
	ServerURL string `json:"server_url"`
	DbURL     string `json:"db_url"`
}
