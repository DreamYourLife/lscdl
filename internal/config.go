package internal

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dreamyourlife/lscdl/pkg/tuya"
	"gopkg.in/yaml.v2"
)

type Config struct {
}

const (
	envConfName      = "TSCDL_CONFIG"
	defaultConfFile  = "/etc/tscdl/config.yaml"
)

var (
	debug      bool
	version    bool
	configFile string
)

func ProcessFlags() (err error) {
	if configFile != "" {
		return
	}

	flag.BoolVar(&debug, "x", false, "Add debugging output")
	flag.BoolVar(&version, "v", false, "Show version information")

	flag.StringVar(&configFile, "c", os.Getenv(envConfName), "Path to configfile")

	flag.Parse()

	if version {
		//nolint
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	if configFile == "" {
		configFile = defaultConfFile
	}

	configFile, err = filepath.EvalSymlinks(configFile)
	return err
}

func NewConfig() (config Config, err error) {
	if err = ProcessFlags(); err != nil {
		return
	}

	// This only parsed as yaml, nothing else
	// #nosec
	yamlConfig, err := os.ReadFile(configFile)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlConfig, &config)
	if debug {
		config.Debug = true
	}

	config.Direction = direction
	config.Initialize()

	return config, err
}

func (config *Config) Initialize() {
	//if err := config.RabbitMqConfig.Initialize(); err != nil {
	//	log.Fatalf("failed to initialize config: %e", err)
	//}
}
