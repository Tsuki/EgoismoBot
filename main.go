package main

import (
	"EgoismoBot/cmd"
	"EgoismoBot/config"
	"EgoismoBot/logging"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"time"
)

var version string

func init() {
	// Set GOMAXPROCS if not set
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	// Init random seed
	rand.Seed(time.Now().UnixNano())

	// Save info
	cmd.Version = version
	cmd.StartTime = time.Now()
	cmd.Hostname, _ = os.Hostname()

}
func main() {
	cmd.Execute(func(cfg *config.Config) {

		// Configure logging
		logging.Configure(cfg.LoggingConfig.AppName, cfg.LoggingConfig.Output, cfg.LoggingConfig.Level)
		log := logging.For("main")
		if cfg.Debug {
			log.Info("Enable debug")

		}
		log.Infof("%s starting v.%s", cfg.LoggingConfig.AppName, version)

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		log.Info("Shutdown")
	})
}
