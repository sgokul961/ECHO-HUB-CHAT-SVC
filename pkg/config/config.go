package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthHubUrl    string `mapstructure:"auth_hub_url"`
	DB_URL        string `mapstructure:"DB_URL"`
	DBName        string `mapstructure:"DB_NAME"`
	DBUsername    string `mapstructure:"MONGO_INITDB_ROOT_USERNAME"`
	DBPassword    string `mapstructure:"MONGO_INITDB_ROOT_PASSWORD"`
	AuthMechanism string `mapstructure:"AUTH_MECHANISM"`
}

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./pkg/config/envs")

	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
