package config

import "github.com/spf13/viper"

type Config struct {
	ServicePort    string `mapstructure:"SERVICE_PORT"`
	UserClientPort string `mapstructure:"USER_SERVICE_PORT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBName         string `mapstructure:"DB_NAME"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
}

var envs = []string{
	"SERVICE_PORT", "USER_SERVICE_PORT",
	"DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", "DB_PASSWORD",
	"JWT_SECRET",
}

func LoadEnvs() (config *Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err = viper.BindEnv(env); err != nil {
			return
		}
	}

	err = viper.Unmarshal(&config)

	return
}
