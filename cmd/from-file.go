package cmd

import (
	"EgoismoBot/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"routegate/utils/codec"
)

func init() {
	RootCmd.AddCommand(FromFileCmd)
}

var FromFileCmd = &cobra.Command{
	Use:   "from-file <path>",
	Short: "Start using config from file",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			cmd.Help()
			return
		}

		data, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}

		var cfg config.Config
		if err = codec.Decode(string(data), &cfg, codec.TOML); err != nil {
			log.Fatal(err)
		}
		cfg.Debug = debug
		Configuration = struct {
			Kind string `json:"kind"`
			Path string `json:"path"`
		}{"file", args[0]}

		start(&cfg)
	},
}
