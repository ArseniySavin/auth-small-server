package config

const (
	PORT       = "PORT"
	ETCD_ULRS  = "ETCD_ULRS"
	JWT_SECRET = "JWT_SECRET"

	PG_HOST     = "PG_HOST"
	PG_PORT     = "PG_PORT"
	PG_LOGIN    = "PG_LOGIN"
	PG_PASS     = "PG_PASS"
	PG_DB       = "PG_DB"
	PG_SSL_MODE = "PG_SSL_MODE"
)

type Config struct {
	DB         DB     ` validate:"struct"`
	Port       string ` validate:"string"`
	ETCD_ULRS  string ` validate:"string"`
	JWT_SECRET string ` validate:"string"`
}

type DB struct {
	Login   string ` validate:"string"`
	Pass    string ` validate:"string"`
	Host    string ` validate:"string"`
	Port    int    ` validate:"number,min=1,max=10000"`
	Name    string ` validate:"string"`
	SslMode string ` validate:"string"`
}
