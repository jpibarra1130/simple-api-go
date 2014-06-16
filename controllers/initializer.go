package controllers

import (
	"flag"
)

// global options. available to any subcommands. This was taken from goose library
var flagPath = flag.String("path", "db", "folder containing db info")
var flagEnv = flag.String("env", "development", "which DB environment to use")
