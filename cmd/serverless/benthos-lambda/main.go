package main

import (
	"github.com/ivan-zherdev/benthos/v4/internal/serverless/lambda"

	// Import all plugins defined within the repo.
	_ "github.com/ivan-zherdev/benthos/v4/public/components/all"
)

func main() {
	lambda.Run()
}
