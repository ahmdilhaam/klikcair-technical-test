package helper

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory.
	viper.SetConfigFile(".env")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Multiple path config file to support local and container deployment.
	var (
		_, b, _, _ = runtime.Caller(0)
		root       = filepath.Join(filepath.Dir(b), "../../../")
	)
	viper.AddConfigPath(root)
	viper.AddConfigPath("/digital-wallet")
	viper.AddConfigPath("./")

	// Find and read the config file.
	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error.
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true.
	// ok will make sure the program not break.
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
