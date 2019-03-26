package config

import "os"

// Config хранит все данные для конфигурации сервера
// которые беруться из переменных окружения
type Config struct {
	// MongoURI Хранит URI для подключения к базе данных
	MongoURI string

	// Port Порт на котором будет сервер, используеться пременная gin PORT
	Port string

	// Host — домен, на котором работает API
	Host string
}

func parseEnv(envName, defaultValue string) string {
	if env := os.Getenv(envName); env != "" {
		return env
	}
	return defaultValue
}

// New коструктор для конфига
func New() Config {
	return Config{
		MongoURI: parseEnv("MONGODB_URI", "mongodb://localhost:27017"),
		Port:     parseEnv("PORT", "8080"),
		Host:     parseEnv("HOST_DOMAIN", "localhost"),
	}
}
