// Command speck turns a rough idea into a structured SPEC.md.
package main

import (
	"os"

	"github.com/palladius/respec/internal/cli"
	"github.com/palladius/respec/internal/dotenv"
)

func main() {
	_ = dotenv.Load(".env")
	if err := cli.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
