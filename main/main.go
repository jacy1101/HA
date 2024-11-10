package main

import (
	"HA/cmd"
	_ "HA/cmd/config" // import sub command as module
	_ "HA/cmd/version"
)

func init() {
}

func main() {
	cmd.Execute()
}
