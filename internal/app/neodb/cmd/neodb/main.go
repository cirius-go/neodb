package main

import (
	"context"
	"flag"
	"time"

	"cirius-go/neodb/internal/app/neodb/config"
	"cirius-go/neodb/internal/common"
	"cirius-go/neodb/internal/common/envloader"
)

// flags
var (
	configFlag = flag.String("config", "config.yaml", "Path to the configuration file")
)

func main() {
	flag.Parse()

	var (
		infraContext, cancelInfraContext = context.WithTimeout(context.Background(), 30*time.Second)
		envLoaders                       []common.EnvLoader
	)
	defer cancelInfraContext()

	if *configFlag != "" {
		envLoaders = append(envLoaders, envloader.FromDotEnvFiles(*configFlag))
	}

	cfg, err := envloader.Load[config.Config](infraContext, envLoaders...)
	panicIf(err)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
