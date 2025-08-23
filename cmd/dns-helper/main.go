package main

import "dns-helper/internal/cli"

// Version information set during build
var (
	version   = "dev"
	commit    = "unknown"
	buildTime = "unknown"
)

func main() {
	cli.Execute()
}
