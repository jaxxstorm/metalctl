package main

import (
	"fmt"
	"github.com/jaxxstorm/metalctl/cmd/metalctl/create"
	"github.com/jaxxstorm/metalctl/cmd/metalctl/destroy"
	"github.com/jaxxstorm/metalctl/pkg/contract"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	org    string
	debug  bool
)

func configureCLI() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:  "metalctl",
		Long: "Create resources on Equinix Metal",
	}

	// commands
	rootCommand.AddCommand(create.Command())
	rootCommand.AddCommand(destroy.Command())

	rootCommand.PersistentFlags().StringVarP(&org, "org", "o", "", "Pulumi org to use for your stack")
	rootCommand.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug output")

	return rootCommand
}

func init() {
	log.SetLevel(log.InfoLevel)
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	if debug {
		log.SetLevel(log.DebugLevel)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.metalctl") // adding home directory as first search path
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file: ", viper.ConfigFileUsed())
	}
}

func main() {
	rootCommand := configureCLI()

	if err := rootCommand.Execute(); err != nil {
		contract.IgnoreIoError(fmt.Fprintf(os.Stderr, "%s", err))
		os.Exit(1)
	}
}
