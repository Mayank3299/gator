package config

type Config struct {
	DbUrl       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}
