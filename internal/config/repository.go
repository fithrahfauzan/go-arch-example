package config

type RepositoryConfig struct {
	TokenRepository TokenRepositoryConfig
}

type TokenRepositoryConfig struct {
	TokenExpiration int // Hour.
}
