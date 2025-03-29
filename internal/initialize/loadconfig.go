package initialize

import (
	"ecommerce/global"
	"fmt"

	"github.com/spf13/viper"
)
func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // ten file
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	// configure structur
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
