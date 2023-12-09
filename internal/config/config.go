package config

type Config struct {
	DefaultFrom    string
	AllowedOrigins []string
	Redis          Redis
	Kafka          Kafka
	SMTP           SMTP
	RateLimit      int
	RabbitMQ       RabbitMQ
}

type Redis struct {
	Password string
	Host     string
	Port     int
	Topic    string
	Enabled  bool
}

type Kafka struct {
	Host    string
	Port    int
	Enabled bool
	Topic   string
	Brokers []string
}

type RabbitMQ struct {
	Host     string
	Port     int
	Username string
	Password string
}

type SMTP struct {
	Host     string
	Port     int
	Username string
	Password string
}
