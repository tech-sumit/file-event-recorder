package env

import (
	"flag"
	"fmt"
	"os"

	"github.com/caarlos0/env/v9"
)

// Config maps the environment variables into a struct.
type Config struct {
	// AppEnv is the application environment that determines `configs/<APP_ENV>.env` to load.
	AppEnv string `env:"APP_ENV" envDefault:"development"`

	// DBPath is the primary database URI.
	DBPath string `env:"DB_PATH,notEmpty"`

	// DirPath is the folder which is going to be watched
	DirPath string `env:"DIR_PATH,notEmpty"`

	// MaxConcurrency is the concurrent tasks can run in any unit time
	MaxConcurrency int `env:"MAX_CONCURRENCY" envDefault:"100"`
}

// InitEnvConfig loads environment variables into Config struct.
func InitEnvConfig() (*Config, error) {
	config := Config{}

	// Attempt to parse environment variables
	if err := env.Parse(&config); err != nil {
		fmt.Printf("Error parsing environment variables: %+v\n", err)
		fmt.Println("Failed to load environment variables, attempting to load from command line args...")
		parseCommandLineArgs(&config) // Fallback to command line args
	} else {
		parseCommandLineArgs(&config) // Also parse CLI args for possible overrides
	}

	if config.DBPath == "" || config.DirPath == "" {
		fmt.Println("Critical configuration missing, cannot start application.")
		os.Exit(1) // Terminate if config is still incomplete
	}

	fmt.Printf("Loaded Configuration: %+v\n", config)
	return &config, nil
}

// parseCommandLineArgs defines and parses command line arguments to set config
func parseCommandLineArgs(config *Config) {
	flag.StringVar(&config.AppEnv, "appenv", config.AppEnv, "Application environment")
	flag.StringVar(&config.DBPath, "dbpath", config.DBPath, "Database path")
	flag.StringVar(&config.DirPath, "dirpath", config.DirPath, "Directory to watch")
	flag.IntVar(&config.MaxConcurrency, "maxconcurrency", config.MaxConcurrency, "Maximum concurrency level")

	// Parse command line arguments
	flag.Parse()

	// Extra check to print if command line flags are used
	if flag.NFlag() > 0 {
		fmt.Println("Command line arguments provided and parsed.")
	}
}
