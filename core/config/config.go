package config

type Config struct {
	DB DataBaseConfig
}
type DataBaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}
