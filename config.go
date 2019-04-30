package bamboohr

// Config contains configuration parameters of the BambooHR API.
type Config struct {
	apiDomain string
	apiKey    string
}

// NewConfig creates and returns a new Config object based on the given api domain and api key.
func NewConfig(apiDomain, apiKey string) *Config {
	return &Config{apiDomain, apiKey}
}

// APIDomain returns the api domain.
func (c *Config) APIDomain() string {
	return c.apiDomain
}

// APIKey returns the api key.
func (c *Config) APIKey() string {
	return c.apiKey
}
