package main

import (
	"context"

	"github.com/ivan-zherdev/benthos/v4/internal/cli"

	// Import all plugins defined within the repo.
	_ "github.com/ivan-zherdev/benthos/v4/public/components/all"
)

func main() {
	cli.Run(context.Background())
}
