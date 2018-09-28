package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

/* Parsed options */
var configPath string

/* Show version */
var showVersion bool

/* enable debug */
var debug bool

/* enable debug */
var help bool

var Version string
var StartTime time.Time
var Configuration interface{}
var Hostname string

func init() {
	RootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Print version information and quit")
	RootCmd.PersistentFlags().BoolVarP(&help, "help", "h", false, "Print  helpinformation and quit")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Path to configuration file")
	RootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config.toml", "Path to configuration file")
}

var RootCmd = &cobra.Command{
	Use:   "EgoismoBot",
	Short: "EgoismoBot",
	Run: func(cmd *cobra.Command, args []string) {

		if showVersion {
			fmt.Println(Version)
			return
		}

		if help {
			cmd.Help()
			return
		}

		FromFileCmd.Run(cmd, []string{configPath})
	},
}
