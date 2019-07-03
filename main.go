package main

import (
	"fmt"

	"github.com/Depado/memegen/cmd"
	"github.com/Depado/memegen/generator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)

// Main command that will be run when no other command is provided on the
// command-line
var rootCmd = &cobra.Command{
	Use:  "memegen <picture> [options]",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := generator.Generate(
			args[0],
			viper.GetString("top"),
			viper.GetString("bottom"),
			viper.GetString("font.path"),
			viper.GetString("output"),
			viper.GetFloat64("font.size"),
		)
		if err != nil {
			logrus.WithError(err).Fatal("Unable to generate meme")
		}
	},
}

// Version command that will display the build number and version (if any)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run:   func(cmd *cobra.Command, args []string) { fmt.Printf("Build: %s\nVersion: %s\n", Build, Version) }, // nolint: unparam
}

func main() {
	cmd.AddAllFlags(rootCmd)
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Couldn't start")
	}
}
