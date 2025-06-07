package initialize

import (
	"ecommerce/global"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)
func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // path to config
	viper.SetConfigName("dev")     // ten file
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	// configure structur
	if err := viper.Unmarshal(&global.Config); err != nil {
		global.Logger.Errorf("Unable to decode configuration %v", err)
	}
}

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		global.Logger.Errorf("No .env file found or error loading .env file")
	}
}
