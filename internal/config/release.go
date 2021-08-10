//+build !dev

package config

type Config struct {
	DebugMode bool `env:"DEBUG" envDefault:"false"`
	Verbose   bool `env:"VERBOSE" envDefault:"false"`

	// Address for publishing the API
	ApiHost   string `env:"API_HOST" envDefault:"127.0.0.1"`
	ApiPort   int `env:"API_PORT" envDefault:"3000"`

	// Timeouts (in seconds)
	ApiReadTimeout int `env:"API_READ_TIMEOUT" envDefault:"15"`
	ApiWriteTimeout int `env:"API_WRITE_TIMEOUT" envDefault:"15"`

	// Redis
	RedisHost string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort int `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDatabase int`env:"REDIS_DB" envDefault:"0"`
}
