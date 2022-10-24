package defenderclienthttp

type Config struct {
	APIKey, APISecret, UserPoolID, ClientPoolID, APIURL string
}

func (cfg *Config) SetDefault() *Config {
	if cfg.UserPoolID == "" {
		cfg.UserPoolID = UserPoolIDDefault
	}

	if cfg.ClientPoolID == "" {
		cfg.ClientPoolID = ClientPoolIDDefault
	}

	if cfg.APIURL == "" {
		cfg.APIURL = APIURLDefault
	}

	return cfg
}
