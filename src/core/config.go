package core

type ServerMode string

const (
	ServerDevMode     = "dev"
	ServerProdMode    = "prod"
	ServerTestingMode = "test"
)

type ServerConfig struct {
	Mode    ServerMode
	ShowLog bool
}
