package config

type PresentationConfig struct {
	HTTP HTTPConfig
	Cron CronConfig
}

type HTTPConfig struct {
	BaseURL string
}

type CronConfig struct {
	// TODO
}
