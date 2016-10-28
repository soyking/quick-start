package config

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

const (
	APP_NAME     = "user_tags"
	APP_REVISION = "0.9.0"

	// flag params
	FLAG_REV_NAME    = "rev"
	FLAG_REV_USAGE   = "print rev"
	FLAG_REV_DEFAULT = false

	FLAG_PORT_NAME    = "port"
	FLAG_PORT_USAGE   = "runing port"
	FLAG_PORT_DEFAULT = "8080"
)

type Config struct {
	Port string
}

func printRev() {
	fmt.Printf("%s v%s build on %s [%s, %s]\n",
		APP_NAME, APP_REVISION,
		runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

func InitConfig() *Config {
	var (
		rev  bool
		port string
	)

	flag.StringVar(&port, FLAG_PORT_NAME, FLAG_PORT_DEFAULT, FLAG_PORT_USAGE)
	flag.BoolVar(&rev, FLAG_REV_NAME, FLAG_REV_DEFAULT, FLAG_REV_USAGE)
	flag.Parse()

	printRev()
	if rev {
		os.Exit(0)
	}

	return &Config{
		Port: port,
	}
}
