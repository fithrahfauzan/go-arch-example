package config

type MainConfig struct {
	AppName      string
	AppModule    AppModuleConfig
	Repository   RepositoryConfig
	UseCase      UseCaseConfig
	Presentation PresentationConfig
}

func NewMainConfig() MainConfig {
	// TODO, get config from files.

	return MainConfig{}
}
