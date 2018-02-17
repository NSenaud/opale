package cmd

import (
	"fmt"
	"os"

	"github.com/NSenaud/opale"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var conf, conferr = opale.LoadConfig()

var rootCmd = &cobra.Command{
	Use:   "opale",
	Short: "Opale is a tool to request the Opale server.",
	Long: `Opale allow you to request data in the Opale database via the
Opale server.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO Do Stuff Here
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Opale",
	Long:  `All software has versions. This is Opale's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Opale CLI tool v0.1 -- HEAD")
	},
}

var getCmd = &cobra.Command{
	Use:   "get [sensor]",
	Short: "Get a sensor's value",
	Long:  `Request the value of a sensor to the Opale server`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		get(&conf, args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(logInit)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(getCmd)
}

func logInit() {
	if conferr != nil {
		// FIXME Should use default settings instead
		log.Panic("Failed to load configuration file:", conferr)
	}

	if conf.Client.Debug {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.WarnLevel)
	}
}
