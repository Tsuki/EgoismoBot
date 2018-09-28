package cmd

import "EgoismoBot/config"

var start func(*config.Config)

/**
 * Execute processing flags
 */
func Execute(f func(*config.Config)) {
	start = f
	RootCmd.Execute()
}
