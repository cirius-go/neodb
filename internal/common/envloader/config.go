package envloader

import (
	"cirius-go/neodb/internal/common"
	"context"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Load loads configuration of type ConfigType using the provided EnvLoaders.
func Load[ConfigType any](ctx context.Context, loaders ...common.EnvLoader) (z ConfigType, err error) {
	if len(loaders) > 0 {
		for _, loader := range loaders {
			if err := loader(ctx); err != nil {
				return z, err
			}
		}
	}

	return env.ParseAs[ConfigType]()
}

// FromDotEnvFiles creates an EnvLoader that loads environment variables
// from the specified files.
func FromDotEnvFiles(files ...string) common.EnvLoader {
	return func(context.Context) error {
		return godotenv.Overload(files...)
	}
}
