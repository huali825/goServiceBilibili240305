package config

type config struct {
	DB    DBConfig
	Redis RedisConfig
	HS    HelloString
}

type DBConfig struct {
	DSN string
}
type RedisConfig struct {
	Addr string
}

type HelloString struct {
	str string
}
