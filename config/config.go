package config

import "github.com/spf13/viper"

func GetConfig(key string) (string, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		return "", err
	}

	value := viper.GetString(key)

	return value, nil
}
